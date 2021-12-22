package main

import (
	"fmt"
	"strings"
)

func main() {
	welcome := "Welcome to my structs!"
	fmt.Println(welcome)

	garotm := User{"garotm", "garotm@go.dev", true, 34}
	fmt.Println(garotm)
	fmt.Println(garotm.Email)
	fmt.Printf("The details of the struct are: %+v\n", garotm)
	// use it in a more meaningful sentence
	fmt.Printf("My name is %v and %v's email is %v and he is %v years old.", strings.Title(garotm.Name), garotm.Name, garotm.Email, garotm.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
