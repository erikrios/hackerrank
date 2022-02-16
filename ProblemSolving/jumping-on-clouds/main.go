package main

import "fmt"

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 1, 0, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 0, 1, 0}))
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0}))
}

func jumpingOnClouds(c []int32) int32 {
	var position int
	var count int32

	for position < len(c)-1 {

		if len(c) == 1 {
			return count
		}

		if len(c) == 2 {
			if c[position+1] == 1 {
				return count
			} else {
				return count + 1
			}
		}

		if position == len(c)-1-2 && c[len(c)-1] == 1 {
			count++
			break
		}

		if position == len(c)-1-1 && c[len(c)-1] == 1 {
			count++
			break
		} else if position == len(c)-1-1 && c[len(c)-1] == 0 {
			count++
			break
		}

		if c[position+2] == 1 {
			position++
		} else {
			position += 2
		}
		count++
	}

	return count
}
