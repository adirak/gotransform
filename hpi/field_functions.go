package hpi

import "strings"

// HpiRenameFields is function to rename field from input to output
func (c *Transform) HpiRenameFields(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_RenameFields, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of this function
	// =======================================

	// Prepare size
	inFields := config.Input
	outFields := config.Output
	size := len(inFields)
	if size > len(outFields) {
		size = len(outFields)
	}
	inputBus := c.GetInputBus()

	// Loop to rename
	for i := 0; i < size; i++ {
		inName := inFields[i]
		outName := outFields[i]
		value := inputBus.Value(inName)

		if value != nil {
			// Set data to output
			c.SetOutputData(outName, value)
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

// HpiDeleteField is function to delete field form data bus
func (c *Transform) HpiDeleteField(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_DeleteField, false, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	oName := config.Output[0]
	c.DeleteOutputData(oName)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiFixValue is function to declare field and its value
func (c *Transform) HpiFixValue(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_FixValue, false, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	valueType := config.ValueType
	value := config.Value
	var result interface{}

	if strings.EqualFold(valueType, "string") {
		result = value
	} else if strings.EqualFold(valueType, "integer") {
		result, err = c.ToInt64(value)
		if err != nil {
			return nil, err
		}
	} else if strings.EqualFold(valueType, "decimal") {
		result, err = c.ToFloat64(value)
		if err != nil {
			return nil, err
		}
	} else if strings.EqualFold(valueType, "number") {
		result, err = c.ToFloat64(value)
		if err != nil {
			return nil, err
		}
	} else if strings.EqualFold(valueType, "boolean") {
		result = c.ToBoolean(value)
	} else {
		result = value
	}

	oName := config.Output[0]
	c.SetOutputData(oName, result)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
