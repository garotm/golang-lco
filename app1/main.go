package main

import (
	"fmt"
	"strings"
)

func main() {

	const f_name, l_name = "Garot", "Conklin"
	fmt.Println("Hello", f_name, l_name, "... Welcome to GO!")

	str := "Garot Michael Conklin"
	substr := "a"
	occurance := strings.Count(str, substr)
	length := len(str)

	fmt.Println("The number of occurrences of the letter:", substr, ", in the string", str, "is: ", occurance)
	fmt.Println("The length of the string is", length, "characters long.")
	fmt.Printf("%v, %T", str, str)

	fizz_buzz()
}

func fizz_buzz() {

	count := 100
	b := "buzz"
	f := "fizz"
	fb := "fizzbuzz"

	for i := 0; i < count; i++ {
		if i%15 == 0 {
			fmt.Println(fb)
		} else if i%3 == 0 {
			fmt.Println(f)
		} else if i%2 == 0 {
			fmt.Println(b)
		} else {
			fmt.Println(i)
		}
	}
}
