package hpi

import (
	"math/rand"
	"strings"
	"time"
)

// HpiRandomInteger is function to random integer for transform node
func (c *Transform) HpiRandomInteger(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_RandomInteger, false, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	value := config.Value
	max, err := c.ToInt64(value)
	if err != nil {
		return nil, err
	}

	// Set seed
	rand.Seed(time.Now().Unix())

	// Random Data
	ran := rand.Intn(int(max))

	oName := config.Output[0]
	c.SetOutputData(oName, ran)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}

// HpiRandomAlphabet is function to random alphabet for transform node
func (c *Transform) HpiRandomAlphabet(tConfig interface{}) (output map[string]interface{}, err error) {

	// Get transform config
	config, err := c.GetBasicConfig(tConfig)
	if err != nil {
		return nil, err
	}

	// Validate TransformType and input/output data
	err = c.Validate(config, HPI_TType_RandomAlphabet, false, true)
	if err != nil {
		return nil, err
	}

	// Logic of function
	// =======================================

	valueType := config.ValueType
	value := config.Value
	max, err := c.ToInt64(value)
	if err != nil {
		return nil, err
	}

	// init
	m := int(max)
	list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	l := len(list)
	var sb strings.Builder

	// Set seed
	rand.Seed(time.Now().Unix())

	if strings.EqualFold(valueType, "string") {
		for i := 0; i < m; i++ {
			ran := rand.Intn(l)
			sb.WriteString(list[ran])
		}
	} else {
		for i := 0; i < m; i++ {
			ran := rand.Intn(10)
			sb.WriteString(list[ran])
		}
	}

	// Random Data
	result := sb.String()

	oName := config.Output[0]
	c.SetOutputData(oName, result)

	// =======================================

	// No Return Value
	if c.IgnoreReturn {
		return nil, nil
	}
	// Return
	return c.GetOutput(), nil

}
