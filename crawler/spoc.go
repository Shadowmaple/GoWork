package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	client := http.Client{}
	requestUrl := "http://spoc.ccnu.edu.cn/userLoginController/getUserProfile"
	postData := make(url.Values)
	postData.Add("loginName", "***")
	postData.Add("password", "***")

	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(postData.Encode()))
	rp, err := client.Do(request)
	if err != nil {
		panic(err)
		return
	}

	//content, _ := ioutil.ReadAll(rp.Body)
	fmt.Println(rp.StatusCode, rp.Status)
}
