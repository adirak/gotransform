package hpi

import (
	"math"
	"strings"
)

// HpiCalculateNumber is function to calculate number value
func (c *Transform) HpiCalculateNumber(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CalculateNumber, true, true)
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
	operator := config.Operator
	valueType := config.ValueType

	// input
	inputBus := c.GetInputBus()
	inName1 := config.Input[0]
	inName2 := config.Input[1]
	inValue1 := inputBus.Value(inName1)
	inValue2 := inputBus.Value(inName2)

	// Convert to float
	fValue1, err := c.ToFloat64(inValue1)
	fValue2, err2 := c.ToFloat64(inValue2)
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}

	// Calulating..
	result := 0.0
	operator = strings.ToLower(operator)
	if strings.Contains(operator, "plus") {
		result = fValue1 + fValue2
	} else if strings.Contains(operator, "minus") {
		result = fValue1 - fValue2
	} else if strings.Contains(operator, "multipl") {
		result = fValue1 * fValue2
	} else if strings.Contains(operator, "divide") {
		result = fValue1 / fValue2
	} else if strings.Contains(operator, "power") {
		result = math.Pow(fValue1, fValue2)
	} else if strings.Contains(operator, "mod") {
		result = math.Mod(fValue1, fValue2)
	} else if strings.Contains(operator, "dim") {
		result = math.Dim(fValue1, fValue2)
	} else if strings.Contains(operator, "max") {
		result = math.Max(fValue1, fValue2)
	} else if strings.Contains(operator, "min") {
		result = math.Min(fValue1, fValue2)
	}

	// Set Result
	oName := config.Output[0]
	if strings.EqualFold(valueType, "integer") {
		result = math.Round(result)
		rInt := int64(result)
		c.SetOutputData(oName, rInt)
	} else {
		c.SetOutputData(oName, result)
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiMathUtil is function about utilities of math package
func (c *Transform) HpiMathUtil(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_MathUtil, true, true)
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
	valueType := config.ValueType
	operator := config.Operator

	// input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	inValue := inputBus.Value(inName)

	// Convert to float
	fValue, err := c.ToFloat64(inValue)
	if err != nil {
		return nil, err
	}

	// Calulating..
	result := 0.0
	operator = strings.ToLower(operator)
	if strings.EqualFold(operator, "Abs") {
		result = math.Abs(fValue)
	} else if strings.EqualFold(operator, "Acos") {
		result = math.Acos(fValue)
	} else if strings.EqualFold(operator, "Asin") {
		result = math.Asin(fValue)
	} else if strings.EqualFold(operator, "Atan") {
		result = math.Atan(fValue)
	} else if strings.EqualFold(operator, "Ceil") {
		result = math.Ceil(fValue)
	} else if strings.EqualFold(operator, "Cos") {
		result = math.Cos(fValue)
	} else if strings.EqualFold(operator, "Floor") {
		result = math.Floor(fValue)
	} else if strings.EqualFold(operator, "Gamma") {
		result = math.Gamma(fValue)
	} else if strings.EqualFold(operator, "Log10") {
		result = math.Log10(fValue)
	} else if strings.EqualFold(operator, "Log2") {
		result = math.Log2(fValue)
	} else if strings.EqualFold(operator, "Log") {
		result = math.Log(fValue)
	} else if strings.EqualFold(operator, "Round") {
		result = math.Round(fValue)
	} else if strings.EqualFold(operator, "RoundToEven") {
		result = math.RoundToEven(fValue)
	} else if strings.EqualFold(operator, "Sin") {
		result = math.Sin(fValue)
	} else if strings.EqualFold(operator, "Sqrt") {
		result = math.Sqrt(fValue)
	} else if strings.EqualFold(operator, "Tan") {
		result = math.Tan(fValue)
	} else if strings.EqualFold(operator, "Trunc") {
		result = math.Trunc(fValue)
	}

	// Set Result
	oName := config.Output[0]
	if strings.EqualFold(valueType, "integer") {
		result = math.Round(result)
		rInt := int64(result)
		c.SetOutputData(oName, rInt)
	} else {
		c.SetOutputData(oName, result)
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
