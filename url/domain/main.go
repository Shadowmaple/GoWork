package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "http://www.baidu.com/"
	s, _ := parseDomain(url)
	fmt.Println(s)
}

func parseDomain(url string) (string, error) {
	a := strings.Split(url, "//")
	var s string
	if len(a) == 1 {
		s = strings.Split(a[0], "/")[0]
	} else {
		s = strings.Split(a[1], "/")[0]
	}
	return s, nil
}
