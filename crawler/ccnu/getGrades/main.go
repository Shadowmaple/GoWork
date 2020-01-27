package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
)

func main() {

	url := "http://xk.ccnu.edu.cn/cjcx/cjcx_cxCjxq.html?time=" + strconv.Itoa(int(time.Now().UnixNano())) + "&gnmkdm=N305005"
	method := "POST"

	payload := strings.NewReader("jxb_id=8AF0256B86F10FA8E053CADCDC0A2E88&xnm=2019&xqm=3&kcmc=%E6%A6%82%E7%8E%87%E7%BB%9F%E8%AE%A1A")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "text/html, */*; q=0.01")
	req.Header.Add("Origin", "http://xk.ccnu.edu.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Cookie", "JSESSIONID=53874C814E6C443A86421EA9CD78078A")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	fmt.Println("--------------------")

	html := string(body)

	rg, err := regexp.Compile(`<td valign="middle">[^%]*&nbsp;</td>`)
	if err != nil {
		fmt.Println(err)
		return
	}
	s := rg.FindAllString(html, 2)
	fmt.Println(s)
}
