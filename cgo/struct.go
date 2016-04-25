package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
struct stat buf;
struct stat func_to_export() {
  stat("test.txt", &buf);
	return buf;
}
*/
import "C"
import "fmt"

func main() {
	n, _ := C.func_to_export()
	fmt.Println(n.st_ino)
	fmt.Println(n.st_uid)
	fmt.Println(n.st_gid)
	fmt.Println(n.st_size)
}
