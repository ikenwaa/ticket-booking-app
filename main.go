package main

import (
	"fmt"
	"strings"
)

var appName = "Atlantic City Train Ticket Application"
const seatsPerTrip uint = 200
var remTickets uint = 200
var bookings = []string{}

func main() {
	greetUser()

	for {
		// User input plus validation
		firstName, lastName, email, userTickets := getUserInput()
		validName, validEmail, validTickets := validateUsersInput(firstName, lastName, email, userTickets, remTickets)

		if validName && validEmail && validTickets {
			bookTickets(remTickets, userTickets, bookings, firstName, lastName, email)
			// Print the first names of users that paid for tickets
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first name of customers that have paid for tickets include: %v\n", firstNames)

			if remTickets == 0 {
				fmt.Println("The tickets have been sold out. Kindly wait for the next round.")
				break
			}
		} else {
			if !validName {
				fmt.Println("Your first name or second name is too short. Please try again.")
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

func greetUser() {
	fmt.Printf("Welcome to %v.\n", appName)
	fmt.Printf("There are %v tickets per trip and %v tickets are left.\n", seatsPerTrip, remTickets)
	fmt.Println("Hurry and purchase your ticket.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUsersInput(firstName string, lastName string, email string, userTickets uint, remTickets uint) (bool, bool, bool) {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validTickets := userTickets > 0 && userTickets <= remTickets
	return validName, validEmail, validTickets
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(remTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string) {
	remTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Hi %v, you have bought %d ticket(s). You will receive a confirmation email at %v.\n", firstName, userTickets, email)
	fmt.Printf("Number of train tickets left is: %v\n", remTickets)
}