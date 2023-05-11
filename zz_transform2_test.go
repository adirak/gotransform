package main

import (
	"testing"

	"github.com/adirak/gotransform/hpi"
)

func TestReplaceString(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/basic/replace_string.json")
	if err != nil {
		t.Error(err)
	}

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		t.Error("data is not map")
	}

	dataBus, _ := mapObj["input"].(map[string]interface{})
	transform, _ := mapObj["transform"].(map[string]interface{})

	res, err := hpi.ProcessTransform(dataBus, transform)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	hpi.WriteJsonFile("./conf_result/replace_string.json", res)
}

func TestNumberFormat(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/basic/number_format.json")
	if err != nil {
		t.Error(err)
	}

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		t.Error("data is not map")
	}

	dataBus, _ := mapObj["input"].(map[string]interface{})
	transform, _ := mapObj["transform"].(map[string]interface{})

	res, err := hpi.ProcessTransform(dataBus, transform)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
	hpi.WriteJsonFile("./conf_result/number_format.json", res)
}

func TestDmaToSSQ(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/ssq_dma/dma_to_ssq.json")
	if err != nil {
		t.Error(err)
	}

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		t.Error("data is not map")
	}

	input, _ := mapObj["input"].(map[string]interface{})
	output := map[string]interface{}{}
	transform, _ := mapObj["transform"].(map[string]interface{})

	res, err := hpi.TransformData(input, output, transform)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
	hpi.WriteJsonFile("./conf_result/dma_to_ssq.json", res)
}
