package main

import (
	"fmt"
	"regexp"
)

func main() {
	room := "国4-0114照明公用"
	s := ProcessDormName(room)
	fmt.Println(s)
}

func ProcessDormName(dorm string) string {
	s := `(.+)-.*([1-9][\d]{2,3}[A|B]?).*`

	rgx := regexp.MustCompile(s)
	result := rgx.FindStringSubmatch(dorm)

	fmt.Println(result)

	return ""
}
