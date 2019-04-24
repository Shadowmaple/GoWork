package main
import "fmt"

func longestValidParentheses(s string) int {
	total := 0
	
	for begin, char := range s {
		if char != '(' {
			continue
		}
		flag := 1
		for _, ch := range s[begin+1:] {
			if ch == '(' {
				flag++
			} else if ch == ')' {
				flag--
			}
			if flag == 0 {
				total += 2
				break
			}
		}
	}
	return total
}

func main() {
	//var s string
	//fmt.Scanf(s)
	num := longestValidParentheses(")(())")
	fmt.Println(num)
}
