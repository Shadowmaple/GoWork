/*
现在有一个程序：在main goroutine中对于变量i不断进行修改，另外开一个goroutine每隔一秒钟将i的值进行输出
如果不出意外，应该会每隔一秒输出一个-1，这和我们预想中的不太一样。

任务要求便是：使用不同的思路改写程序，使之可以达到预期的效果，即每隔一段时间可以输出不断增长的值
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	go func() {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
	}
}

func method1() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	go func() {
		for {
			i += 1
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
	}
}

func method2() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	ch := make(chan int64, 1)

	go func() {
		for {
			fmt.Println(<-ch)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
		ch <- i
	}
}

func method3() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	for {
		i += 1

		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(i)
			time.Sleep(time.Second)
		}()
		wg.Wait()
	}
}
