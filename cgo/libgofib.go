package main

import (
	"C"
	"log"
)

//export fib
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func init() {
	log.Println("Loaded!!")
}

func main() {
}
