package main

import(
	"fmt"
	"strings"
)

func main() {

	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking app.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for {
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

		// validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isVaildTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isVaildTicketNumber {

			// book tickets in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName, " "+ lastName)
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
			// print only first names
			firstNames := []string{}
			for _, booking := range bookings {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])		
			}
			fmt.Printf("The first names %v\n, firstNames")

			// exit app if no tickets are left
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isVaildTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}

	// switch statement example
	city := "London"

	switch city {
		//booking for New York conference
	case "New York":
		//booking for singapore and hong kong conference
	case "Singapore", "Hong Kong":
		//booking for london and berlin conference
	case "London", "Berlin":
		//booking for mexico conference
	case "Mexico city":
	default:
		fmt.Print("No vaild city selected")

	}

}

	    