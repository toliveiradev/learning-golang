package main

import (
	"strings"
)

func validateUserInput(firstName string, lastName string, email string, userTicketsCount uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicketsCount > 0 && userTicketsCount < remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
