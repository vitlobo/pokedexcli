package util

import "strings"

// Capitalizes only the first letter of a string
func TitleCase(s string) string {
	if len(s) == 0 || s == "" { return s }
	return strings.ToUpper(s[:1]) + (s[1:])
}