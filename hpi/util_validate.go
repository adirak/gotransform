package hpi

import (
	"errors"
	"fmt"
	"strings"
)

// ValidateList is function to validate list
func (c *Transform) ValidateList(data interface{}) (list []interface{}, err error) {
	if data != nil {
		list, ok := data.([]interface{})
		if !ok {
			return nil, errors.New("data is not list")
		}

		return list, nil
	}
	return make([]interface{}, 0), nil
}

// ValidateListOfMap is function to validate list of map
func (c *Transform) ValidateListOfMap(data interface{}) (list []interface{}, err error) {
	if data != nil {

		list, err := c.ValidateList(data)

		if err != nil {
			return nil, err
		}

		if len(list) > 0 {
			obj := list[0]
			_, ok2 := obj.(map[string]interface{})
			if !ok2 {
				return nil, errors.New("data is not list of map")
			}
		}

		return list, nil
	}

	return make([]interface{}, 0), nil
}

// ValidateObjectByType is function to validate object by validate type
func (c *Transform) ValidateObjectByType(data interface{}, validateType string) bool {

	// Validate Null
	if strings.Contains(validateType, "Not Null") {
		if data == nil {
			return false
		}
	}

	// Validate Empty
	checkNumber := true
	if strings.Contains(validateType, "Not Empty") {
		str, ok := data.(string)
		if ok {
			if len(str) == 0 {
				return false
			}
			checkNumber = false
		}
		list, ok := data.([]interface{})
		if ok {
			if len(list) == 0 {
				return false
			}
			checkNumber = false
		}
		mapObj, ok := data.(map[string]interface{})
		if ok {
			if len(mapObj) == 0 {
				return false
			}
			checkNumber = false
		}
	}

	// Validate Zero
	if strings.Contains(validateType, "Not Zero") && checkNumber {
		fValue, err := c.ToFloat64(data)
		if err == nil {
			if fValue == 0 {
				return false
			}
		}
	}

	return true
}

// ValidateNullData is function to validate null and empty data
func (c *Transform) ValidateNullData(config *TransformConfig, checkEmpty bool) (err error) {

	inputBus := c.GetInputBus()

	for _, name := range config.Input {

		// Get value
		obj := inputBus.Value(name)
		if obj == nil {
			err = fmt.Errorf("%s is null", name)
			return err
		}

		// Check Empty
		if checkEmpty {

			// Map type
			mapObj, ok := obj.(map[string]interface{})
			if ok {
				if len(mapObj) <= 0 {
					err = fmt.Errorf("%s is empty map", name)
					return err
				}
			}

			// List Type
			list, ok := obj.([]interface{})
			if ok {
				if len(list) <= 0 {
					err = fmt.Errorf("%s is empty list", name)
					return err
				}
			}

			// List Type
			listMap, ok := obj.([]map[string]interface{})
			if ok {
				if len(listMap) <= 0 {
					err = fmt.Errorf("%s is empty list", name)
					return err
				}
			}

		}

	}

	return
}
