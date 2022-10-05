package main

import (
	"fmt"
	"strings"
)

func main() {
	agency := "LASRMA"
	appName := "Ticket Booking Application"
	const seatsPerTrip int = 200
	var remTickets uint = 200
	bookings := []string{}

	fmt.Printf("Welcome to %v %v - Purchase and book your train tickets here.\n", agency, appName)
	fmt.Printf("There are %v seats available and %v tickets left for sale. Book yours now!\n", seatsPerTrip, remTickets)

	for {
		// Get user info
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Enter your first name: \n")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: \n")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: \n")
		fmt.Scan(&email)
		fmt.Printf("Welcome %v, how many tickets do you want to buy? \n", firstName)
		fmt.Scan(&userTickets)

		// User input validation
		validName := len(firstName) >= 2 && len(lastName) >= 2
		validEmail := strings.Contains(email, "@")
		validTickets := userTickets > 0 && userTickets <= remTickets

		if validName && validEmail && validTickets {
			remTickets -= userTickets
			bookings = append(bookings, firstName + " " + lastName)
			
			fmt.Printf("Hi %v, you have bought %d ticket(s). You will receive a confirmation email at %v.\n", firstName, userTickets, email)
			fmt.Printf("Number of train tickets left is: %v\n", remTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("This is the first name of our bookings are: %v\n", firstNames)

			if remTickets == 0 {
				fmt.Println("The tickets have been sold out. Kindly wait for the next round.")
				break
			} 
		} else {
			if !validName {
				fmt.Println("You entered an invalid first name or second name. Please try again.")
			} 
			if !validEmail {
				fmt.Println("'@' symbol is missing in email. Please try again.")
			} 
			if !validTickets {
				fmt.Println("You entered an invalid number of tickets. Please try again.")
			}
		}
	}
}