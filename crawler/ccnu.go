package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	requestUrl := "https://account.ccnu.edu.cn/cas/login?"
	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return
	}
	statusCode := resp.StatusCode
	content := string(data)
	fmt.Println(statusCode)
	fmt.Println(content)
}