package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	g := gin.Default()

	g.POST("", getData)

	g.Run(":8080")
}

func getData(c *gin.Context) {
	var err error

	// 获取head数据
	cookie := c.Request.Header.Get("cookie")
	if cookie != "" {
		log.Println(cookie)
	}

	var user User

	// 绑定json
	err = c.BindJSON(&user)
	if err != nil {
		log.Println("json get failed")
	}
	c.JSON(http.StatusOK, user.Username+" "+user.Password)

	// 读取form-data参数
	name := c.PostForm("name")
	log.Println(name)

	// 获取row数据
	// data, err := c.GetRawData()
}
