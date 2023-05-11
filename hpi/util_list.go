package hpi

// AppendToList is function to append data by first or last index
func (c *Transform) AppendToList(list []interface{}, value interface{}, toFirst bool) []interface{} {

	if toFirst {

		index := 0
		if len(list) == index {
			list = append(list, value)
			return list
		} else {
			// Append at first
			list = append([]interface{}{value}, list...)
		}

	} else {

		// To Last Index
		list = append(list, value)
	}

	return list
}

// JoinRecord is function to join record by some field and condition
func (c *Transform) JoinRecord(fields1 []string, fields2 []string, data1 map[string]interface{}, data2 map[string]interface{}, isAnd bool) (data3 map[string]interface{}, success bool) {

	// Prepare join
	count := 0
	size := len(fields1)
	if size > len(fields2) {
		size = len(fields2)
	}

	// Loop checking
	for i := 0; i < size; i++ {
		name1 := fields1[i]
		name2 := fields2[i]
		value1 := data1[name1]
		value2 := data2[name2]
		if value1 == value2 {
			count++
		}
	}

	// Check condition to join
	canJoin := false
	if isAnd && count >= size {
		canJoin = true
	} else if !isAnd && count > 0 {
		canJoin = true
	}

	if canJoin {

		// Merge Data
		data3 = make(map[string]interface{})

		// Add data2
		for key, value := range data2 {
			data3[key] = value
		}

		// Add data1
		for key, value := range data1 {
			data3[key] = value
		}

		// Return success
		return data3, true
	}

	return
}

// CheckFilterRecord is function to check condition for filtering record
func (c *Transform) CheckFilterRecord(fields []string, values []interface{}, data map[string]interface{}, isAnd bool, compareType string) (success bool) {

	// Prepare join
	count := 0
	size := len(fields)
	if size > len(values) {
		size = len(values)
	}

	// Loop checking
	for i := 0; i < size; i++ {

		field := fields[i]
		valueA := data[field]
		valueB := values[i]

		if c.IsEqual(valueA, valueB, compareType) {
			count++
		}
	}

	// Checking match case
	success = false
	if isAnd && count >= size && count > 0 {
		success = true
	} else if !isAnd && count > 0 {
		success = true
	}

	return success
}

// RemoveIndexInList is function to remove index in list
func (c *Transform) RemoveIndexInList(list []interface{}, index int) []interface{} {
	return append(list[:index], list[index+1:]...)
}

// ReverseString is function to reverse string value
func (c *Transform) ReverseString(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}
