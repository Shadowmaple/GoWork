package main

import (
	"fmt"
	"time"
)

// main函数中新开一个协程，协程中panic了，main函数defer中recover能接收到吗？
// 验证结果：不能，且会导致主函数panic异常退出
// 原因：a协程新开了一个协程b，这个协程b就独立于该a了（main除外），a结束后协程b还在运行，所以不会被defer捕获。
//      但它依赖于main协程，main结束后会强制b结束，所以b协程panic会导致main panic，但却也无法被main的defer捕获。
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("no panic")
		}
	}()

	go func() {
		panic("panic comes")
	}()

	// foo()
	time.Sleep(3 * time.Second)
	fmt.Println("main end")
}

func foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("no panic")
		}
	}()

	go func() {
		panic("panic comes")
	}()

	fmt.Println("foo end")
}
