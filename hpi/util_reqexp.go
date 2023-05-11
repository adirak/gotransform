package hpi

import (
	"fmt"
	"regexp"
)

// FindValuesByRegExp is function find values by reqular expression
func (c *Transform) FindValuesByRegExp(obj interface{}, strRegExp string) ([]string, error) {

	r, err := regexp.Compile(strRegExp)
	if err != nil {
		return nil, err
	}

	// Call Recursive to find values
	return c.doFindValuesByRegExp(obj, r), nil

}

// doFindValuesByRegExp is function to find values by reqular expression object and recursive to find
func (c *Transform) doFindValuesByRegExp(obj interface{}, r *regexp.Regexp) []string {

	// String
	strObj, ok := obj.(string)
	if ok {
		return r.FindAllString(strObj, -1)
	}

	// Map
	mapObj, ok := obj.(map[string]interface{})
	if ok {
		values := make([]string, 0)
		for _, value := range mapObj {
			results := c.doFindValuesByRegExp(value, r)
			if len(results) > 0 {
				values = append(values, results...)
			}
		}
		return values
	}

	// Array
	arrObj, ok := obj.([]interface{})
	if ok {
		values := make([]string, 0)
		for _, value := range arrObj {
			results := c.doFindValuesByRegExp(value, r)
			if len(results) > 0 {
				values = append(values, results...)
			}
		}
		return values
	}

	// Other object
	strObj = fmt.Sprintf("%v", obj)
	return r.FindAllString(strObj, -1)
}
