package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	strs := []string{}
	for _, v := range words {
		strs = append(strs, strings.ToLower(v))
	}
	return strs
}
