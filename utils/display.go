package utils

import (
	"fmt"
	"strings"
)
func DisplayText(input string, contentlines []string) {
	if input == "" {
		return
	}
if input == "\\n" || input == "\n" {
	fmt.Println()
	return
}
input = strings.ReplaceAll(input, "\n", "\\n")
input = strings.ReplaceAll(input, "\t", "    ")
wordSlice := strings.Split(input, "\\n")
for _, word := range wordSlice{
	if word == "" {
		fmt.Println()
	} else {
		if IsEnglish(word) {
			PrintWord(word, contentlines)
		} else {
			fmt.Println("invalid input:", word)
		}
	}
}
}