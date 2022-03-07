package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(funnyString("acxz"))
	fmt.Println(funnyString("bcxz"))
}

func funnyString(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if math.Abs(float64(s[i+1])-float64(s[i])) != math.Abs(float64(s[len(s)-i-1])-float64(s[len(s)-i-2])) {
			return "Not Funny"
		}
	}

	return "Funny"
}
