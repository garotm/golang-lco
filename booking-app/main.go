package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	welcome := fmt.Sprintf("Welcome to the %v booking app!", conferenceName)
	const conferenceTickets = 50
	var remainingTickets uint
	var merchCount uint = 10

	fmt.Printf("conferenceTickets is of type %T\nremainingTickets is of type %T\nconferenceName is of type %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println(welcome)
	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("We also have %v merch items left\n\n", merchCount)

	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint
	prefix := "Please enter"

	// ask user for their first/last name, email addr and number of tickets they want
	fmt.Println(prefix, "your first name: ")
	fmt.Scan(&firstName) // the '&' is a pointer

	fmt.Println(prefix, "your last name: ")
	fmt.Scan(&lastName)

	fmt.Println(prefix, "your email address: ")
	fmt.Scan(&email)

	fmt.Println(prefix, "the number of tickets you'd like to purchase today: ")
	fmt.Scan(&numberOfTickets)

	// update the number of tickets remaining
	remainingTickets = conferenceTickets - numberOfTickets

	fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, numberOfTickets, email)
	fmt.Printf("There are now %v tickets remaining from the original %v we started with", remainingTickets, conferenceTickets)

}
