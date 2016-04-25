package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>

typedef int (*intFunc) ();
int bridge_int_func(intFunc f) {
  return f();
}

int fortytwo() {
  return 42;
}
*/
import "C"
import "fmt"

func main() {
	// 関数ポインタを保持
	f := C.intFunc(C.fortytwo)
	// 関数ポインタを渡す
	fmt.Println(int(C.bridge_int_func(f)))
}
