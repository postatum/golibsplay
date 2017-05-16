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
		return nil, err
	}
	return ParseJSON(dat)
}

func ParseJSON(dat []byte) (interface{}, error) {
	parsed := []interface{}{}
	if err := json.Unmarshal(dat, &parsed); err != nil {
		return nil, err
	}
	return parsed, nil
}

func WrapError(e error) map[string]string {
	return map[string]string{"error": e.Error()}
}

func DumpError(e error) string {
	res, _ := json.Marshal(WrapError(e))
	return string(res)
}

func GetFirstFileElement(path string) (string, error) {
	parsed, err := ParseJSONFile(path)
	if err != nil {
		return "", err
	}
	return getFirstCommon(parsed)
}

func GetFirstContentElement(content []byte) (string, error) {
	parsed, err := ParseJSON(content)
	if err != nil {
		return "", err
	}
	return getFirstCommon(parsed)
}

func getFirstCommon(data interface{}) (string, error) {
	arrDat := data.([]interface{})
	res, err := json.Marshal(arrDat[0])
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func GetLexicalInfo(path string) (string, error) {
	jsonb, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	jsons := string(jsonb)
    lf := rune(0x0A)
    line := 1
    var character int

    for _, b := range jsons {
        if b == lf {
            line++
            character = 0
        }
        character++
        fmt.Printf("%d - %d - %s\n", line, character, string(b))
    }
    return "", nil
}
