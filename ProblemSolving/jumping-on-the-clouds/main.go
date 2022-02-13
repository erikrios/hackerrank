package main

import "fmt"

func main() {
	fmt.Println(jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3))
}

func jumpingOnClouds(c []int32, k int32) int32 {
	var energy int32 = 100

	var i int32 = 0

	i = k

	for {
		if i == int32(len(c)) {
			decreaseEnergy(&energy, k, c[0])
			break
		}

		if i > int32(len(c)) {
			i = i - int32(len(c))
		}

		decreaseEnergy(&energy, k, c[i])
		i += k
	}

	return energy
}

func decreaseEnergy(energy *int32, k int32, v int32) {
	if v == 0 {
		*energy = *energy - 1
	} else {
		*energy = *energy - 1 - 2
	}
}
