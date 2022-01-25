package main

import (
	"fmt"
	"time"
)

func main() {
	welcome_t := "Welcome to time!"
	fmt.Println(welcome_t)

	presentTime := time.Now()
	fmt.Println(presentTime)
	// This formatting is NOT arbitraty, this is what must be used... ???
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	if presentTime.Format("15:04:05") >= "17:00:00" {
		fmt.Println("It's 5 'oclock somewhere !!")
	} else {
		fmt.Println("Get back to work !!")
	}
}
