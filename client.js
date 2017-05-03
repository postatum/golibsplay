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

var lib = ffi.Library("./mylib.so", {
  PrintGoStr: ["longlong", [GoString]],
  PrintInt: ["longlong", ["longlong"]],
  PrintCStr: ["longlong", ["string"]],
  // PrintMap: ["longlong", ["void"]],
  PrintStruct: ["longlong", [Coords]],
});

var str = new GoString();
str["p"] = "Hello Node!";
str["n"] = 11;

lib.PrintGoStr(str);
lib.PrintInt(123);
lib.PrintCStr("Hello C!");
// lib.PrintMap({1: 2, 3: 4})
var coords = new Coords();
coords["x"] = 1
coords["y"] = 2
lib.PrintStruct(coords);
