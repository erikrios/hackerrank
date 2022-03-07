package main

import "fmt"

func main() {
	fmt.Println(camelCase("saveChangesInTheEditor"))
}

func camelCase(s string) int32 {
	var counter int32 = 1

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			counter++
		}
	}

	return counter
}
