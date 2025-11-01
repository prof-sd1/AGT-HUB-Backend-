package main

import (
	"fmt"
	"regexp"
	"strings"
)

// WordFrequency returns a map of word counts 
func WordFrequency(text string) map[string]int {
	wordCount := make(map[string]int)

	// Convert to lowercase
	text = strings.ToLower(text)

	// Remove punctuation using regex
	re := regexp.MustCompile(`[^\w\s]`)
	cleanText := re.ReplaceAllString(text, "")

	// Split into words
	words := strings.Fields(cleanText)

	// Count frequency
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
