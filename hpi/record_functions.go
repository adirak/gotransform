package hpi

import (
	"fmt"
	"strings"
)

// HpiCopyToList is function to copy value from field and put it to all records in list
func (c *Transform) HpiUpdateRecord(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_UpdateRecord, true, true)
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

	// Loop the input
	inputBus := c.GetInputBus()
	inName := config.Input[0]

	// Input record
	record := inputBus.Value(inName)

	outputBus := c.GetOutputBus()
	oName := config.Output[0]
	listObj := outputBus.Value(oName)

	// Make empty list
	if listObj == nil {
		listObj = make([]interface{}, 0)
	}

	// Validate List of Map
	list, err := c.ValidateList(listObj)
	if err != nil {
		return nil, err
	}

	indexKey := config.Index
	index := c.GetIndexValue(indexKey, len(list)-1)

	// Loop to copy to all record
	if index >= 0 && len(list) > 0 {
		list[index] = record
	}

	// Keep result
	c.SetOutputData(oName, list)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiDeleteRecord is function to delete some record in list
func (c *Transform) HpiDeleteRecord(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_DeleteRecord, true, true)
	if err != nil {
		return nil, err
	}

	// Validate null input data
	err = c.ValidateNullData(config, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Prepare input list
	inName := config.Input[0]
	inputBus := c.GetInputBus()
	listObj := inputBus.Value(inName)

	// Validate List of Map
	list1, err := c.ValidateList(listObj)
	if err != nil {
		return nil, err
	}

	indexKey := config.Index
	index := c.GetIndexValue(indexKey, len(list1)-1)

	// Remove Index
	list1 = c.RemoveIndexInList(list1, index)

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, list1)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiFilterRecord is function to filter some match record condition in list
func (c *Transform) HpiFilterRecord(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_FilterRecord, true, true)
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
	inListObj0 := inputBus.Value(inName0)

	// Validate List of Map
	list0, err := c.ValidateListOfMap(inListObj0)
	if err != nil {
		return nil, err
	}

	// Prepare Loop
	list := make([]map[string]interface{}, 0)

	fields := c.ToArrayFields(config.Fields)
	valuses := c.ToArrayValues(config.Value)
	compare := config.Compare
	isAnd := true
	if strings.EqualFold(config.Operator, "OR") {
		isAnd = false
	}

	// Looping to merge
	for _, obj1 := range list0 {
		data, ok1 := obj1.(map[string]interface{})
		if ok1 {
			success := c.CheckFilterRecord(fields, valuses, data, isAnd, compare)
			if success {
				list = append(list, data)
			}
		}
	}

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, list)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiFindOneRecord is function to find match one record condition in list
func (c *Transform) HpiFindOneRecord(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_FindOneRecord, true, true)
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
	inListObj0 := inputBus.Value(inName0)

	// Validate List of Map
	list0, err := c.ValidateList(inListObj0)
	if err != nil {
		return nil, err
	}

	// Prepare Loop
	fields := c.ToArrayFields(config.Fields)
	valuses := c.ToArrayValues(config.Value)
	compare := config.Compare
	isAnd := true
	if strings.EqualFold(config.Operator, "OR") {
		isAnd = false
	}

	// Looping to find
	var result interface{}
	index := -1
	for i, obj1 := range list0 {

		// ignore null data
		if obj1 == nil {
			continue
		}

		// Map Type
		data, ok1 := obj1.(map[string]interface{})
		if ok1 {
			success := c.CheckFilterRecord(fields, valuses, data, isAnd, compare)
			if success {
				index = i
				result = obj1
				break
			}
			continue
		}

		// String type
		str, ok2 := obj1.(string)
		if ok2 {
			if len(valuses) > 0 {
				val0 := valuses[0]
				success := c.IsEqual(str, val0, compare)
				if success {
					index = i
					result = obj1
					break
				}
			}
			continue
		}

		// Other type
		str = fmt.Sprintf("%v", obj1)
		if len(valuses) > 0 {
			val0 := valuses[0]
			success := c.IsEqual(str, val0, compare)
			if success {
				index = i
				result = obj1
				break
			}
		}

	}

	// Put data to output
	oResultName := config.Output[0]
	c.SetOutputData(oResultName, result)
	if len(config.Output) > 1 {
		oResultIndex := config.Output[1]
		c.SetOutputData(oResultIndex, index)
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
