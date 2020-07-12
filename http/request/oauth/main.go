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

var (
	code         string
	refreshToken string
	clientID     = "4b194ad8-7d97-4dca-b078-6c3c65b31c75"
	clientSecret = "8c066b19-e507-4887-88f3-7e7edd99bfd8"
)

type RequestData struct {
	Query       map[string]string
	FormData    map[string]string
	ContentType string
}

type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// Data    map[string]interface{}
}

type CodeResponse struct {
	BasicResponse
	Data CodeItem `json:"data"`
}

type CodeItem struct {
	Code    string `json:"code"`
	Expired int64  `json:"expired"`
}

type TokenResponse struct {
	BasicResponse
	Data TokenItem `json:"data"`
}

type TokenItem struct {
	AccessToken    string `json:"access_token"`
	AccessExpired  int64  `json:"access_expired"`
	RefreshToken   string `json:"refresh_token"`
	RefreshExpired int64  `json:"refresh_expired"`
}

func main() {
	CodeRequest()
	TokenRequest()
	RefreshTokenRequest()
}

func TokenRequest() {
	url := "http://localhost:8083/auth/api/oauth/token"
	query := map[string]string{"client_id": clientID, "response_type": "token", "grant_type": "authorization_code"}
	formData := map[string]string{"code": code, "client_secret": clientSecret}

	s, err := SendHTTPRequest(url, "POST", RequestData{
		Query:       query,
		FormData:    formData,
		ContentType: FormData,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

	var data TokenResponse
	err = MarshalBodyForCustomData([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Data)
	refreshToken = data.Data.RefreshToken
}

func CodeRequest() {
	url := "http://localhost:8083/auth/api/oauth"
	query := map[string]string{"client_id": clientID, "response_type": "code"}
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

	var data CodeResponse
	err = MarshalBodyForCustomData([]byte(s), &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Data.Code)

	code = data.Data.Code
}

func RefreshTokenRequest() {
	// refreshToken = "MOGKAKGKU7IXFQNKMSTJWQ"
	url := "http://localhost:8083/auth/api/oauth/token/refresh"
	query := map[string]string{"client_id": clientID, "grant_type": "refresh_token"}
	bodyData := map[string]string{"refresh_token": refreshToken, "client_secret": clientSecret}

	s, err := SendHTTPRequest(url, "POST", RequestData{
		Query:       query,
		FormData:    bodyData,
		ContentType: FormData,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
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
