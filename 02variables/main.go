package main

import (
	"fmt"
)

func main() {
	var username string = "Garot"
	fmt.Println(username)
	fmt.Printf("Variable %v is of type: %T \n", username, username)

	// var isLoggedIn bool = true <-- this can be streamlined as done below
	// which should default to 'false'
	var isLoggedIn bool
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable %v is of type: %T \n", isLoggedIn, isLoggedIn)
}
