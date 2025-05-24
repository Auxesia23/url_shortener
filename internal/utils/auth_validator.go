package utils

import (
	"net/mail"
	"regexp"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidatePassword(password string) bool {
	// regex for atleast one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
	
	// regex for atleast one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString
	
	//return true if password is valid
	return len(password) >= 8 && hasUpper(password) && hasDigit(password)
}