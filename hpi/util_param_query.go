package hpi

import (
	"fmt"
	"strings"
)

// GetIndexValue is function to get index value in array
func (c *Transform) GetIndexValue(indexKey string, maxIndex int) int {

	if strings.HasPrefix(indexKey, "$") {

		// Get index from input fields
		name := strings.Replace(indexKey, "$", "", 1)
		inBus := c.GetInputBus()
		value := inBus.Value(name)

		index, err := c.ToInt64(value)
		if err != nil {
			return -1
		}
		return int(index)

	} else if strings.Contains(indexKey, "last") {

		// Get last index
		return maxIndex

	} else if strings.Contains(indexKey, "first") {

		// Get first index
		return 0

	} else {

		index, err := c.ToInt64(indexKey)
		if err != nil {
			return -1
		}
		return int(index)

	}
}

// ToFieldArray is function to convert fields that it's saparated by comma to array
func (c *Transform) ToArrayFields(fields string) []string {

	arr := strings.Split(fields, ",")
	for i, str := range arr {
		str = strings.TrimSpace(str)
		arr[i] = str
	}

	return arr
}

// ToArrayValues is function to convert value and separated by comma to array value
func (c *Transform) ToArrayValues(values string) (valArr []interface{}) {

	arr := strings.Split(values, ",")

	valArr = make([]interface{}, len(arr))

	for i, str := range arr {

		str = strings.TrimSpace(str)

		// Reference to value in input
		if strings.HasPrefix(str, "$") {

			// Get value from input fields
			name := strings.Replace(str, "$", "", 1)
			inBus := c.GetInputBus()
			value := inBus.Value(name)
			valArr[i] = value

		} else {

			valArr[i] = str

		}

	}

	return valArr
}

// ToValueParam is function to convert value from sting and $param
func (c *Transform) ToValueParam(inValue interface{}) (value interface{}) {

	strValue, ok := inValue.(string)
	if ok {
		arr := c.ToArrayValues(strValue)
		if len(arr) > 0 {
			return arr[0]
		} else {
			return nil
		}
	} else {
		strValue = fmt.Sprintf("%v", inValue)
		arr := c.ToArrayValues(strValue)
		if len(arr) > 0 {
			return arr[0]
		} else {
			return nil
		}
	}

}
