package hpi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ExitOnError(err error, message string) {
	fmt.Println(message)
	log.Fatalln(err)
	os.Exit(1)
}

func WriteFile(filePath string, data ...string) error {

	f, err := os.Create(filePath)
	check(err)
	defer f.Close()

	allData := ""

	for _, d := range data {
		allData = allData + d + "\n"
	}
	_, err = f.WriteString(allData)
	check(err)
	//fmt.Printf("wrote %d bytes\n", n)

	f.Sync()

	return err

}

func CopyFile(source string, target string) error {
	b, err := os.ReadFile(source)
	if err != nil && err != io.EOF {
		return err
	}
	err = os.WriteFile(target, b, os.ModePerm|0755)
	return err
}

func ReadFile(source string) ([]byte, error) {
	b, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func IsFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ListFiles(path string) []string {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		ExitOnError(err, "Cannot list the files in directory : "+path)
	}
	return files
}

func MakeDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func ClearOutputDirectory(dir string) error {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		d, err := os.Open(dir)
		if err != nil {
			return err
		}
		defer d.Close()
		names, err := d.Readdirnames(-1)
		if err != nil {
			return err
		}
		for _, name := range names {
			err = os.RemoveAll(filepath.Join(dir, name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func IsFileExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveFile(file string) error {
	err := os.Remove(file)
	return err
}

func RemoveFolder(file string) error {
	err := os.RemoveAll(file)
	return err
}

func JsonFromFile(file_path string) (interface{}, error) {
	b, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}

	// Convert to map
	mapObj := map[string]interface{}{}
	err = json.Unmarshal(b, &mapObj)
	if err == nil {
		return mapObj, nil
	}

	// Convert to array
	arrObj := []interface{}{}
	err = json.Unmarshal(b, &arrObj)
	if err == nil {
		return arrObj, nil
	}

	return nil, err
}

func JsonToMap(b []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func CastToMap(in interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	return JsonToMap(b)
}

func JsonMapFromFile(file_path string) (map[string]interface{}, error) {
	obj, err := JsonFromFile(file_path)
	if err != nil {
		return nil, err
	}

	mapObj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, errors.New("data is not map")
	}
	return mapObj, nil
}

func JsonArrayFromFile(file_path string) ([]interface{}, error) {
	obj, err := JsonFromFile(file_path)
	if err != nil {
		return nil, err
	}

	arrObj, ok := obj.([]interface{})
	if !ok {
		return nil, errors.New("data is not array")
	}
	return arrObj, nil
}

// WriteJsonFile is function to write json data to file
func WriteJsonFile(path string, data interface{}) (err error) {

	bData, err := json.Marshal(data)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(path, bData, 0644)

	return

}

func Interface2Str(o interface{}) string {
	switch v := o.(type) {
	case map[string]interface{}:
		return json2Str(v)
	case []interface{}:
		return jsonArray2Str(v)
	default:
		return o.(string)
	}
}

func jsonArray2Str(v []interface{}) string {
	size := len(v)
	var jsonStrArray []string
	for i, s := range v {
		b, _ := json.Marshal(s)
		jsonStr := string(b)
		if i < size-1 {
			jsonStr = jsonStr + ","
		}
		jsonStrArray = append(jsonStrArray, jsonStr)
	}
	str := fmt.Sprintf("%v", jsonStrArray)
	return str
}

func json2Str(v map[string]interface{}) string {
	str, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(str)
}
