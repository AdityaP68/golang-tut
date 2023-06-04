package main

import "fmt"

func main() {
	fmt.Println("arrays not used that much in golang")
	var fruits [5]string
	fruits[0] = "orange"
	fruits[1] = "apple"


	fmt.Println("the array has, ", fruits)
	fmt.Println(len(fruits))

	var veg = [4]string{"potatoes", "tomato"}
	fmt.Println(veg)

	
}
