package main

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

type BasicResponse struct {
	Code    int
	Message string
	// Data    map[string]interface{}
	Data CodeResponse
}

type CodeResponse struct {
	Code    string
	Expired int64
}

type TokenResponse struct {
	AccessToken    string
	AccessExpired  int64
	RefreshToken   string
	RefreshExpired int64
}

func main() {
	CodeRequest()
	// TokenRequest()
}

func TokenRequest() {
	url := "http://localhost:8083/auth/api/oauth/token"
	query := map[string]string{"client_id": "a850da64-310e-416f-a6c3-ad9a7ad7eb25", "response_type": "token", "grant_type": "authorization_code"}
	formData := map[string]string{"code": "P9U6D6COPZYUNNV4WCTIMA", "client_secret": "2a839568-67a5-47b8-9027-00207b3d5072"}

	s, err := SendHTTPRequest(url, "POST", RequestData{
		Query:       query,
		FormData:    formData,
		ContentType: FormData,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func CodeRequest() {
	url := "http://localhost:8083/auth/api/oauth"
	query := map[string]string{"client_id": "a850da64-310e-416f-a6c3-ad9a7ad7eb25", "response_type": "code"}
	formData := map[string]string{"username": "shadow", "password": "MTIz"}

	s, err := SendHTTPRequest(url, "POST", RequestData{
		Query:       query,
		FormData:    formData,
		ContentType: JsonData,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	var data BasicResponse
	err = MarshalBodyForCustomData([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Data.Code)

	// if err := json.Unmarshal([]byte(s), &data); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data.Data)

	// code, ok := data.Data["code"].(string)
	// if !ok {
	// 	log.Fatal("error")
	// }
	// fmt.Println(code)
}

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

	// req.Header.Set("Content-Type", "multipart/form-data")
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
