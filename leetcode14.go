package main
import "fmt"

func longestCommonPrefix(strs []string) string {	
	length := len(strs[0])
	for index := 0; index < length; index++ {
		prefix := strs[0][:index+1]
		for _, ch := range strs[1:] {
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
