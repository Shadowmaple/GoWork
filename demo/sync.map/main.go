package main

import (
	"fmt"
	"sync"
)

func main() {
	m := new(sync.Map)

	// 存储键值对
	m.Store("k", 1)
	m.Store("a", 2)
	m.Store("x", 3)

	// 从sync.Map中根据键取值
	fmt.Println(m.Load("a"))

	// 根据键删除对应的键值对
	m.Delete("a")

	// 遍历所有sync.Map中的键值对
	m.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
