package main
import "fmt"

func longestValidParentheses(s string) int {
	max := 0
	
	for begin, char := range s {
		if char != '(' {
			continue
		}
		flag := 1
		for index, ch := range s[begin+1:] {
			if ch == '(' {
				flag++
			} else if ch == ')' {
				flag--
			}
			if flag == 0 {
				total := index + 1 + 1
				if total > max {
					max = total
				}
				break
			}
		}
	}
	return max
}

func main() {
	//var s string
	//fmt.Scanf(s)
	num := longestValidParentheses(")(())")
	fmt.Println(num)
}
