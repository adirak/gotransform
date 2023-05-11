package hpi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// GetBasicConfig is function to get basic data of transfrom
func (c *Transform) GetBasicConfig(transform interface{}) (*TransformConfig, error) {

	if transform != nil {

		// Map json transform
		mapTransform, ok4 := transform.(map[string]interface{})
		if ok4 {

			// Convert to Json Byte Data
			bData, err := json.Marshal(mapTransform)
			if err == nil {
				tConf := TransformConfig{}
				err = json.Unmarshal(bData, &tConf)
				if err == nil {
					return &tConf, nil
				}
			}
		}

		// Check config is *TransformConfig
		tfConfig2, ok2 := transform.(*TransformConfig)
		if ok2 {
			return tfConfig2, nil
		}

		// Check config is TransformConfig
		tfConfig1, ok1 := transform.(TransformConfig)
		if ok1 {
			return &tfConfig1, nil
		}

		// String json transform
		strTransform, ok3 := transform.(string)
		if ok3 {
			tConf := TransformConfig{}
			err := json.Unmarshal([]byte(strTransform), &tConf)
			if err == nil {
				return &tConf, err
			}
		}

		return nil, errors.New("transform config data format is error")

	} else {
		return nil, errors.New("transform config is null")
	}

}

// getInputBytes is function to get transform binary data
func (c *Transform) getBytesData(data interface{}) ([]byte, error) {
	if data != nil {

		byteInput, ok := data.([]byte)
		if ok {
			return byteInput, nil
		}

		// Input String
		strInput, ok1 := data.(string)
		if ok1 {
			return []byte(strInput), nil
		}

		// Input Map
		inputMap, ok2 := data.(map[string]interface{})
		if ok2 {
			jsonString, err := json.Marshal(inputMap)
			return jsonString, err
		}

		return nil, errors.New("data format is error")

	} else {
		return nil, errors.New("data is null")
	}
}

// GetInputBus is function to create inputBus if it is null and return it
func (c *Transform) GetInputBus() *JsonPath {

	// Check null data
	if c.InputBus == nil {

		// init nil data
		if c.Input == nil {
			c.Input = map[string]interface{}{}
		}

		// Convert from map
		mapObj, ok := c.Input.(map[string]interface{})
		if ok {
			jp, err := NewJsonPathWithRoot(&mapObj)
			if err == nil {
				c.InputBus = jp
				return jp
			}
		}

		bData, err := c.getBytesData(c.Input)

		// get bytes error
		if err != nil {
			fmt.Println("*** input data is error : " + err.Error())
			bData = []byte("{}")
		}

		// Convert from byte array
		mapObj = map[string]interface{}{}
		err = json.Unmarshal(bData, &mapObj)
		if err != nil {
			panic(err)
		}
		jp, err := NewJsonPathWithRoot(&mapObj)
		if err != nil {
			panic(err)
		}
		c.InputBus = jp
		return jp
	}

	return c.InputBus
}

// GetOutputBus is function to create outbutBus if it is null and return it
func (c *Transform) GetOutputBus() *JsonPath {

	// Check null data
	if c.OutputBus == nil {

		// init nil data
		if c.Output == nil {
			c.Output = map[string]interface{}{}
		}

		// Convert from map
		mapObj, ok := c.Output.(map[string]interface{})
		if ok {
			jp, err := NewJsonPathWithRoot(&mapObj)
			if err == nil {
				c.OutputBus = jp
				return jp
			}
		}

		bData, err := c.getBytesData(c.Output)
		// get bytes error
		if err != nil {
			fmt.Println("*** input data is error : " + err.Error())
			bData = []byte("{}")
		}

		// Convert from byte array
		mapObj = map[string]interface{}{}
		err = json.Unmarshal(bData, &mapObj)
		if err != nil {
			panic(err)
		}
		jp, err := NewJsonPathWithRoot(&mapObj)
		if err != nil {
			panic(err)
		}
		c.OutputBus = jp
		return jp
	}

	return c.OutputBus
}

// SetOutputData is function to set output data by path
func (c *Transform) SetOutputData(path string, value interface{}) {

	if strings.HasPrefix(path, "_tmp") && strings.HasSuffix(path, "_") {

		// This is temporary variable
		inputBus := c.GetInputBus()
		inputBus.Set(path, value)

	} else {

		// Set result to output
		outputBus := c.GetOutputBus()
		outputBus.Set(path, value)
	}

}

// DeleteOutputData is function to delete output field
func (c *Transform) DeleteOutputData(path string) {

	// Set result to output
	outputBus := c.GetOutputBus()
	outputBus.Delete(path)

}

// GetOutput is function to convert outputBut to map data
func (c *Transform) GetOutput() map[string]interface{} {

	// initial data bus
	if c.OutputBus == nil {
		c.GetOutputBus()
	}

	// Convert but to map
	if c.OutputBus != nil {
		jparser := *c.OutputBus
		return jparser.ToMap()
	}

	// No data return empty
	emptyMap := make(map[string]interface{})
	return emptyMap
}

// ValidateTransformType is function to validate transform type
func (c *Transform) validateType(trnsConf *TransformConfig, expectedType string) error {

	if trnsConf != nil {
		if strings.EqualFold(trnsConf.Type, expectedType) {
			return nil
		} else {
			return errors.New("transform type is not corrected by type=" + expectedType)
		}
	} else {
		return errors.New("transform is null")
	}

}

// Validate is function to validate transform type and input output data
func (c *Transform) Validate(trnsConf *TransformConfig, expectedType string, needInput bool, needOutput bool) error {

	// Validate transform type
	err := c.validateType(trnsConf, expectedType)
	if err != nil {
		return err
	}

	// Validate input data
	if needInput {
		if c.Input == nil {
			return errors.New("input is null")
		}
		if trnsConf.Input == nil || len(trnsConf.Input) <= 0 {
			return errors.New("input field is empty")
		}
	}

	// Validate output data
	if needOutput {
		if trnsConf.Output == nil || len(trnsConf.Output) <= 0 {
			return errors.New("output field is empty")
		}
	}

	// No error
	return nil
}
