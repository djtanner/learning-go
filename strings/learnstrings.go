package main

import (
	"strings"
)

func ContainsZ(word string) bool {
	return strings.Contains(strings.ToLower(word), "z")
}

func Splitting(sentence string) []string {
	return strings.Split(sentence, " ")
}

func Lower(word string) string {

	return strings.ToLower(word)
}
