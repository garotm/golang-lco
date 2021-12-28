package main

import "fmt"

func main() {
	welcome_maps := "Welcome to maps"
	fmt.Println(welcome_maps)

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["NJ"] = "Nodejs"
	languages["PY"] = "Python"

	fmt.Println("List of all languages", languages)
	fmt.Println("'js' is short for", languages["JS"])

	for key, values := range languages {
		fmt.Printf("for key %v, value is %v\n", key, values)
	}
}
