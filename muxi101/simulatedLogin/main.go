package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/howeyc/gopass"
)

func main() {
	requestUrl := "https://accounts.douban.com/j/mobile/login/basic"

	// 输入账号和密码
	var name string
	fmt.Print("豆瓣账号：")
	_, _ = fmt.Scanln(&name)
	fmt.Print("输入密码：")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}

	data := url.Values{}
	data.Set("name", name)
	data.Set("password", string(password))
	data.Set("remember", "false")
	data.Set("ck", "")
	data.Set("ticket", "")

	payload := strings.NewReader(data.Encode())

	request, err := http.NewRequest("POST", requestUrl, payload)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Origin", "https://accounts.douban.com")
	request.Header.Add("Referer", "https://accounts.douban.com/passport/login_popup?login_source=anony")
	request.Header.Add("Sec-Fetch-Mode", "cors")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)

	fmt.Println(string(body))
}
