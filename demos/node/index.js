var ref = require("ref");
var ffi = require("ffi");
var Struct = require("ref-struct")

var GoString = Struct({
  p: "string",
  n: "longlong"
});

var Coords = Struct({
  x: "longlong",
  y: "longlong"
});

var lib = ffi.Library("../../mylib.so", {
  PrintGoStr: ["longlong", [GoString]],
  PrintInt: ["longlong", ["longlong"]],
  PrintCStr: ["longlong", ["string"]],
  PrintStruct: ["longlong", [Coords]],
  PrintStructMeth: ["longlong", [Coords]],
  // PrintMap: ["longlong", ["void"]],
  GetFirstJSONElement: ["longlong", ["string"]],
});

var str = new GoString();
str["p"] = "Hello Node!";
str["n"] = 11;

lib.PrintGoStr(str);
lib.PrintInt(123);
lib.PrintCStr("Hello C!");
var coords = new Coords();
coords["x"] = 1
coords["y"] = 2
lib.PrintStruct(coords);
lib.PrintStructMeth(coords);
// lib.PrintMap({1: 2, 3: 4})
lib.GetFirstJSONElement("../testdata.json")