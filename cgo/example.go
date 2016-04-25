package main

/*
#include <stdlib.h>
extern void func_to_export(const char*);
*/
import "C"
import "unsafe"

func main() {
	str := C.CString("イェーイ")
	defer C.free(unsafe.Pointer(str))
	C.func_to_export(str)
}
