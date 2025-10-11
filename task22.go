package main

import (
	"fmt"
	"regexp"
	"strings"
)

// IsPalindrome checks if a string is a palindrome (ignoring spaces, punctuation, and case)
func IsPalindrome(s string) bool {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove all non-alphanumeric characters
	re := regexp.MustCompile(`[^a-z0-9]`)
	cleaned := re.ReplaceAllString(s, "")

	// Compare with its reverse
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}
