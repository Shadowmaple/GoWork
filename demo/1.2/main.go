package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for index, cmd := range os.Args[:] {
		fmt.Println(index, cmd)
	}
	fmt.Println("执行的命令是：" , os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}