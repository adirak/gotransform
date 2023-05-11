package hpi

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	uuidP "github.com/google/uuid"
)

// HpiMergeString is function to merge field in input param to output param with combine text
func (c *Transform) HpiMergeString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_MergeString, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// MergeString function
	// Loop to Merge String
	// =======================================
	inputBus := c.GetInputBus()
	var buffer bytes.Buffer
	for i, name := range config.Input {
		if i > 0 {
			buffer.WriteString(config.Combine)
		}
		obj := inputBus.Value(name)
		value := fmt.Sprintf("%v", obj)
		buffer.WriteString(value)
	}

	// Set it to outputBus
	name := config.Output[0]
	c.SetOutputData(name, buffer.String())
	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiSplitString is function to split field value in input param to output param with split text
func (c *Transform) HpiSplitString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SplitString, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Split String function
	// Logic to Split String
	// =======================================
	inputBus := c.GetInputBus()
	name := config.Input[0]
	value := inputBus.String("", name)
	if value != "" {
		splitResults := strings.Split(value, config.Split)
		size := len(config.Output)
		for i, val := range splitResults {
			if i < size {
				oName := config.Output[i]
				c.SetOutputData(oName, val)
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

// HpiSplitStringArray is function to split field value in input param to output param with split text
func (c *Transform) HpiSplitStringArray(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SplitStringArray, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Split String function
	// Logic to Split String
	// =======================================
	inputBus := c.GetInputBus()
	name := config.Input[0]
	value := inputBus.String("", name)
	if value != "" {
		splitResults := strings.Split(value, config.Split)
		oName := config.Output[0]
		c.SetOutputData(oName, splitResults)
	}
	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiReplaceString is function to replace string from old strong to new string
func (c *Transform) HpiReplaceString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_ReplaceString, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Replace String function
	// Logic to replace String
	// =======================================
	inputBus := c.GetInputBus()
	name := config.Input[0]
	value := inputBus.String("", name)
	if value != "" {

		// number of replacement
		num := 1
		if strings.EqualFold(config.Position, "All") {
			num = -1
		} else if strings.EqualFold(config.Position, "Last") {
			num = 0
		}

		from, _ := c.ToValueParam(config.From).(string)
		to, _ := c.ToValueParam(config.To).(string)

		result := value
		if num != 0 {
			result = strings.Replace(value, from, to, num)
		} else {
			var sb strings.Builder
			arr := strings.Split(value, from)
			l := len(arr)
			for i := 0; i < l; i++ {
				if i > 0 {
					if i == l-1 {
						sb.WriteString(to)
					} else {
						sb.WriteString(from)
					}

				}
				s := arr[i]
				sb.WriteString(s)
			}
			result = sb.String()
		}

		oName := config.Output[0]
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

// HpiTrimString is function to remove all white space form string
func (c *Transform) HpiTrimString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_TrimString, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Replace String function
	// Logic to replace String
	// =======================================
	inputBus := c.GetInputBus()
	name := config.Input[0]
	value := inputBus.String("", name)
	if value != "" {

		result := strings.TrimSpace(value)

		oName := config.Output[0]
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

// HpiGenUUID is function to generate uuid by num bit of string
func (c *Transform) HpiGenUUID(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_GenUUID, false, true)
	if err != nil {
		return nil, err
	}

	// Gen UUID
	// Logic to gen uuid function
	// =======================================

	uuidObj := uuidP.New()
	uuid := uuidObj.String()

	oName := config.Output[0]
	c.SetOutputData(oName, uuid)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiGenUniqueID is function to generate uniqueID by num bit of string
func (c *Transform) HpiGenUniqueID(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_GenUniqueId, false, true)
	if err != nil {
		return nil, err
	}

	// Gen UUID
	// Logic to gen uuid function
	// =======================================
	uuid := ""
	num := config.Num
	if num <= 8 {
		dice, _ := rand.Int(rand.Reader, big.NewInt(799999999))
		uuid = fmt.Sprintf("%08x", dice)
	} else if num <= 16 {
		dice, _ := rand.Int(rand.Reader, big.NewInt(5999999999999999999))
		uuid = fmt.Sprintf("%016x", dice)
	} else if num <= 32 {
		uuidObj := uuidP.New()
		uuid := uuidObj.String()
		uuid = strings.Replace(uuid, "-", "", -1)
	} else {

		uuidObj1 := uuidP.New()
		uuid1 := uuidObj1.String()

		// Gen from random
		dice, _ := rand.Int(rand.Reader, big.NewInt(5999999999999999999))
		uuid2 := fmt.Sprintf("%016x", dice)
		dice, _ = rand.Int(rand.Reader, big.NewInt(5999999999999999999))
		uuid3 := fmt.Sprintf("%016x", dice)

		uuid = uuid3 + uuid2 + uuid1
		uuid = strings.Replace(uuid, "-", "", -1)
	}

	oName := config.Output[0]
	c.SetOutputData(oName, uuid)
	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiModifyString is function to modify the string and set it back to output
func (c *Transform) HpiModifyString(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_ModifyString, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, false)
	if err != nil {
		return nil, err
	}

	// Replace String function
	// Logic to replace String
	// =======================================

	// init config
	prefix := config.Prefix
	suffix := config.Suffix
	center := config.Center
	defaltValue := config.Value
	valueType := config.ValueType

	inputBus := c.GetInputBus()
	name := config.Input[0]
	value := inputBus.Value(name)
	oName := config.Output[0]

	if value == nil {

		// Convert type
		result, err := c.ConvertValueByType(value, valueType, defaltValue)
		if err != nil {
			return nil, err
		}

		// Set default value
		c.SetOutputData(oName, result)

	} else {

		// Convert to string
		strValue, ok := value.(string)
		if !ok {
			strValue = fmt.Sprintf("%v", value)
		}

		// Check prefix
		if len(prefix) > 0 && strings.HasPrefix(strValue, prefix) {
			strValue = strings.TrimPrefix(strValue, prefix)
		}

		// Check suffix
		if len(suffix) > 0 && strings.HasSuffix(strValue, suffix) {
			strValue = strings.TrimSuffix(strValue, suffix)
		}

		// check center
		if len(center) > 0 && strings.Contains(strValue, center) {
			strValue = strings.Replace(strValue, center, "", 1)
		}

		// Convert type
		result, err := c.ConvertValueByType(strValue, valueType, defaltValue)
		if err != nil {
			return nil, err
		}

		// Set back to output
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
