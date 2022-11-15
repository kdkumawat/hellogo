package main

import "fmt"

const LoginToken string = "slkdfjslkdfjslkdf"

func main() {
	fmt.Println("Welcome to variables")
	var username string = "kuldeep"

	fmt.Println("Hello", username)
	fmt.Printf("Type of username %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn %T \n", isLoggedIn)

	var smallint uint8 = 255
	fmt.Println(smallint)
	fmt.Printf("Type of smallint %T \n", smallint)

	var smallFloat float32 = 255.123456789123
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallFloat %T \n", smallFloat)

	var float float64 = 255.123456789123456789
	fmt.Println(float)
	fmt.Printf("Type of float %T \n", float)

	var myVariable int
	fmt.Println(myVariable)

	var myStr string
	fmt.Println("Default value of string", myStr)

	// implicit
	var implicitVar = "I am string"
	fmt.Println(implicitVar)

	// no var style
	numberOfUsers := 30000.12
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Type of LoginToken %T \n", LoginToken)
}
