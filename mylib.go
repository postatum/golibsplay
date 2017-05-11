package main

/*
typedef struct { long long x; long long y; } Coords;
*/
import "C"
import (
	"fmt"
	"github.com/postatum/golibsplay/core"
	"path/filepath"
)

//export PrintGoStr
func PrintGoStr(s string) {
	core.PrintStr(s)
}

//export PrintInt
func PrintInt(s int) {
	fmt.Println(s)
}

//export PrintCStr
func PrintCStr(s *C.char) {
	fmt.Println(C.GoString(s))
}

//export PrintStruct
func PrintStruct(c C.Coords) {
	fmt.Println(c)
}

//export PrintStructMeth
func (c C.Coords) PrintStructMeth() {
	fmt.Println(c)
}

//export GetFirstJSONElement
func GetFirstJSONElement(p *C.char) *C.char {
	path, err := filepath.Abs(C.GoString(p))
	if err != nil {
		return C.CString(string(core.DumpError(err)))
	}
	return C.CString(core.GetFirstElement(path))
}

func main() {}

// v-- Not used by C yet --v

//export PrintMap
func PrintMap(m map[int]int) {
	fmt.Println(m)
}
