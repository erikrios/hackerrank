package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%+v\n", cutTheSticks([]int32{5, 4, 4, 2, 2, 8}))
	fmt.Printf("%+v\n", cutTheSticks([]int32{1, 2, 3, 4, 3, 3, 2, 1}))
}

var maxValue = math.MaxInt32

func cutTheSticks(arr []int32) []int32 {
	result := make([]int32, 0)

	result = append(result, int32(len(arr)))

	for i := 0; i < len(arr); i++ {
		minValue := findMin(arr)
		if int(minValue) == maxValue {
			break
		}

		nonZeroCount := subtractElements(&arr, minValue)

		if nonZeroCount != 0 {
			result = append(result, nonZeroCount)
		}
	}

	return result
}

func subtractElements(arr *[]int32, subractValue int32) (nonZeroCounter int32) {

	for i, value := range *arr {

		if value == 0 {
			continue
		} else {
			(*arr)[i] = value - subractValue
			if (*arr)[i] != 0 {
				nonZeroCounter++
			}
		}
	}

	return
}

func findMin(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}

	min := maxValue

	for _, value := range arr {
		if value == 0 {
			continue
		}
		if min > int(value) {
			min = int(value)
		}
	}

	return int32(min)
}
