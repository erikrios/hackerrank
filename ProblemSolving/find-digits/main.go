package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findDigits(12))
	fmt.Println(findDigits(1012))
}

func findDigits(n int32) int32 {
	values := toSlice(n)

	var counter int32
	for _, value := range values {
		if value == 0 {
			continue
		}

		if n%value == 0 {
			counter++
		}
	}

	return counter
}

func toSlice(n int32) []int32 {
	stringValues := strconv.Itoa(int(n))

	result := make([]int32, 0)

	for _, stringValue := range stringValues {
		value, _ := strconv.Atoi(string(stringValue))
		result = append(result, int32(value))
	}

	return result
}
