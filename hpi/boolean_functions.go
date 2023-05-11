package hpi

// HpiToBoolean is function to convert other type to boolean
func (c *Transform) HpiToBoolean(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_ToBoolean, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	inputBus := c.GetInputBus()
	inName := config.Input[0]
	inValue := inputBus.Value(inName)

	result := c.ToBoolean(inValue)

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
