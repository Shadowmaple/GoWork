package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	FormData = "application/x-www-form-urlencoded"
	JsonData = "application/json"
)

type RequestData struct {
	Query       map[string]string
	FormData    map[string]string
	ContentType string
}

// 通用发送 HTTP Request
func SendHTTPRequest(requestURL, method string, data RequestData) (string, error) {
	if len(data.Query) != 0 {
		requestURL += "?"
	}
	for key, value := range data.Query {
		requestURL += fmt.Sprintf("%s=%s&", key, value)
	}
	fmt.Println(requestURL)

	var payload string

	if data.ContentType == JsonData {
		body, err := json.Marshal(data.FormData)
		if err != nil {
			return "", err
		}
		payload = string(body)
	} else {
		body := url.Values{}
		for key, value := range data.FormData {
			body.Set(key, value)
		}
		payload = body.Encode()
	}

	fmt.Println(payload)

	req, err := http.NewRequest(method, requestURL, strings.NewReader(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", data.ContentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))
	return string(body), nil
}

func MarshalBodyForCustomData(body []byte, data interface{}) error {
	return json.Unmarshal(body, &data)
}
