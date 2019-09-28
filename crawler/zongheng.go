package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	t := time.Now()
	var lock sync.Mutex
	var wait sync.WaitGroup

	var keyword string
	keyword = "火影"
	//fmt.Scanln(&keyword)

	file, _ := os.OpenFile("zongheng.txt", os.O_WRONLY | os.O_CREATE, 0666)
	num := 1

	for i := 1; i <= 10; i++ {
		requestUrl := "http://search.zongheng.com/s?keyword=" + keyword + "&pageNo=" + strconv.Itoa(i) + "&sort="
		wait.Add(1)
		go func () {
			defer wait.Done()

			rp, _ := http.Get(requestUrl)
			body, _ := ioutil.ReadAll(rp.Body)
			content := string(body)
			defer rp.Body.Close()

			dom, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
			dom.Find(".search-tab").Each(func(i int, selection *goquery.Selection){
				// fmt.Println(selection.Text())
				wait.Add(1)
				go func () {
					defer wait.Done()
					selection.Find(".tit").Each(func(i int, title *goquery.Selection) {
						lock.Lock()
						defer lock.Unlock()
						//fmt.Println(title.Text())
						fmt.Fprintf(file, "%3d   ", num)
						fmt.Fprintln(file, title.Text())
						num++
					})
				}()
				})
		}()

	}
	wait.Wait()
	println("elapsed:", time.Since(t))
}
