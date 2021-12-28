package main

import (
	"fmt"
)

func main() {
	welcome := "Welcome to pointers!"
	fmt.Println(welcome)

	var ptr *int
	fmt.Printf("%v, %T\n", ptr, ptr)

	myNumber := 23

	// & is used for a declared vaiable pointer
	pointer := &myNumber

	fmt.Println("The value in memory of the pointer is", pointer)
	fmt.Println("The value of the pointer is", *pointer)
	fmt.Println("The value in memory of the pointer is", &pointer)

	fmt.Println("The declared value of myNumber is:", myNumber)
	*pointer = *pointer + 2
	fmt.Println("The new value of myNumber is:", myNumber)
}
