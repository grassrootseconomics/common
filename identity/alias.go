package identity

import (
	"regexp"
)

const (
	aliasRegex = `^(?i)[a-z][a-z0-9]*(\.[a-z0-9]+)*$`
)

// IsValidAlias checks if the alias is a valid alias format
func IsValidAlias(alias string) bool {
	match, _ := regexp.MatchString(aliasRegex, alias)
	return match
}
