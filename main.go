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
		var trainTickets uint

		fmt.Printf("Please enter your first name: \n")
		fmt.Scan(&firstName)
		fmt.Printf("Please enter your last name: \n")
		fmt.Scan(&lastName)
		fmt.Printf("Please enter your email address: \n")
		fmt.Scan(&email)
		fmt.Printf("Welcome %v, how many tickets do you want to buy? \n", firstName)
		fmt.Scan(&trainTickets)

		if trainTickets <= remTickets {
			remTickets -= trainTickets
			bookings = append(bookings, firstName + " " + lastName)
			
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
			if remTickets == 1 {
				fmt.Printf("Sorry, you cannot buy %v tickets as there is just %v ticket left!\n", trainTickets, remTickets)
				continue
			} else {
				fmt.Printf("Sorry, you cannot buy %v tickets as there are just %v tickets left!\n", trainTickets, remTickets)
				continue
			}
		}
		
		fmt.Printf("Hi %v, you have bought %d ticket(s). You will receive a confirmation email at %v.\n", firstName, trainTickets, email)
	}
}