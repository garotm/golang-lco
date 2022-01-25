package main

import (
	"fmt"
	"strings"
)

func main() {
	welcomeMethods := "Welcome to 'methods' in GOlang!"
	fmt.Println(welcomeMethods)

	// pulled in the the structs lesson
	garotm := User{"Garot M Conklin Sr.", "garotm", "garotm@go.dev", true, 34, 74}
	fmt.Println(garotm)
	fmt.Println(garotm.Email)
	fmt.Printf("The details of the struct are: %+v\n", garotm)
	// use it in a more meaningful sentence
	fmt.Printf("My name is %v and %v's email is %v and he is %v years old.\n", strings.Title(garotm.Name), garotm.Name, garotm.Email, garotm.Age)

	// call some methods
	garotm.GetStatus()
	garotm.newMail()
	garotm.yourName()
}

type User struct {
	Name       string
	SystemName string
	Email      string
	Status     bool
	Age        int
	oneAge     int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) newMail() {
	u.Email = "test@go.dev"
	fmt.Println("The email of this user is", u.Email)
}

func (u User) yourName() {
	u.Name = "Garot Michael Conklin Sr."
	u.oneAge = 49
	fmt.Println("Hello, my name is", u.Name)
	fmt.Println("My age is", u.oneAge)
	fmt.Println("And my system name is", u.SystemName)
	u.SystemName = "something else" // change the system name
	fmt.Println("And now my system name is", u.SystemName)
}
