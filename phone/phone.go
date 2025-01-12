package phone

import (
	"regexp"
)

// Define the regex patterns as constants
const (
	phoneRegex   = `^(?:\+254|254|0)?((?:7[0-9]{8})|(?:1[01][0-9]{7}))$`
)

// IsValidPhoneNumber checks if the given number is a valid phone number
func IsValidPhoneNumber(phonenumber string) bool {
	match, _ := regexp.MatchString(phoneRegex, phonenumber)
	return match
}
