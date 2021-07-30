package flag

import (
	"regexp"
	"strings"
)

// ToField converts a flag string to a struct field name.
func ToField(flag string) string {
	return strings.Join(strings.Split(strings.Title(flag), "-"), "")
}

// ToFlag converts a struct field name to a flag string .
func ToFlag(field string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	matches := re.FindAllString(field, -1)
	return strings.ToLower(strings.Join(matches, "-"))
}
