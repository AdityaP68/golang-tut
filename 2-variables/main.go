package main

import "fmt"

// capital letter meas its a global var same as public keyword
const LoginToken string = "lmao"

func main() {
	var username string = "aditya"
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	//default values
	var temp int
	fmt.Println(temp)
	fmt.Printf("variable is of type: %T \n", temp)

	//implicit type

	var url = "www.google.com"
	fmt.Println(url)

	// no var style can ony used inside a method
	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)

}
