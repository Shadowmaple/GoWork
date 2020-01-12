package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func main() {
	s, err := HashCourseId("48700001", "2014980017/周光有/教授, 2006982136/彭近兰/副教授")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(s)
}

func HashCourseId(courseNumStr, teachersStr string) (string, error) {
	sqs := strings.Split(teachersStr, ",")
	var array []string
	for _, s := range sqs {
		array = append(array, strings.Split(s, "/")[1])
	}
	teachers := strings.Join(array, ",")
	fmt.Println(teachers)

	// 生成一个MD5的哈希
	hash := md5.New()
	// 写入哈希
	hash.Write([]byte(courseNumStr + teachers))
	// hash算法
	id := hash.Sum(nil)
	// fmt.Println(hash, id)

	// 检验生成的序列是否相同
	for i := 0; i < 3; i++ {
		fmt.Println(hex.EncodeToString(id))
		fmt.Printf("%x\n", md5.Sum(id)) // 十六进制输出
	}

	// 将字节流转化为16进制的字符串
	return hex.EncodeToString(id), nil
}
