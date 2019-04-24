package main
import "fmt"

func longestValidParentheses(s string) int {
	total := 0
	length := len(s)
	for i:=0, j:=length-1; i<j {
		for ; s[i:i+1]=='('; i++ {
			;
		}
		for ; s[j:j+1]==')'; j-- {
			;
		}
	}
}

func main() {
	var s string

	fmt.Scanf(s)
	num := longestValidParentheses(s)
	fmt.Println(num)
}
