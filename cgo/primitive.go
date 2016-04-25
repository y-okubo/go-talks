package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
*/
import "C"
import "fmt"

func main() {
	n, err := C.sqrt(2) // err には errno が入るそう
	fmt.Println(n, err)
}
