package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the pizza: ")

	// comma ok || err ok
	// if value is not being used the use underscore or unsed value 
	//error will be thrown
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the rating is %T, ", input)
}
