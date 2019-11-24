package main

import (
	"fmt"
	"runtime"
//	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		go fmt.Print("-")
		fmt.Print("|")
//		time.Sleep(time.Second * 1)
	}
}

