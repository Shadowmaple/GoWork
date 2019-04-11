package main
import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	//迭代获取字符串的字符
	for _, n := range s {
		fmt.Println(string(n))
	}
}
