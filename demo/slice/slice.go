package main

import (
	"fmt"
)

func main() {
	a := make([]int, 3)
	fmt.Println(cap(a), len(a))
	a = append(a, 1)
	fmt.Println(cap(a), len(a))
	a = append(a, 1)
	fmt.Println(cap(a), len(a))
	a = append(a, 1)
	fmt.Println(cap(a), len(a))
	a = append(a, 1)
	fmt.Println(cap(a), len(a))
	a = append(a, 1, 2, 3, 4)
	fmt.Println(cap(a), len(a))

	// len = cap = 1024
	b := make([]int, 1024)
	fmt.Println(cap(b), len(b))
	b = append(b, 1)
	fmt.Println(cap(b), len(b))

	// length = 1, cap = 1024
	// 必须满足：cap >= length
	c := make([]int, 1, 1024)
	fmt.Println(cap(c), len(c))
	c = append(c, 1)
	fmt.Println(cap(c), len(c))
}

/*
slice 扩容机制：
cap < 1024 时，扩容为原来的2倍，否则为1.25倍

make：
make(Type, size) 时，len=cap=size
make(Type, size1, size2) 时，len = size1, cap = size2，且 size1 <= size2
*/
