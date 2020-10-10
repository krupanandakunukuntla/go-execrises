package main

import "fmt"

func main() {
	x, y := []int{1, 2}, []int{1, 2}
	s := equal(x, y)
	fmt.Println(s)
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] { // slices must be compared through individual element
			return false
		}
	}
	return true
}
