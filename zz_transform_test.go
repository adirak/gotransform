package main

import (
	"testing"

	"github.com/adirak/gotransform/hpi"
)

func TestHpiBase64Decode(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/base64_decode.json")
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
	hpi.WriteJsonFile("./conf/out.json", res)
}

func TestDoConvertSsqToDma(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/ssq_dma/ssq_to_dma.json")
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

	hpi.WriteJsonFile("./conf_result/ssq_to_dma.json", res)

}

func TestHpiRandomInteger(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/basic/random_int.json")
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

	hpi.WriteJsonFile("./conf_result/random_int.json", res)
	t.Log(res)
}

func TestHpiRandomAlphabet(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/basic/random_alphabet.json")
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

	hpi.WriteJsonFile("./conf_result/random_alphabet.json", res)
	t.Log(res)
}
