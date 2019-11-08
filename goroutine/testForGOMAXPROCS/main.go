package main

import (
	"fmt"
	"runtime"
//	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for {
		go fmt.Print("1")
		fmt.Print("0")
//		time.Sleep(time.Second * 1)
	}
}

