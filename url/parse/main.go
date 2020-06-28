package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	uri := "postgres://user:pass@host.com:5432/path?k=-1#f"
	// uri := "https://www.baidu.com/"

	result, err := parse(uri)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func parse(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	fmt.Println(u)

	fmt.Println(u.Scheme)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)

	// fmt.Println(u.RawQuery)
	// m, _ := url.ParseQuery(u.RawQuery)
	// fmt.Println(m)
	// fmt.Println(m["k"][0])

	return "", nil
}
