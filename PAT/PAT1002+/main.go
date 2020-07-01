package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str []byte
	var sum int
	var number = [...]string{"ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"}

	fmt.Scanf("%s", &str)
	//	fmt.Println(str)
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		sum += int(str[i]) - int('0')
	}

	sumArray := []byte(strconv.Itoa(sum))
	for j := 0; j < len(sumArray); j++ {
		if j+1 == len(sumArray) {
			fmt.Printf("%s\n", number[int(sumArray[j])-int('0')])
		} else {
			fmt.Printf("%s ", number[int(sumArray[j])-int('0')])
		}
	}
}
