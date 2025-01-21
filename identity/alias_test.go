package identity

import (
	"testing"
)

func TestAliasesRegext(t *testing.T) {
	tests := []struct {
		alias   string
		isValid bool
	}{
		{
			alias:   "foobar",
			isValid: true,
		},
		{
			alias:   "foo.sarafu.local",
			isValid: true,
		},
		{
			alias:   "foo.sarafu",
			isValid: true,
		},
		{
			alias:   "foo.",
			isValid: false,
		},
		{
			alias:   ".foo..bar",
			isValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.alias, func(t *testing.T) {
			isValid := IsValidAlias(tt.alias)
			if isValid != tt.isValid {
				t.Fatalf("expected %v, got %v", tt.isValid, isValid)
			}
		})
	}
}
