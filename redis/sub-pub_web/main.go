package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var ChStr = "chan"

func main() {
	// redis sub/pub
	SubRdb = OpenRedisClient()
	PubRdb = OpenRedisClient()

	defer PubRdb.Close()
	defer SubRdb.Close()

	go Sub()

	// gin web-service
	g := gin.Default()

	g.POST("", Pub)

	g.Run()
}

func Pub(c *gin.Context) {
	msg := c.DefaultQuery("msg", "none")

	if err := PubRdb.Publish(ChStr, msg).Err(); err != nil {
		c.String(500, "pub failed")
		log.Println(err)
		return
	}

	c.String(200, msg)
}

func Sub() {
	sub := SubRdb.Subscribe(ChStr)

	ch := sub.Channel()

	for msg := range ch {
		log.Println("sub get -- ", msg.Payload)
	}
}
