package main

import "fmt"

func main() {
	fmt.Println(minimumDistances([]int32{3, 2, 1, 2, 3}))
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
}

func minimumDistances(a []int32) int32 {
	var minDistance int32 = int32(len(a) + 1)
	mapper := make(map[int32]int32)

	for i := 0; i < len(a); i++ {
		element := a[i]
		if v, ok := mapper[element]; ok {
			if distance := int32(i) - v; distance < minDistance {
				minDistance = distance
			}
		}
		mapper[element] = int32(i)
	}

	if minDistance == int32(len(a)+1) {
		return -1
	}

	return minDistance
}
