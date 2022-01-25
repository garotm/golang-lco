package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	welcome_m := "All about crypto!"
	fmt.Println(welcome_m)
	/*
		myNumberOne := 4
		myNumberTwo := 4.56

		fmt.Println("The sum is: ", float64(myNumberOne)+myNumberTwo)
	*/
	/* random number(s) from math/rand
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5) + 1)
	*/
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(45))
	fmt.Println(myRandomNum)
}
