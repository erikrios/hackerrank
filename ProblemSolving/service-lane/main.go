package main

import "fmt"

func main() {
	cases := [][]int32{
		{0, 3},
		{4, 6},
		{6, 7},
		{3, 5},
		{0, 7},
	}
	fmt.Printf("%+v", serviceLane(8, []int32{2, 3, 1, 2, 3, 2, 3, 3}, cases))
}

func serviceLane(n int32, width []int32, cases [][]int32) []int32 {
	result := make([]int32, 0)

	for _, caseRange := range cases {
		minValue := searchMin(width[caseRange[0] : caseRange[1]+1])
		result = append(result, minValue)
	}

	return result
}

func searchMin(values []int32) int32 {
	if len(values) == 0 {
		return 0
	}

	min := values[0]

	for _, value := range values {
		if min > value {
			min = value
		}
	}

	return min
}
