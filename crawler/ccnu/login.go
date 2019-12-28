package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {
	client := &http.Client{}
	requestUrl := "https://account.ccnu.edu.cn/cas/login"
	// 初始化 http request
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		panic(err)
	}

	// 发起请求
	rp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	// 读取body
	body, err := ioutil.ReadAll(rp.Body)
	defer rp.Body.Close()
	if err != nil {
		panic(err)
	}

	content := string(body)
	//fmt.Println(content)

	var JSESSIONID string
	var lt string
	var execution string
	// var _eventId string

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range rp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}
	if JSESSIONID == "" {
		log.Println("Can not get JSESSIONID")
		return
	}
	// 正则匹配 HTML 返回的表单字段
	//ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	//executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	//_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	re := regexp.MustCompile(`name="lt" value="([\w-.]+)"`)
	params := re.FindStringSubmatch(content)
	lt = params[1]

	re = regexp.MustCompile(`name="execution" value="([\w]+)"`)
	params = re.FindStringSubmatch(content)
	execution = params[1]

	postData := url.Values{
		"username":  {"***"},
		"password":  {"***"},
		"lt":        {lt},
		"execution": {execution},
		"_eventId":  {"submit"},
		"submit":    {"LOGIN"},
	}

	//cookie := "D7AB9DFCA0408CA3E6709F728C149DB4"
	cookie := rp.Cookies()[0].Value
	fmt.Println(rp.Cookies()[0].Value)
	//fmt.Println(rp.Header["Cookie"])
	//requestUrl = `https://account.ccnu.edu.cn/cas/login;jsessionid=` + cookie

	r, _ := http.NewRequest("POST", requestUrl, strings.NewReader(postData.Encode()))
	r.Header.Set("Cookie", "JSESSIONID="+cookie)
	rp, err = client.Do(r)
	if err != nil {
		panic(err)
	}
	body, _ = ioutil.ReadAll(rp.Body)
	content = string(body)
	fmt.Println(content)
}
