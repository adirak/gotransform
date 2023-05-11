package hpi

import (
	"fmt"
	"strings"
	"time"
)

// IsEqual is function to check value of two fields wether it is equal
func (c *Transform) IsEqual(v1 interface{}, v2 interface{}, compareType string) bool {

	// Not equal if some field is nil
	if v1 == nil || v2 == nil {
		return false
	}

	// Set default type
	if compareType == "" {
		compareType = "EqualIgnorecase"
	}

	// Replace space
	compareType = strings.ReplaceAll(compareType, " ", "")

	// Convert value to string
	strX := fmt.Sprintf("%v", v1)
	strY := fmt.Sprintf("%v", v2)

	// check with case
	if strings.EqualFold(compareType, "Equal") {
		return strX == strY
	} else if strings.EqualFold(compareType, "EqualIgnorecase") {
		return strings.EqualFold(strX, strY)
	} else if strings.EqualFold(compareType, "Contain") {
		return strings.Contains(strX, strY)
	} else if strings.EqualFold(compareType, "ContainIgnorecase") {
		strX = strings.ToLower(strX)
		strY = strings.ToLower(strY)
		return strings.Contains(strX, strY)
	} else if strings.EqualFold(compareType, "Prefix") {
		return strings.HasPrefix(strX, strY)
	} else if strings.EqualFold(compareType, "PrefixIgnorecase") {
		strX = strings.ToLower(strX)
		strY = strings.ToLower(strY)
		return strings.HasPrefix(strX, strY)
	} else if strings.EqualFold(compareType, "Suffix") {
		return strings.HasSuffix(strX, strY)
	} else if strings.EqualFold(compareType, "SuffixIgnorecase") {
		strX = strings.ToLower(strX)
		strY = strings.ToLower(strY)
		return strings.HasSuffix(strX, strY)
	}

	return false
}

// CheckValueRecord is function to check condition for filtering record
func (c *Transform) CheckValueRecord(obj interface{}, values []interface{}, compareType string) bool {

	// init
	if obj == nil || len(values) <= 0 {
		return false
	}

	// Array object
	arr, isArr := obj.([]interface{})
	if isArr {
		for _, item := range arr {
			chk := c.CheckValueRecord(item, values, compareType)
			if chk {
				return true
			}
		}
	} else {

		// Map object
		mapObj, isMap := obj.(map[string]interface{})
		if isMap {
			for _, mapValue := range mapObj {
				chk := c.CheckValueRecord(mapValue, values, compareType)
				if chk {
					return true
				}
			}

		} else {
			// General object
			for _, value := range values {
				if c.IsEqual(obj, value, compareType) {
					return true
				}
			}
		}
	}

	return false
}

// CheckValueRecord is function to check condition for filtering record
func (c *Transform) CompareDateTime(v1 interface{}, v2 interface{}, compareType string, format string, valueType string) (bool, error) {

	// Convert to string
	sv1 := fmt.Sprintf("%v", v1)
	sv2 := fmt.Sprintf("%v", v2)

	tData1, err := time.Parse(format, sv1)
	if err != nil {
		return false, err
	}
	tData2, err := time.Parse(format, sv2)
	if err != nil {
		return false, err
	}

	// init date
	date1 := ""
	date2 := ""
	if strings.EqualFold(valueType, "Date") {
		dFormat := "2006-01-02"
		date1 = tData1.Format(dFormat)
		date2 = tData2.Format(dFormat)
	}

	if strings.EqualFold(compareType, "Equal") {
		if strings.EqualFold(valueType, "Time") {
			return tData1.Equal(tData2), nil
		} else {
			return (strings.Compare(date1, date2) == 0), nil
		}
	} else if strings.EqualFold(compareType, "Before") {
		if strings.EqualFold(valueType, "Time") {
			return tData1.Before(tData2), nil
		} else {
			return (strings.Compare(date1, date2) < 0), nil
		}
	} else if strings.EqualFold(compareType, "After") {
		if strings.EqualFold(valueType, "Time") {
			return tData1.After(tData2), nil
		} else {
			return (strings.Compare(date1, date2) > 0), nil
		}
	} else if strings.EqualFold(compareType, "Before or Equal") {
		if strings.EqualFold(valueType, "Time") {
			return (tData1.Equal(tData2) || tData1.Before(tData2)), nil
		} else {
			return (strings.Compare(date1, date2) <= 0), nil
		}
	} else if strings.EqualFold(compareType, "After or Equal") {
		if strings.EqualFold(valueType, "Time") {
			return (tData1.Equal(tData2) || tData1.After(tData2)), nil
		} else {
			return (strings.Compare(date1, date2) >= 0), nil
		}
	}

	return false, nil

}
