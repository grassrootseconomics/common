package identity

import (
	"regexp"
)
// Define the regex patterns as constants
const (
	addressRegex = `^0x[a-fA-F0-9]{40}$`
)

// IsValidAddress checks if the given address is a valid Ethereum address
func IsValidAddress(address string) bool {
	match, _ := regexp.MatchString(addressRegex, address)
	return match
}
