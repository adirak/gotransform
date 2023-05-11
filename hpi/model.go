package hpi

type Transform struct {

	// Input data
	Input interface{} `structs:"input" json:"input" bson:"input"`

	// Output data if you want to merge function
	Output interface{} `structs:"output" json:"output" bson:"output"`

	// Shared databus between transform function
	InputBus  *JsonPath `structs:"inputBus" json:"inputBus" bson:"inputBus"`
	OutputBus *JsonPath `structs:"outputBus" json:"outputBus" bson:"outputBus"`

	// Flag to ignore converting output databus to map
	IgnoreReturn bool `structs:"ignoreReturn" json:"ignoreReturn" bson:"ignoreReturn"`
}

// TransformConfig is data model of transform configuration
type TransformConfig struct {
	Input     []string `structs:"input" json:"input" bson:"input"`
	Output    []string `structs:"output" json:"output" bson:"output"`
	Type      string   `structs:"type" json:"type" bson:"type"`
	Split     string   `structs:"split" json:"split" bson:"split"`
	Combine   string   `structs:"combine" json:"combine" bson:"combine"`
	Compare   string   `structs:"compare" json:"compare" bson:"compare"`
	Format    string   `structs:"format" json:"format" bson:"format"`
	From      string   `structs:"from" json:"from" bson:"from"`
	To        string   `structs:"to" json:"to" bson:"to"`
	Position  string   `structs:"position" json:"position" bson:"position"`
	Num       int64    `structs:"num" json:"num" bson:"num"`
	Step      int64    `structs:"step" json:"step" bson:"step"`
	ValueType string   `structs:"valueType" json:"valueType" bson:"valueType"`
	Value     string   `structs:"value" json:"value" bson:"value"`
	Index     string   `structs:"index" json:"index" bson:"index"`
	Fields    string   `structs:"fields" json:"fields" bson:"fields"`
	Fields2   string   `structs:"fields2" json:"fields2" bson:"fields2"`
	Operator  string   `structs:"operator" json:"operator" bson:"operator"`

	RegExp string `structs:"regexp" json:"regexp" bson:"regexp"`
	Prefix string `structs:"prefix" json:"prefix" bson:"prefix"`
	Suffix string `structs:"suffix" json:"suffix" bson:"suffix"`
	Center string `structs:"center" json:"center" bson:"center"`

	// Number Format
	IntDigit          int64  `structs:"intDigit" json:"intDigit" bson:"intDigit"`
	Padding           string `structs:"padding" json:"padding" bson:"padding"`
	DecDigit          int64  `structs:"decDigit" json:"decDigit" bson:"decDigit"`
	ThousandSeparator string `structs:"thousandSeparator" json:"thousandSeparator" bson:"thousandSeparator"`
	DecimalSeparator  string `structs:"decimalSeparator" json:"decimalSeparator" bson:"decimalSeparator"`
}

// TransformConfigFile is data of transfrom config from file
type TransformConfigFile struct {
	File   string                 `structs:"file" json:"file" bson:"file"`
	Config map[string]interface{} `structs:"config" json:"config" bson:"config"`
}
