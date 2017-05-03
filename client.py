from ctypes import *

lib = cdll.LoadLibrary("./mylib.so")
# lib = ctypes.CDLL("mylib.so")


class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]


class Coords(Structure):
    _fields_ = [("x", c_longlong), ("y", c_longlong)]

msg = GoString("Hello Python!", 13)

lib.PrintGoStr(msg)
lib.PrintInt(123)
lib.PrintCStr("Hello C!")
lib.PrintStruct(Coords(1, 2))
lib.PrintStructFancy(Coords(3, 4))
# lib.PrintMap({1: 2, 3: 4})
