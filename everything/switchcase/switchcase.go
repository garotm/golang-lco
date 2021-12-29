package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	welcome_sw := "Welcome to switch and case in GO"
	fmt.Println(welcome_sw)

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		diceNumber := rand.Intn(45) + 1
		fmt.Println("Value of diceNumber is", diceNumber)
	}

	DiceNumber := rand.Intn(6) + 1

	switch DiceNumber {
	case 1:
		fmt.Println("Dice value is", DiceNumber, "you can move", DiceNumber, "spots.")
	case 2:
		fmt.Println("Dice value is", DiceNumber, "you can move", DiceNumber, "spots.")
	case 3:
		fmt.Println("Dice value is", DiceNumber, "you can move", DiceNumber, "spots.")
		fallthrough
	case 4:
		fmt.Println("Dice value is", DiceNumber, "you can move", DiceNumber, "spots.")
	case 5:
		fmt.Println("Dice value is", DiceNumber, "you can move", DiceNumber, "spots.")
		fallthrough
	case 6:
		fmt.Println("Dice value is", DiceNumber, "you can move", DiceNumber, "spots and roll again !!")
	}
}
