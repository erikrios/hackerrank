package main

import "fmt"

func main() {
	fmt.Println(marsExploration("SOSTOT"))
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
	fmt.Println(marsExploration("SOSSOT"))
	fmt.Println(marsExploration("SOSSOSSOS"))
}

func marsExploration(s string) int32 {
	var counter int32

	for i := 0; i < len(s); i += 3 {
		chars := s[i : i+3]
		if string(chars) != "SOS" {
			if chars[0] != 'S' {
				counter++
			}

			if chars[1] != 'O' {
				counter++
			}

			if chars[2] != 'S' {
				counter++
			}
		}
	}

	return counter
}
