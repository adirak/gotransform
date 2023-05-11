package hpi

import (
	"fmt"
	"strings"
)

// HpiGetKeys is function to list key of map value to output
func (c *Transform) HpiGetKeysInMap(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_GetKeysInMap, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Loop the input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	obj := inputBus.Value(inName)
	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("object is not map")
	}

	// list key in map
	keys := make([]string, 0)
	for key := range mapObj {
		keys = append(keys, key)
	}

	oName := config.Output[0]

	// Keep result
	c.SetOutputData(oName, keys)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiGetValuesInMap is function to get values by key of map value to output
func (c *Transform) HpiGetValuesInMap(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_GetValuesInMap, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Loop the input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	obj := inputBus.Value(inName)
	data, ok := obj.(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("object is not map")
	}

	// Convert to jsonPath
	jsonPath, err := NewJsonPathWithRoot(&data)
	if err != nil {
		return nil, err
	}

	// field mapping
	fields := config.Fields
	arrFields := c.ToArrayValues(fields)
	arrLen := len(arrFields)

	// list key in map
	for i, oName := range config.Output {
		path := oName

		// Make path
		if i < arrLen {
			field := arrFields[i]
			nPath := fmt.Sprintf("%v", field)
			nPath = strings.TrimSpace(nPath)
			if len(nPath) > 0 {
				path = nPath
			}
		}

		value := jsonPath.Value(path)

		// Keep result
		c.SetOutputData(oName, value)
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiCopyObjectToSameParent is function to copy value in the map to same parent of output
func (c *Transform) HpiCopyObjectToSameParent(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CopyObjectToSameParent, true, true)
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
	data, ok := obj.(map[string]interface{})

	if ok {

		parent := c.GetParentFromPath(oName)

		// Map Object copy all keys
		for key, value := range data {
			path := key
			if len(parent) > 0 {
				path = parent + "." + key
			}
			c.SetOutputData(path, value)
		}

	} else {
		// Keep result
		c.SetOutputData(oName, obj)
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiValidateAndCopy is function to validate value before copy it to output
func (c *Transform) HpiValidateAndCopy(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_ValidateAndCopy, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	compare := config.Compare

	// Loop the input
	inputBus := c.GetInputBus()
	for i, inName := range config.Input {
		if i < len(config.Output) {
			outName := config.Output[i]
			inValue := inputBus.Value(inName)
			if c.ValidateObjectByType(inValue, compare) {
				// Copy data
				c.SetOutputData(outName, inValue)
			} else {
				// clear data
				c.DeleteOutputData(outName)
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

// HpiCheckExistingField is function to check existing field in object
func (c *Transform) HpiCheckExistingField(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CheckExistingField, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	fields := config.Fields
	combine := config.Combine

	// Loop the input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	oName := config.Output[0]
	obj := inputBus.Value(inName)

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("object is not map")
	}

	arrFields := c.ToArrayFields(fields)

	// No key
	if len(arrFields) == 0 {
		c.SetOutputData(oName, false)
	}

	// Loop to check
	jsonPath, err := NewJsonPathWithRoot(&mapObj)
	if err != nil {
		return nil, err
	}

	count := 0
	for _, field := range arrFields {
		value := jsonPath.Value(field)
		if value != nil {
			count++
		}
	}

	// AND condition
	if strings.EqualFold(combine, "AND") {
		// Check existing by count
		if count >= len(arrFields) {
			c.SetOutputData(oName, true)
		} else {
			c.SetOutputData(oName, false)
		}
	} else {
		// Check existing by count
		if count > 0 {
			c.SetOutputData(oName, true)
		} else {
			c.SetOutputData(oName, false)
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
