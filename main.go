package main

import "fmt"

func main() {
	var agency string = "LASRMA"
	var appName string = "Ticket Booking System"
	const seatsPerTrip int = 200
	var remTickets int = 200

	fmt.Printf("Welcome to %v %v - Purchase and book your train tickets here.\n", agency, appName)
	fmt.Printf("There are %v seats available and %v tickets left for sale. Book yours now!\n", seatsPerTrip, remTickets)

	// Get user info
	var userName string
	var ticketsBought int

	fmt.Printf("Please enter your username: ")
	fmt.Scanf("%v", &userName)
	fmt.Printf("Welcome %v, how many tickets do you want to purchase? ", userName)
	fmt.Scanf("%d", &ticketsBought)
	fmt.Printf("Hi %v, you have bought %d ticket(s).\n", userName, ticketsBought)
}