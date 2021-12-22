package main

import (
	"fmt"
	"sort"
)

func main() {
	welcome := "Welcome to slices!"
	fmt.Println(welcome)

	// arrays are fixed size which has to be declared at initialization
	// var carList = [5]string{"Toyota","Honda","Subaru","Ford","Chevrolet"}
	// slices do not have a defined length at declaration...
	var carList = []string{"Toyota", "Honda", "Subaru", "Ford", "Chevrolet"}
	fmt.Printf("The type of 'carList' is %T, the value is %v\n", carList, carList)

	carList = append(carList, "Buick", "Jeep", "Tesla")
	fmt.Println("The new value of 'carList', post append, is:", carList)

	// slice some stuff up !
	carList = carList[:3]
	fmt.Println("Now 'carList' is:", carList)

	highScores := make([]int, 4)
	highScores[0] = 245
	highScores[1] = 345
	highScores[2] = 445
	highScores[3] = 545

	fmt.Println(highScores)

	highScores = append(highScores, 645, 745, 845)
	fmt.Println(highScores)

	sort.IntsAreSorted(highScores)
	fmt.Printf("%v, %T\n", highScores, highScores)

	var courses = []string{"golang", "nodejs", "javascript", "python", "java", "C++", "ruby"}
	fmt.Println(courses)

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	courses = courses[index:4]
	fmt.Println(courses)
	fmt.Println(len(courses) - index)

}
