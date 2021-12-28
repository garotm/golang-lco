package main

import "fmt"

func main() {

	const f_name, l_name = "Garot", "Conklin"
	fmt.Println("Hello", f_name, l_name, "... Welcome to GO!")

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
