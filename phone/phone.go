package phone

import (
	"errors"
	"regexp"
	"strings"
)

// Define the regex patterns as constants
const (
	// TODO: This should rather use a phone package to determine whether valid phone number for any region.
	// Kenyan phone numbers: must be exactly 10 digits (07XXXXXXXX or 01XXXXXXXX) when starting with 0
	// Or start with 254 / +254 and still follow the same pattern
	phoneRegex = `^(?:\+254|254|0)(7\d{8}|1[01]\d{7})$`
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

// FormatToLocalPhoneNumber converts a phone number like "+2547XXXXXXXX"
// or "2547XXXXXXXX" into the local format "07XXXXXXXX" / "01XXXXXXXX".
func FormatToLocalPhoneNumber(phone string) (string, error) {
	// Remove leading "+" and spaces
	phone = strings.TrimPrefix(phone, "+")
	phone = strings.ReplaceAll(phone, " ", "")

	// Must start with 254
	if !strings.HasPrefix(phone, "254") {
		return "", errors.New("invalid international phone format")
	}

	// Remove "254" prefix → get 7XXXXXXXX or 1XXXXXXXX
	rest := phone[3:]

	// Validate: must be 9 digits
	if len(rest) != 9 {
		return "", errors.New("invalid phone number length")
	}

	// Kenyan mobile numbers start with 7 or 1 (Safaricom/Airtel/Telkom prefixes)
	if !(strings.HasPrefix(rest, "7") || strings.HasPrefix(rest, "1")) {
		return "", errors.New("invalid Kenyan mobile prefix")
	}

	// Convert to local: prepend "0"
	local := "0" + rest

	return local, nil
}
