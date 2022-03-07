package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize"))
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))
}

func pangrams(s string) string {
	s = strings.ToLower(s)

	alphabets := make(map[rune]int)

	for i := 'a'; i <= 'z'; i++ {
		alphabets[i] = 0
	}

	for _, char := range s {
		if _, ok := alphabets[char]; ok {
			alphabets[char]++
		}
	}

	for _, count := range alphabets {
		if count == 0 {
			return "not pangram"
		}
	}

	return "pangram"
}
