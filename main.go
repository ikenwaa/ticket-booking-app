package main

import "fmt"

func main() {
	agency := "LASRMA"
	appName := "Ticket Booking Application"
	const seatsPerTrip int = 200
	var remTickets uint = 200
	var bookings []string

	fmt.Printf("Welcome to %v %v - Purchase and book your train tickets here.\n", agency, appName)
	fmt.Printf("There are %v seats available and %v tickets left for sale. Book yours now!\n", seatsPerTrip, remTickets)

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
	fmt.Printf("Hi %v, you have bought %d ticket(s). You will receive a confirmation email at %v.\n", firstName, trainTickets, email)

	remTickets -= trainTickets
	bookings = append(bookings, firstName + " " + lastName)
	
	fmt.Printf("Number of train tickets left is: %v\n", remTickets)
	fmt.Printf("Type of bookings: %T\n", bookings)
	fmt.Printf("First element: %v\n", bookings[0])
	fmt.Println(bookings)
	fmt.Printf("Length: %v\n", len(bookings))
}