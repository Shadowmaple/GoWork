package main

import (
	"fmt"
	"time"
)

func main() {
	// t := GetCurrentTime()
	// fmt.Println(t)

	// TimeFormat()

	ok := CheckDateTime("2020-02-13")
	fmt.Println(ok)
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

func ParseTimeString(s string) (time.Time, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation("2006-01-02", s, loc)
}

// 格式化，提取标准时间：2020-01-01
func CheckDateTime(dates ...string) bool {
	for _, date := range dates {
		_, err := ParseTimeString(date)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}
