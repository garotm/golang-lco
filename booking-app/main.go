package main

import "fmt"

func main() {
	var welcome string = "Welcome to the Go Conference booking app!"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets int
	var merchCount uint = 10

	fmt.Printf("conferenceTickets is of type %T\nremainingTickets is of type %T\nconferenceName is of type %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println(welcome)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("We also have %v merch items left\n", merchCount)

	var firstName string
	var lastName string
	var emailAddr string
	var numOfTickets int
	prefix := "Please enter"

	// ask user for their first/last name, email addr and number of tickets they want
	fmt.Println(prefix, "your first name: ")
	fmt.Scan(&firstName) // the '&' is a pointer

	fmt.Println(prefix, "your last name: ")
	fmt.Scan(&lastName)

	fmt.Println(prefix, "your email address: ")
	fmt.Scan(&emailAddr)

	fmt.Println(prefix, "the number of tickets you'd like to purchase today: ")
	fmt.Scan(&numOfTickets)

	fmt.Printf("User %v %v with email address %v booked %v tickets.\n", firstName, lastName, emailAddr, numOfTickets)

	// update the number of tickets remaining
	remainingTickets = conferenceTickets - numOfTickets

	fmt.Printf("There are now %v tickets remaining from the original %v we started with", remainingTickets, conferenceTickets)

}
