package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(theLoveLetterMystery("abc"))
	fmt.Println(theLoveLetterMystery("abcba"))
	fmt.Println(theLoveLetterMystery("abcd"))
	fmt.Println(theLoveLetterMystery("cba"))
}

func theLoveLetterMystery(s string) int32 {
	var counter int32

	for i := 0; i < len(s)/2; i++ {
		lastChar := s[len(s)-1-i]
		if s[i] != lastChar {
			res := int(s[i]) - int(lastChar)
			counter += int32(math.Abs(float64(res)))
		}
	}

	return counter
}
