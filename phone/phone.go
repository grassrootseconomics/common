package phone

import (
	"errors"
	"regexp"
	"strings"
)

// Define the regex patterns as constants
const (
	// TODO: This should rather use a phone package to determine whether valid phone number for any region.
	phoneRegex = `^(?:\+254|254|0)?((?:7[0-9]{8})|(?:1[01][0-9]{7}))$`
)

// IsValidPhoneNumber checks if the given number is a valid phone number
func IsValidPhoneNumber(phonenumber string) bool {
	match, _ := regexp.MatchString(phoneRegex, phonenumber)
	return match
}

// FormatPhoneNumber formats a Kenyan phone number to "+254xxxxxxxx".
func FormatPhoneNumber(phone string) (string, error) {
	if !IsValidPhoneNumber(phone) {
		return "", errors.New("invalid phone number")
	}

	// Remove any leading "+" and spaces
	phone = strings.TrimPrefix(phone, "+")
	phone = strings.ReplaceAll(phone, " ", "")

	// Replace leading "0" with "254" if present
	if strings.HasPrefix(phone, "0") {
		phone = "254" + phone[1:]
	}
	// Add "+" if not already present
	if !strings.HasPrefix(phone, "254") {
		return "", errors.New("unexpected format")
	}

	return "+" + phone, nil
}
