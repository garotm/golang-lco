package main

import "fmt"

func main() {
	var welcome string = "Welcome to the Go Conference booking app!"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	var merchCount uint = 10

	fmt.Printf("conferenceTickets is of type %T\nremainingTickets is of type %T\nconferenceName is of type %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println(welcome)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("We also have %v merch items left\n", merchCount)

	userName := "Garot"
	userTickets := 2
	fmt.Printf("User %v booked %v ticket.\n", userName, userTickets)
}
