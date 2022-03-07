package main

import "fmt"

func main() {
	numbers := []int32{2, 1, 3, 1, 2}
	fmt.Println(runningTime(numbers))
	fmt.Println(numbers)
	numbers = []int32{1, 1, 2, 2, 3, 3, 5, 5, 7, 7, 9, 9}
	fmt.Println(runningTime(numbers))
	fmt.Println(numbers)
}

func runningTime(arr []int32) int32 {
	return insertionSort(arr)
}

func insertionSort(numbers []int32) (totalShift int32) {
	if len(numbers) <= 1 {
		return
	}

	for i := 1; i < len(numbers); i++ {
		marker := i
		for j := i - 1; j >= 0; j-- {
			if numbers[marker] < numbers[j] {
				totalShift++
				numbers[marker], numbers[j] = numbers[j], numbers[marker]
				marker--
			} else {
				break
			}
		}
	}

	return totalShift
}
