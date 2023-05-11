package hpi

import (
	"fmt"
	"strings"
)

// HpiNumberFormat is function to format the number value to string format
func (c *Transform) HpiNumberFormat(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_NumberFormat, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Logic of this function
	// =======================================
	intDigit := config.IntDigit
	padding := config.Padding
	decDigit := config.DecDigit
	thousandSeparator := config.ThousandSeparator
	decimalSeparator := config.DecimalSeparator

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)
	num, err := c.ToFloat64(value)
	if err != nil {
		return nil, err
	}

	// Get data by format
	result := c.FormatNumber(num, intDigit, padding, decDigit, thousandSeparator, decimalSeparator)

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

// HpiNameFormat is function to format string by naming type
func (c *Transform) HpiNameFormat(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_NameFormat, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Logic of this function
	// =======================================

	format := config.Format

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)

	strValue, ok := value.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", value)
	}

	// Get data by format
	result := c.FormatNaming(strValue, format)

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

// HpiIncreaseNumber is function to increase number
func (c *Transform) HpiIncreaseNumber(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_IncreaseNum, true, true)
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
	step := config.Step

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)
	inValue, err := c.ToInt64(value)
	if err != nil {
		return nil, err
	}

	result := inValue + step

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

// HpiDecreaseNumber is function to decrease number
func (c *Transform) HpiDecreaseNumber(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_DecreaseNum, true, true)
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
	step := config.Step

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)
	inValue, err := c.ToInt64(value)
	if err != nil {
		return nil, err
	}

	result := inValue - step

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

// HpiToNumber is function to convert other type to number data
func (c *Transform) HpiToNumber(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_ToNumber, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================
	valueType := config.ValueType

	// input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	inValue := inputBus.Value(inName)
	var result interface{}

	if inValue == nil {
		return nil, fmt.Errorf("input is null")
	}

	if strings.EqualFold(valueType, "integer") {
		result, err = c.ToInt64(inValue)
		if err != nil {
			return nil, err
		}
	} else if strings.EqualFold(valueType, "decimal") {
		result, err = c.ToFloat64(inValue)
		if err != nil {
			return nil, err
		}
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
