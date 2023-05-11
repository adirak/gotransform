package hpi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// HpiJsonToString is function to convert json object to json string
func (c *Transform) HpiJsonToString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_JsonToString, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Loop the input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	oName := config.Output[0]
	obj := inputBus.Value(inName)

	// Map data
	data, ok := obj.(map[string]interface{})
	if ok {

		// Map to json string
		jsonStr, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		// Keep result
		c.SetOutputData(oName, jsonStr)

	} else {

		arrData, okArr := obj.([]interface{})

		if okArr {

			// Map to json string
			jsonStr, err := json.Marshal(arrData)
			if err != nil {
				return nil, err
			}

			// Keep result
			c.SetOutputData(oName, jsonStr)

		} else {

			str, ok := obj.(string)
			if ok {
				c.SetOutputData(oName, str)
			} else {

				if obj != nil {
					str := fmt.Sprintf("%v", obj)
					c.SetOutputData(oName, str)
				}

			}

		}

	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiStringToJson is function to convert string to json object
func (c *Transform) HpiStringToJson(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_StringToJson, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Loop the input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	oName := config.Output[0]

	obj := inputBus.Value(inName)
	if obj == nil {
		return nil, errors.New("input is null")
	}

	str, ok := obj.(string)
	if !ok {
		return nil, errors.New("input is not string")
	}

	// trim space
	str = strings.TrimSpace(str)

	// Map data
	if strings.HasPrefix(str, "{") {

		var result map[string]interface{}
		err = json.Unmarshal([]byte(str), &result)
		if err != nil {
			return nil, err
		}

		// Keep result
		c.SetOutputData(oName, result)
	} else {

		if strings.HasPrefix(str, "[") {
			var result []interface{}
			err = json.Unmarshal([]byte(str), &result)
			if err != nil {
				return nil, err
			}

			// Keep result
			c.SetOutputData(oName, result)
		} else {
			err = errors.New("cannot convert to json")
		}
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, err
	}
	// Return
	return c.GetOutput(), err

}
