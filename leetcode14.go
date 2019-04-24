package main
import "fmt"

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
	//s := []string{"flower", "flow", "flight"}
	s := []string{}
	fmt.Println(longestCommonPrefix(s))
}
