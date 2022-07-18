package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket int, remainingTicket int) (bool, bool, bool) {
	isValidateName := len(firstName) > 2 && len(lastName) > 2

	isValidateEmail := strings.Contains(email, "@")

	isValidateTicketNumber := userTicket > 0 && userTicket <= remainingTicket

	return isValidateName, isValidateEmail, isValidateTicketNumber
}
