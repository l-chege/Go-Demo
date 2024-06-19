package main

import "fmt"

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"

	fmt.Printf("Welcome to %v booking app.\n
	            We have total of %v tickets and %v are still available.\n
	            Get your tickets here to attend", 
				conferenceName, conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets unit

	// ask for user input
	fmt.Println("Enter your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your Lat Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	// book ticket in the system
	remainingTickets = remainingTickets - userTicketsconferenceName

	fmt.Prrintf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", 
	firstName, lastName, userTickets, email)

}