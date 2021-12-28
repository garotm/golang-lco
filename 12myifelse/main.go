package main

import "fmt"

func main() {
	welcome := "Welcome to my if else instruction."
	fmt.Println(welcome)

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

}
