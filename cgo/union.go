package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>

 union Hoge
 {
  struct Foo{
    int ifoo;
    double dfoo;
  } foo;
  struct Bar{
    int ibar;
    void *pbar;
  } bar;
 };
 union Hoge hoge;
 struct Foo f = { 3, 0.14 };

 union Hoge func_to_export() {
   hoge.foo = f;
 	return hoge;
 }
*/
import "C"
import "fmt"

func main() {
	n, _ := C.func_to_export()
	fmt.Println(n)
}
