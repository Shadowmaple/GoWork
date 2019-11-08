package main

import (
	"fmt"
//	"time"
)

func main() {
	for {
		go fmt.Print("1")
		fmt.Print("0")
//		time.Sleep(time.Second * 1)
	}
}

