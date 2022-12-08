package helper

import "strings"

func ValidateUsersInput(firstName string, lastName string, email string, userTickets uint, remTickets uint) (bool, bool, bool) {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validTickets := userTickets > 0 && userTickets <= remTickets
	return validName, validEmail, validTickets
}