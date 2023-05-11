package hpi

import (
	"fmt"
	"strings"
	"time"
)

// HpiGetCurrentTime is function to get current time from this system
func (c *Transform) HpiDateFormat(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_DateFormat, true, true)
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
	fromFormat := config.From
	toFormat := config.To
	if fromFormat == "" {
		fromFormat = "2006-01-02T15:04:05Z07:00"
	}
	if toFormat == "" {
		toFormat = "2006-01-02T15:04:05Z07:00"
	}

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.String("", inName)
	timeData, err := time.Parse(fromFormat, value)
	if err != nil {
		return nil, err
	}

	result := timeData.Format(toFormat)
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

// HpiGetCurrentTime is function to get current time from this system
func (c *Transform) HpiGetCurrentTime(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_GetCurrentTime, false, true)
	if err != nil {
		return nil, err
	}

	// Logic of this function
	// =======================================
	format := config.Format
	if format == "" {
		format = "2006-01-02T15:04:05Z07:00"
	}

	// Current time
	cTime := time.Now()
	curentTime := cTime.Format(format)

	oName := config.Output[0]
	c.SetOutputData(oName, curentTime)
	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiGetCurrentTimeStamp is function to get current time from this system
func (c *Transform) HpiGetCurrentTimeStamp(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_GetCurrentTimeStamp, false, true)
	if err != nil {
		return nil, err
	}

	// Logic of this function
	// =======================================
	valueType := config.ValueType
	timeData := time.Now()

	rValue := int64(0)
	if strings.EqualFold(valueType, "millisec") {
		// millisec
		rValue = timeData.UnixMilli()
	} else if strings.EqualFold(valueType, "microsec") {
		// microsec
		rValue = timeData.UnixMicro()
	} else if strings.EqualFold(valueType, "nanosec") {
		// nanosec
		rValue = timeData.UnixNano()
	} else {
		// sec
		sec := timeData.UnixMilli() / 1000
		rValue = sec
	}

	oName := config.Output[0]
	c.SetOutputData(oName, rValue)
	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiTimeStampToDate is function to convert time stamp to date format
func (c *Transform) HpiTimeStampToDate(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_TimeStampToDate, true, true)
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
	valueType := config.ValueType
	format := config.Format
	if format == "" {
		format = "2006-01-02T15:04:05Z07:00"
	}

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)

	// Convert to int64
	intValue, err := c.ToInt64(value)
	if err != nil {
		return nil, err
	}

	// Convert to time data
	var timeData time.Time
	if strings.EqualFold(valueType, "millisec") {
		// millisec
		nsec := intValue * 1000000
		timeData = time.Unix(0, nsec)
	} else if strings.EqualFold(valueType, "microsec") {
		// microsec
		nsec := intValue * 1000
		timeData = time.Unix(0, nsec)
	} else if strings.EqualFold(valueType, "nanosec") {
		// nanosec
		timeData = time.Unix(0, intValue)
	} else {
		// sec
		timeData = time.Unix(intValue, 0)
	}

	// Set output
	result := timeData.Format(format)

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

// HpiDateToTimeStamp is function to convert date to time stamp
func (c *Transform) HpiDateToTimeStamp(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_DateToTimeStamp, true, true)
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
	valueType := config.ValueType
	format := config.Format
	if format == "" {
		format = "2006-01-02T15:04:05Z07:00"
	}

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)

	strValue := fmt.Sprintf("%v", value)

	// Convert to time data
	timeData, err := time.Parse(format, strValue)
	if err != nil {
		return nil, err
	}

	rValue := int64(0)
	if strings.EqualFold(valueType, "millisec") {
		// millisec
		rValue = timeData.UnixMilli()
	} else if strings.EqualFold(valueType, "microsec") {
		// microsec
		rValue = timeData.UnixMicro()
	} else if strings.EqualFold(valueType, "nanosec") {
		// nanosec
		rValue = timeData.UnixNano()
	} else {
		// sec
		sec := timeData.UnixMilli() / 1000
		rValue = sec
	}

	// Set output
	oName := config.Output[0]
	c.SetOutputData(oName, rValue)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
