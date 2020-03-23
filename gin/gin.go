package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	route.GET("/user/:name/", func(c *gin.Context) {
		name := c.Param("name")
		message := "hello " + name
		c.String(http.StatusOK, message)
	})
	route.Run(":2000")
}
