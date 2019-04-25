package main

import (
	"fmt"
)

func main() {
	var n, sum int

	fmt.Scan(&n)
	sum = 0

	for ; n > 1; sum++ {
		if n % 2 == 0 {
			n /= 2
		} else {
			n = (3*n + 1) / 2
		}
	}
	fmt.Println(sum)
}
