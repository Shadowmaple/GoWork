package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

    // 运行一个协程
	go func() {
		defer wg.Done()

		fmt.Println(1)
		fmt.Println(2)
        fmt.Println(3)
	}()

	// 运行第二个协程
	go func() {
		defer wg.Done()

		fmt.Println("A")
		fmt.Println("B")

		// 设置个睡眠，让该协程执行超时而被挂起，引起超时调度
        time.Sleep(time.Second)

		fmt.Println("C")
	}()

	wg.Wait()
}

