package main

import "fmt"

func main() {
	welcome_f := "Welcome to functions in golang !"
	fmt.Println(welcome_f)

	// call the function
	greeter() // where you place this matters in the order of execution
	greeterTwo()
	func() {
		fmt.Println("This is an anonymous function")
	}() // this function goes inside main(), others DO NOT!

	result := adder(3, 5)

	fmt.Println("result is :", result)

	proRes, myMessage := proAdder(2, 4, 5, 6, 8, 9)
	fmt.Printf("The sum of the numbers is %v and the type is %T\n", proRes, proRes)
	fmt.Println("This is my message to you -->", myMessage)
}

func greeter() {
	fmt.Println("Namaste from GOlang!")
}

func greeterTwo() {
	fmt.Println("This is the 2nd greeter function.")
}

func adder(valOne int, valTwo int) int { // don't forget the function sigature @!
	return valOne + valTwo
}

func proAdder(numbers ...int) (int, string) { // you can call multiple things here
	total := 0

	for _, val := range numbers {
		total += val
	}
	return total, "This is my message"
}
