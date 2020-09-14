package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "233A"
	s1 := strings.TrimSuffix(strings.TrimSuffix(s, "B"), "A")

	fmt.Println(s1)
}
