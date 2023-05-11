package hpi

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// SortList is function to sort list
func (c *Transform) SortList(list []interface{}, sortType string, field string, valueType string, format string) []interface{} {

	// Never sort
	if len(list) < 2 {
		return list
	}

	// Check child
	child, ok := list[0].(map[string]interface{})
	if ok {
		if strings.TrimSpace(field) == "" {
			return list
		}
		if child[field] == nil {
			return list
		}
		// Sort map list
		return c.SortListMap(list, sortType, field, valueType, format)
	} else {
		// Sort object list
		return c.SortListObject(list, sortType, valueType, format)
	}
}

func (c *Transform) SortListMap(list []interface{}, sortType string, field string, valueType string, format string) []interface{} {

	sort.SliceStable(list, func(i, j int) bool {

		// get map
		mapi, oki := list[i].(map[string]interface{})
		mapj, okj := list[j].(map[string]interface{})

		if !oki || !okj {
			return false
		}

		// get object
		obji := mapi[field]
		objj := mapj[field]

		// Do compare
		num := c.CompareObject(obji, objj, valueType, format)

		// Compare by type
		if strings.EqualFold(sortType, "ASC") {
			return num < 0
		} else {
			return num > 0
		}
	})

	return list
}

func (c *Transform) SortListObject(list []interface{}, sortType string, valueType string, format string) []interface{} {

	sort.SliceStable(list, func(i, j int) bool {

		obji := list[i]
		objj := list[j]

		// Do compare
		num := c.CompareObject(obji, objj, valueType, format)

		// Compare by type
		if strings.EqualFold(sortType, "ASC") {
			return num < 0
		} else {
			return num > 0
		}
	})

	return list
}

// CompareObject is function to compare object by type
func (c *Transform) CompareObject(obji interface{}, objj interface{}, valueType string, format string) int {
	num := 0
	if strings.EqualFold(valueType, "String") {
		num = c.CompareString(obji, objj)
	} else if strings.EqualFold(valueType, "Number") {
		num = c.CompareNumber(obji, objj)
	} else if strings.EqualFold(valueType, "Number in String") {
		numi, erri := c.GetNumberInString(obji)
		numj, errj := c.GetNumberInString(objj)
		if erri == nil && errj == nil {
			num = c.CompareNumber(numi, numj)
		}
	} else if strings.EqualFold(valueType, "Date Time") {
		num = c.CompareTime(obji, objj, format)
	}
	return num
}

// CompareString is function to compare string value
func (c *Transform) CompareString(vi interface{}, vj interface{}) int {
	si := fmt.Sprintf("%v", vi)
	sj := fmt.Sprintf("%v", vj)
	si = strings.ToLower(si)
	sj = strings.ToLower(sj)
	return strings.Compare(si, sj)
}

// CompareNumber is function to compare number value
func (c *Transform) CompareNumber(vi interface{}, vj interface{}) int {
	si, err := c.ToFloat64(vi)
	if err != nil {
		return 0
	}

	sj, err := c.ToFloat64(vj)
	if err != nil {
		return 0
	}

	if si < sj {
		return -1
	} else if sj < si {
		return 1
	}
	return 0
}

// CompareTime is function to compare date time
func (c *Transform) CompareTime(vi interface{}, vj interface{}, format string) int {
	si := fmt.Sprintf("%v", vi)
	sj := fmt.Sprintf("%v", vj)

	ti, err := time.Parse(format, si)
	if err != nil {
		return 0
	}

	tj, err := time.Parse(format, sj)
	if err != nil {
		return 0
	}

	if ti.Before(tj) {
		return -1
	} else if tj.Before(ti) {
		return 1
	}
	return 0
}
