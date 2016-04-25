package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
struct stat* func_to_export() {
	struct stat* pbuf = malloc(sizeof(struct stat));
  stat("test.txt", pbuf);
	return pbuf;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	n, _ := C.func_to_export()
	fmt.Println(n.st_ino)
	fmt.Println(n.st_uid)
	fmt.Println(n.st_gid)
	fmt.Println(n.st_size)
	C.free(unsafe.Pointer(n))
}
