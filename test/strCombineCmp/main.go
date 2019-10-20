package main

import (
	"fmt"
	"time"
)

func runtime(msg string) func() {
	start := time.Now()
	return func() {
		fmt.Println(msg, time.Since(start))
	}
}

func plus(str string) {
	var s string
	for i := 0; i < 10000; i++ {
		s += str + ","
	}
	fmt.Println(s)
}

func sprint(str string) {
	var s string
	for i := 0; i < 10000; i++ {
		s = fmt.Sprintf("%s,%s", s, str)
	}
	fmt.Println(s)
}

// 测试目的：比较字符串的+号连接和Sprintf两者的效率
func main() {
	defer runtime("Runtime:")()

	plus("abc")
	// 33ms
	sprint("abc")
	// 35ms
	// 实践证明，直接+号连接字符串更高效
}
