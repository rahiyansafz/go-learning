package main

import (
	"fmt"
	"strings"
)

func ManipulateStrings() {
	fruit := " Apple "
	index, err := findLetterIndex(fruit, "A")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("The index of 'l' in '%s' is: %d\n", strings.TrimSpace(fruit), index)
	}
}

func findLetterIndex(s, letter string) (int, error) {
	trimmedS := strings.TrimSpace(s)
	index := strings.Index(trimmedS, letter)
	if index == -1 {
		return -1, fmt.Errorf("string '%s' does not contain the letter '%s'", trimmedS, letter)
	}
	return index, nil
}
