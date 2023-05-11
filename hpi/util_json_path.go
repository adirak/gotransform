package hpi

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// JsonPath is data model of root of path
type JsonPath struct {
	Root *map[string]interface{} `structs:"root" json:"root" bson:"root"`
}

// Create new json path
func NewJsonPath() *JsonPath {
	jp := JsonPath{}
	jp.Root = &map[string]interface{}{}
	return &jp
}

// Create new json path
func NewJsonPathWithRoot(root *map[string]interface{}) (*JsonPath, error) {
	jp := JsonPath{}
	if root == nil {
		return nil, errors.New("root map is null")
	}
	jp.Root = root
	return &jp, nil
}

// ToMap is function to get root map
func (c *JsonPath) ToMap() map[string]interface{} {
	return c.getRoot()
}

// Value is function to get value by path
func (c *JsonPath) Value(path string) interface{} {

	paths := c.splitPath(path)
	if len(paths) == 0 {
		return c.getRoot()
	}

	// value
	var curVal interface{}
	curVal = c.getRoot()

	// loop to get
	for _, spath := range paths {
		spath, idx, err := c.getArrayIndex(spath)
		if err != nil {
			panic(err)
		}
		if idx < 0 {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				panic(err)
			}
			curVal = val
			continue
		} else {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				panic(err)
			}
			val2, err := c.getValueFromArray(val, idx)
			if err != nil {
				panic(err)
			}
			curVal = val2
			continue
		}
	}

	return curVal
}

// Value is function to get value by path
func (c *JsonPath) String(defaultValue, path string) (str string) {

	// Add function to recovery if logic has panic error
	defer func() {
		if r := recover(); r != nil {
			str = defaultValue
		}
	}()
	// --------------------------------------------------

	value := c.Value(path)
	if value == nil {
		value = defaultValue
	}

	str, ok := value.(string)
	if !ok {
		str = fmt.Sprintf("%v", value)
	}

	return
}

// Delete is function to delete value by path
func (c *JsonPath) Delete(path string) error {

	paths := c.splitPath(path)
	if len(paths) == 0 {
		return nil
	}

	// value
	var curVal interface{}
	curVal = c.getRoot()
	delIdx := len(paths) - 1

	// loop to get
	for i, spath := range paths {

		spath, idx, err := c.getArrayIndex(spath)
		if err != nil {
			return err
		}

		// Delete Action
		if i == delIdx {

			if idx < 0 {
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				delete(mapObj, spath)
				return nil
			} else {

				val, err := c.getValueFrom(curVal, spath)
				if err != nil {
					return err
				}
				arr, err := c.removeValueFromArray(val, idx)
				if err != nil {
					return err
				}
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = arr
				return nil
			}

		}

		// Read Action
		if idx < 0 {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			curVal = val
			continue
		} else {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			val2, err := c.getValueFromArray(val, idx)
			if err != nil {
				return err
			}
			curVal = val2
			continue
		}
	}

	return nil
}

// Value is function to get value by path
func (c *JsonPath) Set(path string, value interface{}) error {

	paths := c.splitPath(path)
	if len(paths) == 0 {
		return nil
	}

	// value
	var curVal interface{}
	curVal = c.getRoot()
	setIdx := len(paths) - 1

	// loop to get
	for i, spath := range paths {

		spath, idx, err := c.getArrayIndex(spath)
		if err != nil {
			return err
		}

		// Delete Action
		if i == setIdx {

			if idx < 0 {
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = value
				return nil
			} else {

				val, err := c.getValueFrom(curVal, spath)
				if err != nil {
					return err
				}

				if val == nil {
					val = []interface{}{}
					mapObj, ok := curVal.(map[string]interface{})
					if !ok {
						return errors.New("data at spath=" + spath + " is not map")
					}
					mapObj[spath] = val
				}

				arr, err := c.setValueToArray(val, idx, value)
				if err != nil {
					return err
				}

				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = arr
				return nil
			}

		}

		// Read Action
		if idx < 0 {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			if val == nil {
				val = map[string]interface{}{}
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = val
			}
			curVal = val
			continue
		} else {
			val, err := c.getValueFrom(curVal, spath)
			if err != nil {
				return err
			}
			if val == nil {
				val = []interface{}{}
				mapObj, ok := curVal.(map[string]interface{})
				if !ok {
					return errors.New("data at spath=" + spath + " is not map")
				}
				mapObj[spath] = val
			}
			val2, err := c.getValueFromArray(val, idx)
			if err != nil {
				msg := err.Error()
				msg = strings.ToLower(msg)
				if strings.Contains(msg, "index out of bound") {
					val2 = map[string]interface{}{}
					val, err = c.setValueToArray(val, idx, val2)
					if err != nil {
						return err
					}
					mapObj, ok := curVal.(map[string]interface{})
					if !ok {
						return errors.New("data at spath=" + spath + " is not map")
					}
					mapObj[spath] = val
				}
			}
			curVal = val2
			continue
		}
	}

	return nil

}

// getValueFrom is function to get value from current data
func (c *JsonPath) getValueFrom(curVal interface{}, spath string) (interface{}, error) {
	if spath == "" {
		return curVal, nil
	}

	// Curvalue is map
	mapObj, ok := curVal.(map[string]interface{})
	if !ok {
		return nil, errors.New("value at name=" + spath + " is not map")
	}
	return mapObj[spath], nil
}

// getValueFromArray is function to get value from array
func (c *JsonPath) getValueFromArray(curVal interface{}, index int) (interface{}, error) {

	if index < 0 {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	// Curvalue is map
	arr, ok := curVal.([]interface{})
	if !ok {
		return nil, errors.New("value at is not array")
	}

	l := len(arr)
	if index >= l {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}
	return arr[index], nil
}

// getValueFromArray is function to get value from array
func (c *JsonPath) removeValueFromArray(curVal interface{}, index int) (interface{}, error) {

	if index < 0 {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	// Curvalue is map
	arr, ok := curVal.([]interface{})
	if !ok {
		return nil, errors.New("value at is not array")
	}

	l := len(arr)
	if index >= l {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	arr = append(arr[:index], arr[index+1:]...)

	return arr, nil
}

// setValueToArray is function to set value from array
func (c *JsonPath) setValueToArray(curVal interface{}, index int, value interface{}) (interface{}, error) {

	if index < 0 {
		return nil, fmt.Errorf("array index out of bound, index=%v", index)
	}

	// Curvalue is map
	arr, ok := curVal.([]interface{})
	if !ok {
		return nil, errors.New("value at is not array")
	}

	l := len(arr)
	if index >= l {
		for i := l - 1; i <= index; i++ {
			if i >= l {
				arr = append(arr, nil)
			}
		}
	}

	// set value by index
	arr[index] = value
	return arr, nil
}

// splitPath is function to split path by type
// func (c *JsonPath) splitPath(path string) []string {
// 	arr := strings.Split(path, ".")
// 	return arr
// }

// splitPath is function to split path by type
func (c *JsonPath) splitPath(path string) (paths []string) {

	// init paths
	paths = []string{}
	if path == "" {
		return
	}

	var sb strings.Builder
	iMax := len(path) - 1
	oQuote := false
	for i, r := range path {

		// Check dot operator
		dot := r == '.'
		nQuote := (r == '\'' || r == '"')
		if nQuote {
			oQuote = !oQuote
		}

		// It's changed
		if dot && !oQuote {
			if sb.Len() > 0 {
				spath := sb.String()
				paths = append(paths, spath)
			}
			sb.Reset()
		}

		// Keep data
		if oQuote && !nQuote {
			sb.WriteRune(r)
		} else if !dot && !nQuote {
			sb.WriteRune(r)
		}

		// End
		if i == iMax {
			spath := sb.String()
			paths = append(paths, spath)
		}
	}

	return paths
}

// getRoot is function to get root value
func (c *JsonPath) getRoot() map[string]interface{} {
	root := c.Root
	return *root
}

// getArrayIndex is function to get array index
func (c *JsonPath) getArrayIndex(spath string) (string, int, error) {
	if strings.Contains(spath, "[") && strings.HasSuffix(spath, "]") {
		s := strings.Index(spath, "[") + 1
		e := len(spath) - 1
		ssp := spath[s:e]
		val, err := strconv.ParseInt(ssp, 10, 64)
		if err != nil {
			return spath, -1, err
		}
		npath := spath[:s-1]
		return npath, int(val), nil
	}
	return spath, -1, nil
}
