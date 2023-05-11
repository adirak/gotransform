package hpi

import (
	"errors"
	"strings"
)

// HpiSsqToDmaDataFormat is function to convert SSQ response data to DMA request data format
func (c *Transform) HpiSsqToDmaDataFormat(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SsqToDma, true, true)
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

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)

	result, resultSuccess, resultMessage, err := c.DoConvertSsqToDma(value)
	if err != nil {
		return nil, err
	}

	// length
	outputs := config.Output
	oLength := len(outputs)

	if oLength > 0 {
		oName := outputs[0]
		c.SetOutputData(oName, result)
	}

	if oLength > 1 && outputs[1] != "" {
		resultSuccessName := outputs[1]
		c.SetOutputData(resultSuccessName, resultSuccess)
	} else {
		c.SetOutputData("resultSuccess", resultSuccess)
	}

	if oLength > 2 && outputs[2] != "" {
		resultMessageName := outputs[2]
		c.SetOutputData(resultMessageName, resultMessage)
	} else {
		c.SetOutputData("resultMessage", resultMessage)
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// doConvertSsqToDma is function convertion logic
func (c *Transform) DoConvertSsqToDma(in interface{}) (resultData map[string]interface{}, resultSuccess bool, resultMessage string, err error) {

	inMap, ok := in.(map[string]interface{})
	if !ok {
		err = errors.New("input data is not map")
		return
	}

	root, err := c.getResultRoot(inMap)
	if err != nil {
		return
	}

	// ResponseCode
	Returncode, _ := root["Returncode"].(string)
	if strings.HasSuffix(Returncode, "0000I") {
		resultSuccess = true
	} else {
		resultSuccess = false
	}

	// ResponseMssage
	Runmessage, _ := root["Runmessage"].(string)
	resultMessage = Runmessage

	// content
	content, ok := root["content"].(map[string]interface{})
	if !ok {
		err = errors.New("content data is not map")
		return
	}

	// init result
	resultData = map[string]interface{}{}
	if len(content) <= 0 {
		return
	}

	// Looping in content
	for key, obj := range content {
		list, ok := obj.([]interface{})
		if !ok {
			resultData[key] = obj
		}
		dataList, err := c.convertTo1LevelList(list)
		if err != nil {
			continue
		}
		resultData["dataList"] = dataList
	}

	return
}

// getResultRoot is function to find result map level
func (c *Transform) getResultRoot(inMap map[string]interface{}) (resultMap map[string]interface{}, err error) {

	Returncode, _ := inMap["Returncode"].(string)

	// Get result root
	if Returncode != "" {
		return inMap, nil
	}

	// Missing Returncode
	for _, obj := range inMap {
		mapObj, ok := obj.(map[string]interface{})
		if !ok {
			continue
		}
		resultMap, err = c.getResultRoot(mapObj)
		if err == nil && resultMap != nil {
			return
		}
	}

	return nil, errors.New("input data is not SSQ format")
}

// convertTo1LevelList is function to convert to one level data list
func (c *Transform) convertTo1LevelList(list []interface{}) (dataList []interface{}, err error) {

	// initial
	dataList = make([]interface{}, 0)
	if len(list) <= 0 {
		return
	}

	for _, obj := range list {

		mapObj, _ := obj.(map[string]interface{})
		rMap, err := c.convertTo1LevelMap(mapObj)
		if err != nil {
			return nil, err
		}

		dataList = append(dataList, rMap)
	}

	return
}

// convertTo1LevelList is function to convert to one level data list
func (c *Transform) convertTo1LevelMap(inMap map[string]interface{}) (out map[string]interface{}, err error) {

	// init
	out = map[string]interface{}{}
	if len(inMap) <= 0 {
		return
	}

	for key, obj := range inMap {

		mapObj, ok := obj.(map[string]interface{})
		if !ok {
			out[key] = obj
			continue
		}

		vals, err := c.convertTo1LevelMap(mapObj)
		if err != nil {
			continue
		}

		// Loop to add key
		for sKey, sObj := range vals {
			out[sKey] = sObj
		}
	}

	return

}

// HpiDmaToSsqDataFormat is function to convert DMA response data to SSQ request data format
func (c *Transform) HpiDmaToSsqDataFormat(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_DmaToSsq, true, true)
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

	// Current time
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	value := inputBus.Value(inName)

	// Config
	service := config.Fields
	version := config.ValueType
	listName := config.Operator

	if service == "" {
		service = "ServiceName"
	}
	if version == "" {
		version = "v2.0"
	}
	if listName == "" {
		listName = "list"
	}

	// convert
	result, err := c.DoConvertDmaToSsq(value, service, version, listName)
	if err != nil {
		return nil, err
	}

	// length
	outputs := config.Output

	oName := outputs[0]
	c.SetOutputData(oName, result)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// DoConvertDmaToSsq is function convertion logic
func (c *Transform) DoConvertDmaToSsq(in interface{}, serviceName, ssqVersion, listName string) (resultData map[string]interface{}, err error) {

	inMap, ok := in.(map[string]interface{})
	if !ok {
		err = errors.New("input data is not map")
		return
	}

	// init ssq data
	serviceName = strings.ToLower(serviceName)
	resultData = map[string]interface{}{}
	var body map[string]interface{}
	if strings.EqualFold(ssqVersion, "v1.0") {
		body = map[string]interface{}{}
		reqName := serviceName + "Request"
		resultData[reqName] = body
		body["version"] = "1.0"
	} else {
		body = resultData
		body["version"] = "2.0"
	}

	// body data
	body["Username"] = ""
	body["ewitoken"] = ""
	body["usergroup"] = ""
	body["workstation"] = ""

	// Content
	content := map[string]interface{}{}

	// Loop to add content
	for key, obj := range inMap {

		arr, ok := obj.([]interface{})
		if !ok {

			content[key], _ = c.ToString(obj)
			continue
		}

		// DataList Type
		dataList, callName, err := c.doConvertSSQDataList(arr, ssqVersion, listName)
		if err != nil {
			return resultData, err
		}
		content[callName] = dataList

	}

	// set content
	body["content"] = content

	return
}

// doConvertSSQDataList is function convertion logic
func (c *Transform) doConvertSSQDataList(arr []interface{}, ssqVersion, listName string) (dataList []interface{}, collectionName string, err error) {

	v2Flag := false
	if strings.EqualFold(ssqVersion, "v2.0") {
		collectionName = listName
		v2Flag = true
	} else {
		collectionName = listName + "Collection"
	}

	// init
	dataList = make([]interface{}, 0)

	for _, obj := range arr {
		mapObj, ok := obj.(map[string]interface{})
		if !ok {
			err = errors.New("record is not map")
			return
		}

		// Convert to string
		for key, obj2 := range mapObj {
			mapObj[key], _ = c.ToString(obj2)
		}

		if v2Flag {
			dataList = append(dataList, mapObj)
			continue
		}

		mapList := map[string]interface{}{}
		mapList[listName] = mapObj

		// Add map list
		dataList = append(dataList, mapList)

	}

	return
}
