package main

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		/*
			var bookings = []string{}
			bookings := []string{}
		*/

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName+" "+email+" "+strconv.FormatUint(uint64(userTickets), 10))

		/*
			fmt.Printf("The whole slice is: %v\n", bookings)
			fmt.Printf("The first value in the slice is: %v\n", bookings[0])
			fmt.Printf("The slice type is: %T\n", bookings)
			fmt.Printf("The slice length is: %v\n", len(bookings))
		*/

		fmt.Printf("\nWe have bookings for %v and %v tickets are booked.\n", strings.Join(bookings, ""), userTickets)
		fmt.Printf("There are only %v tickets remaining for the %v.\n\n", remainingTickets, conferenceName)

		if remainingTickets == 0 {
			fmt.Printf("We have no more tickets left :(\n\n")
			break
		}
	}

	fmt.Printf("This is what we have booked:\n")
	for index, booked := range bookings {
		fmt.Printf("\t%v: %v\n", index, booked)
	}
}

func greetUsers() {
	fmt.Printf("\n\tWelcome to the %v booking application\n", conferenceName)
	if remainingTickets == 50 {
		fmt.Printf("\tWe have a total of %v tickets and all %v tickets are still available.\n", conferenceTickets, remainingTickets)
	} else {
		fmt.Printf("\tWe have a total of %v ticket and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	}
	fmt.Printf("\tGet your tickets here to attend\n\n")
}
