package main

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type A struct{}

func (*A) Read() {}

// 用于检查是否实现接口
var _ Reader = &A{}

func main() {}
