package main

import "fmt"

func main() {
	fmt.Println(stringConstruction("abcd"))
	fmt.Println(stringConstruction("abab"))
}

func stringConstruction(s string) int32 {
	mapper := make(map[byte]bool)
	var counter int32

	for i := 0; i < len(s); i++ {
		if _, ok := mapper[s[i]]; !ok {
			counter++
			mapper[s[i]] = true
		}
	}

	return counter
}
