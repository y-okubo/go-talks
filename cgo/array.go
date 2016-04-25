package main

/*
#include <stdio.h>
#include <stdlib.h>
// int a[] = {0, 1, 2, 3};
int a[] = {'A', 'B', 'C', 'D'};
int* func_to_export() {
	return a;
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	n, _ := C.func_to_export()
	// fmt.Println(n[0])

	var arr []C.int
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	slice.Cap = int(4)
	slice.Len = int(4)
	slice.Data = uintptr(unsafe.Pointer(n))
	fmt.Println(arr)
}
