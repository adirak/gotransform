package main

import (
	"fmt"

	"github.com/adirak/gotransform/hpi"
)

func main() {

	fmt.Println("Hello, world.")

	obj, err := hpi.JsonFromFile("./conf/basic/random_alphabet.json")
	if err != nil {
		panic(err)
	}

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		panic("data is not map")
	}

	dataBus, _ := mapObj["input"].(map[string]interface{})
	transform, _ := mapObj["transform"].(map[string]interface{})

	res, err := hpi.ProcessTransform(dataBus, transform)
	if err != nil {
		panic(err)
	}

	hpi.WriteJsonFile("./conf_result/random_alphabet.json", res)
	fmt.Println(res)

}
