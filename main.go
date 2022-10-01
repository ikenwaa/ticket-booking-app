package main

import "fmt"

func main() {
	var agency = "LASRMA"
	var appName = "Ticket Booking System"
	const seatsPerTrip = 200
	var remTickets = 200

	fmt.Printf("Welcome to %s %s - Purchase and book your train tickets here.\n", agency, appName)
	fmt.Printf("There are %d seats available and %d tickets left for sale. Book yours now!\n", seatsPerTrip, remTickets)
}