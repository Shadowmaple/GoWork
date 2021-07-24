package main

func foo() int {
	return 3
}

func test() {
	a := 1
	defer func() {
		println(a)
	}()
	a = foo()
}

func main() {
	test()
}
