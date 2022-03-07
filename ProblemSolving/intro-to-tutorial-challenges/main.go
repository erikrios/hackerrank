package main

import "fmt"

func main() {
	fmt.Println(introTutorial(4, []int32{1, 4, 5, 7, 9, 12}))
}

func introTutorial(V int32, arr []int32) int32 {
	for i, val := range arr {
		if V == val {
			return int32(i)
		}
	}

	return -1
}
