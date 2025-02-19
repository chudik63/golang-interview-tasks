package main

import "fmt"

func main() {
	var x []int
	x = append(x, 0) // x = [0] len = 1 cap = 1
	x = append(x, 1) // x = [0 1] len = 2 cap = 2
	x = append(x, 2) // x = [0 1 2] len = 3 cap = 4
	y := append(x, 3)
	// x = [0 1 2] len = 3 cap = 4       y = [0 1 2 3] len = 4 cap = 4
	z := append(x, 4)
	// x = [0 1 2] len = 3 cap = 4       y = [0 1 2 4] len = 4 cap = 4            z = [0 1 2 4] len = 5 cap = 8

	fmt.Println(x, y, z) // [0 1 2] [0 1 2 4] [0 1 2 4]
}
