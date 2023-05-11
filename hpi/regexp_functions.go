package hpi

import (
	"fmt"
)

// HpiFindValuesByRegExp is function to find values by regular expression
func (c *Transform) HpiFindValuesByRegExp(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_FindValuesByRegExp, true, true)
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

	regExp := config.RegExp

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inValue := inputBus.Value(inName0)
	oName := config.Output[0]

	// Check null
	if inValue == nil {
		return nil, fmt.Errorf("input is null")
	}

	// Call to find value
	values, err := c.FindValuesByRegExp(inValue, regExp)
	if err != nil {
		return nil, err
	}

	// Put data to output
	c.SetOutputData(oName, values)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
