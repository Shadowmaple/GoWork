package main

import (
	"errors"
	"fmt"
)

// 查找子串，返回索引，模拟strings.Index功能
func Find(stra, strb string) (position int, err error) {
	if len(stra) < len(strb) {
		return -1, errors.New("stra should longer than strb ")
	}
	var p int

	// code...

	return p, nil
}

func main() {
	var stra, strb string
	fmt.Scanf("%s %s", stra, strb)
	po, err := Find(stra, strb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(po)
	}
}
