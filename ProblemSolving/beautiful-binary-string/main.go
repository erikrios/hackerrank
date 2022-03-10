package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(beautifulBinaryString("0101010"))
	fmt.Println(beautifulBinaryString("01100"))
	fmt.Println(beautifulBinaryString("0100101010"))
}

func beautifulBinaryString(b string) int32 {
	return int32(strings.Count(b, "010"))
}
