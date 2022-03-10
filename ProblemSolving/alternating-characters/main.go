package main

import "fmt"

func main() {
	fmt.Println(alternatingCharacters("AAAA"))
	fmt.Println(alternatingCharacters("BBBBB"))
	fmt.Println(alternatingCharacters("ABABABAB"))
	fmt.Println(alternatingCharacters("BABABA"))
	fmt.Println(alternatingCharacters("AAABBB"))
}

func alternatingCharacters(s string) int32 {
	starter := s[0]
	var count int32

	for i := 1; i < len(s); i++ {
		if s[i] != starter {
			starter = s[i]
		} else {
			count++
		}
	}

	return count
}
