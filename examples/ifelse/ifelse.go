package main

import (
	"fmt"
)

func main() {
	welcome_ifelse := "Welcome to my if else instruction."
	fmt.Println(welcome_ifelse)

	//var result string

	count := [3]int{4, 10, 45}

	for i, r := range count {
		if r < 10 {
			r := "Regular user"
			fmt.Println(count[i], r)
		} else if r > 10 {
			r := "Irregular user"
			fmt.Println(count[i], r)
		} else {
			r := "Exaclty 10"
			fmt.Println(count[i], r)
		}
	}

	a := 9
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}

	if num := 3; num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "is greater than 10")
	}

}
