package main

// this one is a little confusing to me...

import "fmt"

func main() {
	welcomeDefer := "Welcome to defer-r!"
	fmt.Println(welcomeDefer)
	fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("three")
	myDefer()
}

func myDefer() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
