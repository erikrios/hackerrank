package main

import "fmt"

func main() {
	fmt.Println(saveThePrisoner(5, 2, 1))
	fmt.Println(saveThePrisoner(5, 2, 2))
	fmt.Println(saveThePrisoner(7, 19, 2))
	fmt.Println(saveThePrisoner(3, 7, 3))
}

func saveThePrisoner(n int32, m int32, s int32) int32 {
	return (s+m-2)%n + 1
}
