package main

import (
	"fmt"
	"strings"
)

func main() {
	welcomeMoreMethods := "Welcome to more methods!"
	fmt.Println(welcomeMoreMethods)

	garotm := User{"garotm", "garotm@go.dev", true, 34}
	fmt.Println(garotm)
	fmt.Println(garotm.Email)
	fmt.Printf("The details of the struct are: %+v\n", garotm)
	// use it in a more meaningful sentence
	fmt.Printf("My name is %v and %v's email is %v and he is %v years old.\n", strings.Title(garotm.Name), garotm.Name, garotm.Email, garotm.Age)
	// let me knwow where we are here
	fmt.Println("this is the 'GetStatus' method printing:")
	garotm.GetStatus()
	garotm.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "methods@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
