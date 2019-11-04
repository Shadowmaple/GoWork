// 并发目录遍历
// 模拟 du 命令

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// 遍历以dir为根目录的整个文件树
func walkDir(dir string, fileSizes chan<- int64) {
	// fmt.Println(dir)
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
			continue
		}
		fileSizes <- entry.Size()
	}
}

// 返回dir目录中的条目
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树
	fileSizes := make(chan int64)
	wg := &sync.WaitGroup{}
	for _, root := range roots {
		wg.Add(1)
		go func(root string, fileSizes chan<- int64) {
			defer wg.Done()
			walkDir(root, fileSizes)
		}(root, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}

	fmt.Printf("%d files %d byte %.f GB\n", nfiles, nbytes, float64(nbytes)/1e9)
}
