package main

import (
	"fmt"
)

func main() {
	welcome := "Welcome to arrays"
	fmt.Println(welcome)

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "tomato"
	fruitList[3] = "Peach"
	// notice in the output that since we only defined 3 strings, there is a 'weird' space wehre 2 would be
	fmt.Println("The fruitList contains:", fruitList)
	fmt.Println("The length of the fruitList is:", len(fruitList))

	// declare it all in one line
	var vegList = [4]string{"potato", "corn", "lettuce", "beans"}
	fmt.Printf("The vegList is %v with a length of %v", vegList, len(vegList))

}
