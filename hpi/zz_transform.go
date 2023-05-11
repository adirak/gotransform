package hpi

import (
	"fmt"
)

// Transform Config buffer
var hpiTfConfBuffer []TransformConfigFile

const size_buff_min = 100
const size_buff_max = 200

// Initial buffer data
func init() {
	hpiTfConfBuffer = make([]TransformConfigFile, 0)
}

// readConfigFile is function to read transform config from file
func readConfigFile(path string) (config map[string]interface{}, err error) {

	// Read from buffer
	for _, tfConff := range hpiTfConfBuffer {
		filepath := tfConff.File
		if path == filepath {
			config = tfConff.Config
			return config, nil
		}
	}

	// Read config from file
	config, err = JsonMapFromFile(path)
	if err != nil {
		return nil, err
	}

	// Keep config map
	tfConff := TransformConfigFile{File: path, Config: config}
	hpiTfConfBuffer = append(hpiTfConfBuffer, tfConff)

	// Clearn config file
	if len(hpiTfConfBuffer) > size_buff_max {
		diff := size_buff_max - size_buff_min
		hpiTfConfBuffer = hpiTfConfBuffer[diff:]
	}

	// return value
	return config, err
}

// Transform is function to transfrom data in dataBus by transform config path
func TransformWithFile(dataBus map[string]interface{}, transformFile string) (res map[string]interface{}, err error) {
	return TransformDataWithFile(dataBus, dataBus, transformFile)
}

// ProcessTransform is function to transfrom data in dataBus by transform configuration
func ProcessTransform(dataBus map[string]interface{}, transform map[string]interface{}) (res map[string]interface{}, err error) {
	return TransformData(dataBus, dataBus, transform)
}

// TransformDataWithFile is function to transfrom data by transform config path
func TransformDataWithFile(input map[string]interface{}, output map[string]interface{}, transformFile string) (res map[string]interface{}, err error) {
	transformConfig, err := readConfigFile(transformFile)
	if err != nil {
		return nil, err
	}
	return TransformData(input, output, transformConfig)
}

// TransformData is function to transfrom data by transform configuration
func TransformData(input map[string]interface{}, output map[string]interface{}, transform map[string]interface{}) (res map[string]interface{}, err error) {

	// Add function to recovery if logic has panic error
	defer func() {
		if r := recover(); r != nil {
			res = input
			err = fmt.Errorf("internal error in transform node, %v", r)
		}
	}()
	// --------------------------------------------------

	// Create Transfrom Instance
	trnsf := Transform{Input: input, Output: output, IgnoreReturn: true}

	// Transform Process
	processObj := transform["process"]
	if processObj != nil {
		processes, ok := processObj.([]interface{})
		if ok {

			// If transform process is empty
			// return default output
			if len(processes) <= 0 {
				return output, nil
			}

			// Loop to run transform
			for _, tConfig := range processes {
				err := process(&trnsf, tConfig)
				if err != nil {
					return nil, err
				}
			}

		} else {

			// If it's empty or null
			// return default output
			if len(processes) <= 0 {
				return output, nil
			}

		}

	}

	// Result Data
	res = trnsf.GetOutput()
	return res, err
}
