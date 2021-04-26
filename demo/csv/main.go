package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// generateFile()
	csvFileStream()
}

func generateFile() {
	f, err := os.Create("test.csv") //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data) // 写入数据，内部已包含 flush
	// w.Flush()
}

func csvFileStream() {
	w := csv.NewWriter(os.Stdout)
	data := [][]string{
		{"1", "中国", "23"},
		{"2", "美国", "23"},
		{"3", "bb", "23"},
		{"4", "bb", "23"},
		{"5", "bb", "23"},
	}
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		panic(err)
	}

	fmt.Println(w)
}
