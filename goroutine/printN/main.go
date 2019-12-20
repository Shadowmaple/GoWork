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

// WaitGroup
func method1() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup

	go func() {
		for {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()

			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
		wg.Wait()
	}
}

// channel, CSP
func method2() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	ch := make(chan int64, 1)
	getCh := make(chan bool, 1)

	go func() {
		for {
			getCh <- true
			fmt.Println(<-ch)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
		select {
		case <-getCh:
			ch <- i
		default:
		}
	}
}

// 锁
func method3() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)

	var lock sync.Mutex

	go func() {
		for {
			lock.Lock()
			fmt.Println(i)
			lock.Unlock()
			time.Sleep(time.Second)
		}
	}()

	for {
		lock.Lock()
		i += 1
		lock.Unlock()
	}
}

func method4() {
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
		runtime.Gosched()
	}
}
