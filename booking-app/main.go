package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	welcome := fmt.Sprintf("Welcome to the %v booking app!", conferenceName)
	const conferenceTickets = 50 // we don't want this to change or be negative
	var remainingTickets uint = 50
	var merchCount uint = 10
	var numberOfTickets uint

	fmt.Printf("conferenceTickets is of type %T\nremainingTickets is of type %T\nconferenceName is of type %T\n\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Println(welcome)
	fmt.Printf("\nWe have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("We also have %v merch items left\n\n", merchCount)

	getInput(numberOfTickets, conferenceTickets, remainingTickets)

}

func getInput(numberOfTickets, conferenceTickets, remainingTickets uint) {

	var firstName string
	var lastName string
	var email string
	prefix := "Please enter"

	fmt.Println(prefix, "your first name: ")
	fmt.Scan(&firstName) // the '&' is a pointer

	fmt.Println(prefix, "your last name: ")
	fmt.Scan(&lastName)

	fmt.Println(prefix, "your email address: ")
	fmt.Scan(&email)

	fmt.Println(prefix, "the number of tickets you'd like to purchase today: ")
	fmt.Scan(&numberOfTickets)

	fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, numberOfTickets, email)

	// update the number of tickets remaining
	remainingTickets = conferenceTickets - numberOfTickets
	fmt.Printf("There are now %v tickets remaining from the original %v we started with\n\n", remainingTickets, conferenceTickets)
}
