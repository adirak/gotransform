package hpi

import (
	"strings"
)

// HpiCheckValue is function to check value whether it's match with config value
func (c *Transform) HpiCheckValue(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CheckValue, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inValue := inputBus.Value(inName0)
	valuses := c.ToArrayValues(config.Value)
	compare := config.Compare

	// Call check value
	check := c.CheckValueRecord(inValue, valuses, compare)

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, check)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiCompareNumber is function to compare number value
func (c *Transform) HpiCompareNumber(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CompareNumber, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================
	compare := config.Compare

	// Prepare Input
	inputBus := c.GetInputBus()
	inName1 := config.Input[0]
	inName2 := config.Input[1]
	inValue1 := inputBus.Value(inName1)
	inValue2 := inputBus.Value(inName2)

	// Convert to float
	fValue1, err := c.ToFloat64(inValue1)
	if err != nil {
		return nil, err
	}

	fValue2, err := c.ToFloat64(inValue2)
	if err != nil {
		return nil, err
	}

	// Comparing...
	check := false
	compare = strings.ToLower(compare)
	if strings.Contains(compare, "==") {
		check = fValue1 == fValue2
	} else if strings.Contains(compare, "!=") {
		check = fValue1 != fValue2
	} else if strings.Contains(compare, "<") {
		check = fValue1 < fValue2
	} else if strings.Contains(compare, ">") {
		check = fValue1 > fValue2
	} else if strings.Contains(compare, "<=") {
		check = fValue1 <= fValue2
	} else if strings.Contains(compare, ">=") {
		check = fValue1 >= fValue2
	}

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, check)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiCompareString is function to compare string value
func (c *Transform) HpiCompareString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CompareString, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================
	compare := config.Compare

	// Prepare Input
	inputBus := c.GetInputBus()
	inName1 := config.Input[0]
	inName2 := config.Input[1]
	inValue1 := inputBus.Value(inName1)
	inValue2 := inputBus.Value(inName2)

	// Comparing...
	check := c.IsEqual(inValue1, inValue2, compare)

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, check)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiCompareDateTime is function to compare date and time
func (c *Transform) HpiCompareDateTime(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CompareDateTime, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================
	format := config.Format
	compare := config.Compare
	valueType := config.ValueType

	// Prepare Input
	inputBus := c.GetInputBus()
	inName1 := config.Input[0]
	inName2 := config.Input[1]
	inValue1 := inputBus.Value(inName1)
	inValue2 := inputBus.Value(inName2)

	// Comparing...
	check, err := c.CompareDateTime(inValue1, inValue2, compare, format, valueType)
	if err != nil {
		return nil, err
	}

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, check)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
