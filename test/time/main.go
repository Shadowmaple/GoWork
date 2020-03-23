package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now().Unix()
	timeStr := strconv.FormatInt(t, 10)
	fmt.Println(t, timeStr, len(timeStr))

	t1 := time.Now().Format("2006-01-02 15:04:05")
	t2 := time.Now().Format("2006/01/01")
	t3 := time.Now().Format("15:04:05")
	fmt.Printf("%s\n%s\n%s\n", t1, t2, t3)
}
