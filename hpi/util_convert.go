package hpi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ToInteger is function to convert number to int64
func (c *Transform) ToInt64(value interface{}) (int64, error) {

	if value == nil {
		return 0, nil
	}

	switch num := value.(type) {
	case int:
		return int64(num), nil
	case int64:
		return num, nil
	case int32:
		return int64(num), nil
	case int16:
		return int64(num), nil
	case float64:
		return int64(num), nil
	case float32:
		return int64(num), nil
	default:
		str := fmt.Sprintf("%v", num)
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			offset := strings.Index(str, ".")
			nStr := str[:offset]
			i, err = strconv.ParseInt(nStr, 10, 64)
			if err != nil {
				return 0, err
			}
		}
		return i, nil
	}

}

// ToDecimal is function to convert number to float64
func (c *Transform) ToFloat64(value interface{}) (float64, error) {

	if value == nil {
		return 0, nil
	}

	switch num := value.(type) {
	case int:
		return float64(num), nil
	case int64:
		return float64(num), nil
	case int32:
		return float64(num), nil
	case int16:
		return float64(num), nil
	case float64:
		return num, nil
	case float32:
		return float64(num), nil
	default:
		str := fmt.Sprintf("%v", num)
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	}

}

// ToBoolean is function to convert object to boolean
func (c *Transform) ToBoolean(value interface{}) bool {

	// return false
	if value == nil {
		return false
	}

	switch num := value.(type) {
	case int:
		return num != 0
	case int64:
		return num != 0
	case int32:
		return num != 0
	case int16:
		return num != 0
	case float64:
		return num != 0
	case float32:
		return num != 0
	default:

		str := fmt.Sprintf("%v", num)
		str = strings.ToLower(str)
		str = strings.TrimSpace(str)
		if strings.Contains(str, "true") {
			return true
		} else if strings.Contains(str, "false") {
			return false
		} else if strings.EqualFold(str, "1") {
			return true
		} else if strings.EqualFold(str, "0") {
			return false
		} else if len(str) > 0 {
			return true
		}
		return false

	}

}

// ConvertValueByType is function to convert value by type
func (c *Transform) ConvertValueByType(value interface{}, strType string, defaultValue string) (interface{}, error) {

	// set to default
	if value == nil {
		value = defaultValue
	}

	// error if null
	if value == nil {
		return nil, fmt.Errorf("value is null")
	}

	// String type
	if strings.EqualFold(strType, "string") {
		str, ok := value.(string)
		if ok {
			if str == "" {
				return defaultValue, nil
			}
			return str, nil
		} else {
			return fmt.Sprintf("%v", value), nil
		}
	}

	// Integer type
	if strings.EqualFold(strType, "integer") {
		intValue, err := c.ToInt64(value)
		return intValue, err
	}

	// Decimal type
	if strings.EqualFold(strType, "decimal") {
		floatValue, err := c.ToFloat64(value)
		return floatValue, err
	}

	// boolean type
	if strings.EqualFold(strType, "boolean") {
		boolValue := c.ToBoolean(value)
		return boolValue, nil
	}

	return value, nil

}

// GetNumberInString function to get number in string
func (c *Transform) GetNumberInString(obj interface{}) (float64, error) {

	if obj == nil {
		return 0, fmt.Errorf("value is null")
	}

	str := fmt.Sprintf("%v", obj)

	var sb strings.Builder
	hasDecimal := false
	for _, rune := range str {
		if rune >= 48 && rune <= 57 {
			sb.WriteRune(rune)
		} else if rune == 46 && !hasDecimal {
			sb.WriteRune(rune)
			hasDecimal = true
		}
	}
	if strings.HasSuffix(sb.String(), ".") {
		sb.WriteRune('0')
	}

	// Parse Float
	num, err := strconv.ParseFloat(sb.String(), 64)

	return num, err
}

// ToString is function to convert object to string
func (c *Transform) ToString(value interface{}) (string, error) {

	// Empty String
	if value == nil {
		return "", nil
	}

	str, ok := value.(string)
	if ok {
		return str, nil
	}

	// map
	mapObj, ok := value.(map[string]interface{})
	if ok {
		bData, err := json.Marshal(mapObj)
		if err != nil {
			return "", err
		}
		str = string(bData)
		return str, nil
	}

	// array
	arrObj, ok := value.([]interface{})
	if ok {
		bData, err := json.Marshal(arrObj)
		if err != nil {
			return "", err
		}
		str = string(bData)
		return str, nil
	}

	// Other type
	str = fmt.Sprintf("%v", value)
	return str, nil

}
