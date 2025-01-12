package identity

import (
	"regexp"
)

const (
	aliasRegex   = `^[a-zA-Z0-9]+$`
)

// IsValidAlias checks if the alias is a valid alias format
func IsValidAlias(alias string) bool {
	match, _ := regexp.MatchString(aliasRegex, alias)
	return match
}
