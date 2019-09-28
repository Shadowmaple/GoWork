package main

import (
	"fmt"
	_ "strings"
)

func main() {
	fmt.Print(longestPalindrome("ccc"))
}

func longestPalindrome(s string) string {
    if len(s) <= 1 {
        return s
    }
    max := 1
    k := 0
    for i := 1; i < len(s) - 1; i++ {
        for m, n := i-max, i+max; m >=0 && n < len(s) && s[m] == s[n]; m, n = m-1, n+1 {
			fmt.Println("----")
            max = n - m + 1
            k = n
        }
    }
	fmt.Println(k, max)
    return s[k:max]
}

