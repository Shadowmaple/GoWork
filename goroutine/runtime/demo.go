// 验证 goroutine 是否只依赖于主 goroutine（main），
// 即在某 f() 中创建的 goroutine 是否会随着 f() 的结束而终止，
// 经检验，不会。

package main

import (
	"fmt"
	"time"
)

func main() {
	f()

	select {}
}

func f() {
	go func() {
		for {
			fmt.Println("hello, world.")
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("start..")
}
