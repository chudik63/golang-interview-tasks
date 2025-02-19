package main

import "fmt"

func main() {
	slice := make([]int, 3, 4) // {[0, 0, 0] len = 3 cap = 4}
	appendingSlice(slice[:1])
	fmt.Println(slice) // {[0, 1, 0] len = 3 cap = 4}
}

func appendingSlice(slice []int) {
	// slice - {[0] len = 1 cap = 4}
	slice = append(slice, 1)
	// slice - {[0, 1]] len = 2 cap = 4}
}
