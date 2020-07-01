package main

import "fmt"

//执行用时 : 4 ms, 在Longest Common Prefix的Go提交中击败了92.81% 的用户
//内存消耗 : 2.4 MB, 在Longest Common Prefix的Go提交中击败了48.92% 的用户
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	length := len(strs[0])
	for index := 0; index < length; index++ {
		prefix := strs[0][:index+1]
		for _, ch := range strs[1:] {
			if len(ch) < len(prefix) {
				return strs[0][:index]
			}
			if prefix != ch[:index+1] {
				return strs[0][:index]
			}
		}
	}
	return strs[0][:]
}

func main() {
	s := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(s))
}
