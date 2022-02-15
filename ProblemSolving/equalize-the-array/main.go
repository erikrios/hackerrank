package main

func equalizeArray(arr []int32) int32 {
	arrMap := make(map[int32]int)

	for _, value := range arr {
		arrMap[value] = arrMap[value] + 1
	}

	var mostEqualValue int32

	for _, value := range arrMap {
		if int(mostEqualValue) < value {
			mostEqualValue = int32(value)
		}
	}

	return int32(len(arr)) - mostEqualValue
}
