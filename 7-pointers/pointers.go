package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointers")
	// var ptr *int
	// fmt.Println("the value of ptr is, ", ptr) // nil

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual ptr is, ", ptr)
	fmt.Println("Value of actual ptr is, ", *ptr)

}
