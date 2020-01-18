package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}

	for i, t := range a {
		t = 1
		a[i] = 2
		fmt.Println(t)
	}
	fmt.Println(a)
}
