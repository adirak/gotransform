package hpi

import (
	"encoding/base64"
	"fmt"
)

// HpiBase64Encode is function to encode string to base64
func (c *Transform) HpiBase64Encode(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_Base64Encode, true, true)
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

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inValue := inputBus.Value(inName0)

	strValue, ok := inValue.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", inValue)
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(strValue))

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, encoded)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiBase64Decode is function to decode base64 to string value
func (c *Transform) HpiBase64Decode(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_Base64Decode, true, true)
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

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inValue := inputBus.Value(inName0)

	strValue, ok := inValue.(string)
	if !ok {
		return nil, fmt.Errorf("input is not string cannot decode")
	}

	decoded, err := base64.StdEncoding.DecodeString(strValue)
	if err != nil {
		decoded, err = base64.RawStdEncoding.DecodeString(strValue)
	}
	if err != nil {
		decoded, err = base64.RawURLEncoding.DecodeString(strValue)
	}
	if err != nil {
		return
	}

	decodedStr := string(decoded)

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, decodedStr)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiAes256GcmEncode is function to encode string to AES256-GCM
func (c *Transform) HpiAes256GcmEncode(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_Aes256GcmEncode, true, true)
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

	fields := config.Fields
	passcode := c.ToValueParam(fields)

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inValue := inputBus.Value(inName0)

	strValue, ok := inValue.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", inValue)
	}

	strPasscode, ok := passcode.(string)
	if ok {
		strPasscode = fmt.Sprintf("%v", strPasscode)
	}

	// Call encrypt function
	encoded, err := c.EncryptWithAes256Gcm(strValue, strPasscode)
	if err != nil {
		return nil, err
	}

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, encoded)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiAes256GcmDecode is function to decode string to AES256-GCM
func (c *Transform) HpiAes256GcmDecode(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_Aes256GcmDecode, true, true)
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

	fields := config.Fields
	passcode := c.ToValueParam(fields)

	// Prepare Input
	inputBus := c.GetInputBus()
	inName0 := config.Input[0]
	inValue := inputBus.Value(inName0)

	strValue, ok := inValue.(string)
	if !ok {
		strValue = fmt.Sprintf("%v", inValue)
	}

	strPasscode, ok := passcode.(string)
	if ok {
		strPasscode = fmt.Sprintf("%v", strPasscode)
	}

	// Call encrypt function
	decoded, err := c.DecryptWithAes256Gcm(strValue, strPasscode)
	if err != nil {
		return nil, err
	}

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, decoded)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
