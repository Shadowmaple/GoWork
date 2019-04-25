package main
import "fmt"

func main() {
	/*
	var s string
	fmt.Scan(&s)

	//迭代获取字符串的字符
	for _, n := range s {
		fmt.Println(string(n))
	}
	*/

	/*
	a := "sdfda"
	b := "sd"
	fmt.Println(a[:2], b, a[:2]==b)
	*/

	
	s := "adfaf"
	for index, ch := range s[2:] {
		fmt.Println(index, ch)
	}
	
}
