package main
import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)

	// 创建 10 个 goroutine，但是 onceBody 只会执行 1 次
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	// 等待 10 个 goroutine 结束
	for i := 0; i < 10; i++ {
		<-done
	}
}
