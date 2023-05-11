package hpi

import (
	"strings"
)

// HpiAddToList is function to add value of some field to the list at first or last record
func (c *Transform) HpiAddToList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_AddToList, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	record1 := make(map[string]interface{})

	// Loop the input
	inputBus := c.GetInputBus()
	for _, inPath := range config.Input {
		inName := c.GetNameFromPath(inPath)
		inValue := inputBus.Value(inPath)
		record1[inName] = inValue
	}

	outputBus := c.GetOutputBus()
	oName := config.Output[0]
	listObj := outputBus.Value(oName)

	// Make empty if null
	if listObj == nil {
		listObj = make([]interface{}, 0)
	}

	// Validate List of Map
	list, err := c.ValidateListOfMap(listObj)
	if err != nil {
		return nil, err
	}

	// Append flag
	toFirst := false
	if strings.EqualFold(config.Position, "First") {
		toFirst = true
	}

	list = c.AppendToList(list, record1, toFirst)

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

// HpiAddObjectToList is function to add object to the list at first or last record
func (c *Transform) HpiAddObjectToList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_AddObjectToList, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	// Loop the input
	inputBus := c.GetInputBus()

	// Output
	outputBus := c.GetOutputBus()
	oName := config.Output[0]
	listObj := outputBus.Value(oName)

	// Make empty list
	if listObj == nil {
		listObj = make([]interface{}, 0)
	}

	// Validate List
	list, err := c.ValidateList(listObj)
	if err != nil {
		return nil, err
	}

	// Append flag
	toFirst := false
	if strings.EqualFold(config.Position, "First") {
		toFirst = true
	}

	inLen := len(config.Input)
	if toFirst {

		// Loop to get and set data to list
		for i := inLen - 1; i >= 0; i-- {

			inName := config.Input[i]

			record := inputBus.Value(inName)
			list = c.AppendToList(list, record, toFirst)
		}

	} else {

		// Loop to get and set data to list
		for _, inName := range config.Input {

			record := inputBus.Value(inName)
			list = c.AppendToList(list, record, toFirst)

		}
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

// HpiCopyToList is function to copy value from field and put it to all records in list
func (c *Transform) HpiCopyToList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CopyToList, true, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	record1 := make(map[string]interface{})

	// Loop the input
	inputBus := c.GetInputBus()
	for _, inPath := range config.Input {
		inName := c.GetNameFromPath(inPath)
		inValue := inputBus.Value(inPath)
		record1[inName] = inValue
	}

	outputBus := c.GetOutputBus()
	oName := config.Output[0]
	listObj := outputBus.Value(oName)

	// Validate List of Map
	list, err := c.ValidateListOfMap(listObj)
	if err != nil {
		return nil, err
	}

	// Loop to copy to all record
	for _, recordi := range list {
		mapData := recordi.(map[string]interface{})
		for key, value := range record1 {
			mapData[key] = value
		}
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

// HpiCopyFromList is function to copy field inside of list to target path
func (c *Transform) HpiCopyFromList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CopyFromList, true, true)
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

	// Prepare Input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	inListObj := inputBus.Value(inName)
	fields := config.Fields

	// Validate List of Map
	list, err := c.ValidateListOfMap(inListObj)
	if err != nil {
		return nil, err
	}

	indexKey := config.Index
	index := c.GetIndexValue(indexKey, len(list)-1)

	if index < len(list) {
		mapObj := list[index]
		data := mapObj.(map[string]interface{})

		// Convert to jsonPath
		jsonPath, err := NewJsonPathWithRoot(&data)
		if err != nil {
			return nil, err
		}

		// field mapping
		arrFields := c.ToArrayFields(fields)
		arrLen := len(arrFields)

		// Loop to read output
		for i, path := range config.Output {

			var value interface{}

			if i < arrLen {
				fPath := arrFields[i]
				value = jsonPath.Value(fPath)
				if value == nil {
					value = jsonPath.Value(path)
				}
			} else {
				value = jsonPath.Value(path)
			}

			c.SetOutputData(path, value)
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

// HpiCopyFromListToObject is function to copy data from list to object
func (c *Transform) HpiCopyFromListToObject(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_CopyFromListToObject, true, true)
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

	// Prepare Input
	inputBus := c.GetInputBus()
	inName := config.Input[0]
	inListObj := inputBus.Value(inName)

	// Validate List of Map
	list, err := c.ValidateList(inListObj)
	if err != nil {
		return nil, err
	}

	indexKey := config.Index
	index := c.GetIndexValue(indexKey, len(list)-1)

	if index < len(list) {

		obj := list[index]
		oName := config.Output[0]
		c.SetOutputData(oName, obj)

	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiJoinList is function to join two list by some field and return it to output list
func (c *Transform) HpiJoinList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_JoinList, true, true)
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
	inName1 := config.Input[1]
	inListObj0 := inputBus.Value(inName0)
	inListObj1 := inputBus.Value(inName1)

	// Validate List of Map
	list0, err := c.ValidateListOfMap(inListObj0)
	if err != nil {
		return nil, err
	}

	// Validate List of Map
	list1, err := c.ValidateListOfMap(inListObj1)
	if err != nil {
		return nil, err
	}

	// Prepare Loop
	list := make([]map[string]interface{}, 0)

	fArray1 := c.ToArrayFields(config.Fields)
	fArray2 := c.ToArrayFields(config.Fields2)
	isAnd := true
	if strings.EqualFold(config.Operator, "OR") {
		isAnd = false
	}

	// Looping to merge
	for _, obj1 := range list0 {
		map1, ok1 := obj1.(map[string]interface{})
		for _, obj2 := range list1 {
			map2, ok2 := obj2.(map[string]interface{})
			if ok1 && ok2 {
				record, success := c.JoinRecord(fArray1, fArray2, map1, map2, isAnd)
				if success {

					// Append to result
					list = append(list, record)

					// Remove for 1-1 match case
					//list1 = c.RemoveIndexInList(list1, j)
					//break
				}
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

// HpiListSize is function to count size of the list
func (c *Transform) HpiListSize(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_ListSize, true, true)
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

	// Get size of this list
	size := len(list1)

	// Put data to output
	oName := config.Output[0]
	c.SetOutputData(oName, size)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// SumFieldsInList is function to sum value of field in list
func (c *Transform) HpiSumFieldsInList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SumFieldsInList, true, true)
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

	fields := c.ToArrayFields(config.Fields)
	values := make([]float64, len(fields))

	// Looping to sum value by field
	for _, obj1 := range list0 {
		data, ok1 := obj1.(map[string]interface{})
		if ok1 {
			//jsonPath := c.ToJsonHandler(data)
			jsonPath, err := NewJsonPathWithRoot(&data)
			if err != nil {
				return nil, err
			}
			for j, field := range fields {
				valObj := jsonPath.Value(field)
				value, err := c.ToFloat64(valObj)
				if err == nil {
					values[j] = values[j] + value
				}
			}
		}
	}

	// Put data to output
	for i, oName := range config.Output {
		c.SetOutputData(oName, values[i])
	}

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiSplitListToObject is function to split list to object and set it to output
func (c *Transform) HpiSplitListToObject(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SplitListToObject, true, true)
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
	list, err := c.ValidateList(inListObj0)
	if err != nil {
		return nil, err
	}

	// Loop output
	inLen := len(list)
	for i, oName := range config.Output {
		if i < inLen {
			iValue := list[i]
			c.SetOutputData(oName, iValue)
		} else {
			// No data set to null
			c.SetOutputData(oName, nil)
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

// HpiSearchValueInList is function to search value of each record in the list
func (c *Transform) HpiSearchValueInList(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_SearchValueInList, true, true)
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
	list := make([]interface{}, 0)
	valuses := c.ToArrayValues(config.Value)
	compare := config.Compare

	// Looping to check value
	for _, obj1 := range list0 {

		success := c.CheckValueRecord(obj1, valuses, compare)
		if success {
			list = append(list, obj1)
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
