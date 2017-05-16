package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/jstoiko/golibsplay/core"
)

type Vertex struct {
	A, B int
}

func ReturnGoStr(s string) string {
	return s
}

func ReturnInt(s int) int {
	return s
}

func ReturnMap(m map[int]int) map[int]int {
	return m
}

func GetLen(s string) int {
	return len(s)
}

func PrintGoStr(s string) {
	core.PrintStr(s)
}

func GetFirstJSONElement(c string) string {
	content := []byte(c)
	el, err := core.GetFirstContentElement(content)
	if err != nil {
		return core.DumpError(err)
	}
	return el
}

func main() {
	// js.Module.Set("exports", map[string]interface{}{ // node
	js.Global.Set("mylib", map[string]interface{}{ // browser
		"ReturnGoStr":         ReturnGoStr,
		"ReturnInt":           ReturnInt,
		"ReturnMap":           ReturnMap,
		"GetLen":              GetLen,
		"PrintGoStr":          PrintGoStr,
		"GetFirstJSONElement": GetFirstJSONElement,
	})
}
