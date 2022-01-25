package main

import "fmt"

func main() {
	welcome_l := "Welcome to loops in GOlang!"
	fmt.Println(welcome_l)

	daysOfTheWeek := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(daysOfTheWeek)

	for i := 0; i < len(daysOfTheWeek); i++ {
		fmt.Println(daysOfTheWeek[i])
	}

	for i := range daysOfTheWeek {
		fmt.Println(daysOfTheWeek[i])
	}

	for index, daysOfTheWeek := range daysOfTheWeek {
		fmt.Printf("The index is %v and the day is %v\n", index, daysOfTheWeek)
	}

	for i, daysOfTheWeek := range daysOfTheWeek {
		fmt.Printf("The day of the week is %v\n", daysOfTheWeek)
		if daysOfTheWeek == "Monday" {
			fmt.Println("Someone has a case of the", daysOfTheWeek, "'s...")
		} else if daysOfTheWeek == "Tuesday" || daysOfTheWeek == "Wednesday" {
			fmt.Println("Just getting started here...")
		} else if daysOfTheWeek == "Friday" {
			fmt.Println("TGIF !!")
			dayOfTheWeek := i
			if dayOfTheWeek == 5 {
				fmt.Println("Man, already", dayOfTheWeek, "days gone...")
			}
		} else {
			fmt.Println("Just another day in the life...")
		}
	}

	someValue := 1
	for someValue <= 10 {
		fmt.Println("count is:", someValue)
		someValue++
	}

lco:
	fmt.Println("There used to be a 'other' right here")

	myArray := []string{"this", "that", "and", "the", "other", "thing"}
	for index := range myArray {
		fmt.Println(myArray[index])
		if myArray[index] == "the" {
			break
		} else if myArray[index] == "other" {
			goto lco
		} else {
			continue
		}

	}

}
