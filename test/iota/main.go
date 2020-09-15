package main

import "fmt"

const (
	// iota 从0开始递增
	a = 1 << iota
	b
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}
