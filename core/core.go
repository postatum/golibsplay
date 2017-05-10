package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func PrintStr(s string) {
	fmt.Println(s)
}

func ParseJSONFile(path string) (interface{}, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return WrapError(err), err
	}
	return ParseJSON(dat)
}

func ParseJSON(dat []byte) (interface{}, error) {
	parsed := []interface{}{}
	if err := json.Unmarshal(dat, &parsed); err != nil {
		return WrapError(err), err
	}
	return parsed, nil
}

func WrapError(e error) map[string]string {
	return map[string]string{"error": e.Error()}
}

func DumpError(e error) []byte {
	res, _ := json.Marshal(WrapError(e))
	return res
}

func GetFirstElement(path string) string {
	dat, err := ParseJSONFile(path)
	if err != nil {
		return dat.(string)
	}
	arrDat := dat.([]interface{})
	res, err := json.Marshal(arrDat[0])
	if err != nil {
		return string(DumpError(err))
	}
	return string(res)
}
