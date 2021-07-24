package main

import "fmt"

type A struct {
	X string
}

// 会先调用return，再调用defer
func foo() (ret A) {
	defer func() {
		fmt.Println(ret.X)
	}()
	return A{X: "hao"}
}

func main() {
	a := foo()
	fmt.Println(a.X)
}
