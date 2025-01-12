package identity

import (
	"fmt"

	"git.grassecon.net/grassrootseconomics/common/phone"
)

// CheckRecipient validates the recipient format based on the criteria
func CheckRecipient(recipient string) (string, error) {
	if phone.IsValidPhoneNumber(recipient) {
		return "phone number", nil
	}

	if IsValidAddress(recipient) {
		return "address", nil
	}

	if IsValidAlias(recipient) {
		return "alias", nil
	}

	return "", fmt.Errorf("invalid recipient: must be a phone number, address or alias")
}
