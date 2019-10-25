package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	time := time.Now().Unix()
	timeStr := strconv.FormatInt(time, 10)
	fmt.Println(time, timeStr, len(timeStr))
}
