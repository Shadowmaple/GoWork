package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var ch = [10]string{"ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"}
	var sum = 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	fmt.Println(sum)
	l := len(strconv.Itoa(sum))
	for i, s := range strconv.Itoa(sum) {
		m, _ := strconv.Atoi(string(s))
		fmt.Printf("%s", ch[m])
		if i != l-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
