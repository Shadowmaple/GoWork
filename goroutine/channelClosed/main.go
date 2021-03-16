package main

import (
	"fmt"
	"time"
)

// 测试关闭一个 channel，会从关闭的 channel 中收到仅一个关闭信息还是都收到
func main() {
	stopChan := make(chan int)
	go func() {
		for {
			select {
				case <- stopChan:
					fmt.Println("1 closed.")
					return
				default:
					fmt.Println("1")
			}
		}
	}()

	go func() {
		for {
			select {
				case <- stopChan:
					fmt.Println("2 closed.")
					return
				default:
					fmt.Println("2")
			}
		}
	}()

	go func() {
		for {
			select {
				case <- stopChan:
					fmt.Println("3 closed.")
					return
				default:
					fmt.Println("3")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("stop..")
	close(stopChan)
	time.Sleep(2 * time.Second)
}
