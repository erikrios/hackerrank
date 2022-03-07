package main

import "fmt"

func main() {
	fmt.Println(minimumDistances([]int32{3, 2, 1, 2, 3}))
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
}

func minimumDistances(a []int32) int32 {
	minDistance := len(a) + 1

	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				if j-i < minDistance {
					minDistance = j - i
				}
				break
			}
		}
	}

	if minDistance == len(a)+1 {
		return -1
	}

	return int32(minDistance)
}
