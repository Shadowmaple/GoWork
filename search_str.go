package main

import "fmt"

func Find(stra, strb string) (position int, err error) {
	for index,
}

func main() {
	var stra, strb string
	fmt.Scanf("%s %s", stra, strb)
	po, err := Find(stra, strb)
	if err != nil {
		return -1, err
	} else {
		return po, nil
	}
}

