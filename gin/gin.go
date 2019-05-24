package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.GET("/user/:name/", func(c *gin.Context) {
		name := c.Param("name")
		message := "hello " + name
		c.String(http.StatusOK, message)
	})
	route.Run()
}
