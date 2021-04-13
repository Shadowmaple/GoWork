package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	s := "站内-横版视频-0313（优选）-SH270-"
	symbol, id, err := GetSymbol(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(symbol, id)
}

// GetSymbol 解析计划名称中的特殊字符
// symbol：特殊字符，如 "-SH181-"
// id：特殊字符中的房间号/分区号
func GetSymbol(s string) (symbol string, id int64, err error) {
	// 只解析 -SHxxx- 或 -SRxxx-
	rgx, err := regexp.Compile("(-S[H|R](\\d*)-)")
	if err != nil {
		return
	}
	result := rgx.FindStringSubmatch(s)
	if len(result) < 3 {
		err = errors.New("匹配解析失败")
		return
	}
	symbol = result[1]
	id, err = strconv.ParseInt(result[2], 10, 32)
	if err != nil {
		return
	}
	return
}
