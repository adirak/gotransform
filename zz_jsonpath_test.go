package main

import (
	"testing"

	"github.com/adirak/gotransform/hpi"
)

func TestJsonPathRead(t *testing.T) {

	obj, err := hpi.JsonFromFile("./conf/random_int.json")
	if err != nil {
		t.Error(err)
	}

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		t.Error("data is not map")
	}

	jp, err := hpi.NewJsonPathWithRoot(&mapObj)
	if err != nil {
		t.Error(err)
	}

	val := jp.Value("transform.process[0].value")

	err = jp.Delete("transform.process[0]")
	if err != nil {
		t.Error(err)
	}

	val = jp.Value("transform.process[0].value")

	t.Log(val)
}

func TestJsonPathWrite(t *testing.T) {

	mapObj := map[string]interface{}{}
	jp, err := hpi.NewJsonPathWithRoot(&mapObj)
	if err != nil {
		t.Error(err)
	}

	jp.Set("name", "Adirak")
	jp.Set("surname", "Kaewmahing")
	jp.Set("a.b.arr[3]", "3")
	jp.Set("a.b.arr[2]", "2")
	jp.Set("a.b.arr[1]", "1")
	jp.Set("a.b.arr[0]", "0")

	jp.Set("x.y.z[1].name", "supote")
	jp.Set("x.y.z[1].surname", "sirimaha")

	jp.Set("x.y.z[0]", map[string]interface{}{})
	jp.Set("x.y.z[0].name", "aaa")
	jp.Set("x.y.z[0].surname", "bbb")

	jp.Delete("x.y.z[0].surname")

	hpi.WriteJsonFile("./conf/jsonMap.json", jp.ToMap())

}
