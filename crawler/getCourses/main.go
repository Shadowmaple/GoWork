package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	GetCourses()
}

func GetCourses() {
	sid := "2018214830"
	cookie := "JSESSIONID=9775CD2E38E1E9F85B6C3E07D5C28606"
	url := "http://xk.ccnu.edu.cn/xkcx/xkmdcx_cxXkmdcxIndex.html?doType=query&gnmkdm=N255010&su=" + sid

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Origin", "http://xk.ccnu.edu.cn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Add("Cookie", cookie)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(string(body))
}
