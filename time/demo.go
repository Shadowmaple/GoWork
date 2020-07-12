package main

import (
	"fmt"
	"time"
)

func main() {
	t := GetCurrentTime()
	fmt.Println(t)

	TimeFormat()
}

func GetCurrentTime() *time.Time {
	//	var loc, _ = time.LoadLocation("Asia/Shanghai")
	//	loc := time.FixedZone("CST", 7*3600)
	t := time.Now().UTC().Add(8 * time.Hour)
	return &t
}

func TimeFormat() {
	now := time.Now()
	s := now.Format("2006/01/02 15:04:05")
	fmt.Println(s)
}
