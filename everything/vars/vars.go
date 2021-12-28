package main

import (
	"fmt"
)

const LoginToken string = "askdjfhkuyer/fhaerij" // Public/Global variable

func main() {
	fmt.Println("Welcome to vars.")
	vars()
}

func vars() {
	var username string = "Garot"
	fmt.Println(username)
	fmt.Printf("Variable %v is of type: %T \n", username, username)

	// var isLoggedIn bool = true <-- this can be streamlined as done below
	// which should default to 'false'
	var isLoggedIn bool
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable %v is of type: %T \n", isLoggedIn, isLoggedIn)

	// var smallVal uint8 = 255
	var smallVal byte = 254
	fmt.Println(smallVal)
	fmt.Printf("Variable %v is of type: %T \n", smallVal, smallVal)

	var smallFloat float32 = 254.567238389
	fmt.Println(smallFloat)
	fmt.Printf("Variable %v is of type: %T \n", smallFloat, smallFloat)

	// default value is always 0
	fmt.Printf("%v, %T\n", LoginToken, LoginToken)
	// print out each letter in LoginToken
	for i, letter := range LoginToken {
		fmt.Println(i, string(letter))
	}
}
