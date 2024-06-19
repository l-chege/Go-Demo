package main

import "fmt"

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking app.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask for user input
	fmt.Println("Enter your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your Last Name:")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Enter number of Tickets:")
	fmt.Scanln(&userTickets)

	// book tickets in system
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName, " "+ lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)

}