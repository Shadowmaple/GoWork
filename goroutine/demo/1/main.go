package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var ss = []string{"hello", "dd"}

	for _, s := range ss {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 3)
			fmt.Println(s)
		}()
	}
	wg.Wait()
	fmt.Println("OK")
}
