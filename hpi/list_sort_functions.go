package hpi

// HpiSortList is function to sort data in the list
func (c *Transform) HpiSortList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SortList, true, true)
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

	sortType := config.Compare
	field := config.Fields
	valueType := config.ValueType
	format := config.Format

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inListObj0 := inputBus.Value(inName0)

	// Validate List of Map
	list0, err := c.ValidateList(inListObj0)
	if err != nil {
		return nil, err
	}

	// Prepare Loop
	list := c.SortList(list0, sortType, field, valueType, format)

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, list)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
