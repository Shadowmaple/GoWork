package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	a := uniquePaths(m-1, n)
	b := uniquePaths(m, n-1)
	return a + b
}

func main() {
	n := 3
	m := 7
	x := uniquePaths(m, n)
	fmt.Println(x)
}
