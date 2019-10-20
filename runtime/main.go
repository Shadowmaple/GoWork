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

func main() {
	defer runtime("Runtime:")()
	time.Sleep(5 * time.Second)
}
