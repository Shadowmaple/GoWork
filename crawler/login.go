package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"io/ioutil"
	"strings"
)

func main() {
	client := &http.Client{}
	requestUrl := "https://account.ccnu.edu.cn/cas/login?"
	r, _ := http.NewRequest("GET", requestUrl, nil)
	rp, err := client.Do(r)
	if err != nil {
		panic(err)
		return
	}
	body, _ := ioutil.ReadAll(rp.Body)
	defer rp.Body.Close()
	content := string(body)
	//fmt.Println(content)

	re := regexp.MustCompile(`name="lt" value="([\w-.]+)"`)
	params := re.FindStringSubmatch(content)
	lt := params[1]

	re = regexp.MustCompile(`name="execution" value="([\w]+)"`)
	params = re.FindStringSubmatch(content)
	execution := params[1]

	postData := url.Values{
		"username": {"***"},
		"password": {"***"},
		"lt": {lt},
		"execution": {execution},
		"_eventId": {"submit"},
		"submit": {"LOGIN"},
	}

	//cookie := "D7AB9DFCA0408CA3E6709F728C149DB4"
	cookie := rp.Cookies()[0].Value
	fmt.Println(rp.Cookies()[0].Value)
	//fmt.Println(rp.Header["Cookie"])
	//requestUrl = `https://account.ccnu.edu.cn/cas/login;jsessionid=` + cookie

	r, _ = http.NewRequest("POST", requestUrl, strings.NewReader(postData.Encode()))
	r.Header.Set("Cookie", "JSESSIONID=" + cookie)
	rp, err = client.Do(r)
	if err != nil {
		panic(err)
		return
	}
	body, _ = ioutil.ReadAll(rp.Body)
	defer rp.Body.Close()
	content = string(body)
	fmt.Println(content)
}
