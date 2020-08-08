package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/lexkong/log"
	"golang.org/x/net/publicsuffix"
)

func GetCardInfo(sid, password string) error {
	params, err := MakeAccountPreflightRequest()
	if err != nil {
		log.Error("MakeAccountPreflightRequest function error", err)
		return err
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}

	if err := MakeAccountRequest(sid, password, params, client); err != nil {
		log.Error("MakeAccountRequest function err", err)
		return err
	}

	// fmt.Println(client.Jar)

	requestURL := "http://one.ccnu.edu.cn/ecard_portal/get_info"
	req, err := http.NewRequest("POST", requestURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Origin", "http://xk.ccnu.edu.cn")
	req.Header.Set("Host", "xk.ccnu.edu.cn")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Request error", err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
