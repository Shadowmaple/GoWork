package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientModel struct {
	ClientId int
	SecretId string
	Addr     string
	AuthCode string
}

var store = make(map[int]ClientModel)

func main() {
	g := gin.Default()

	g.POST("/apply", Apply)
	g.POST("/auth", Auth)
	g.POST("/token", Token)
	g.GET("/info", UserInfo)

	g.Run(":9999")
}

// 客户端备案
func Apply(g *gin.Context) {
	log.Println("apply")

	clientId := 22

	store[clientId] = ClientModel{
		ClientId: clientId,
		SecretId: "9sdffd",
		Addr:     "localhost",
	}
	g.JSON(200, clientId)
}

// request: login form
// return: auth code
func Auth(g *gin.Context) {
	log.Println("Auth")
	clientIdStr, ok := g.GetQuery("client_id")
	if !ok {
		g.JSON(400, "error")
		return
	}

	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		g.JSON(400, "error")
		return
	}

	authCode := "22222323asdf"
	m := store[clientId]
	m.AuthCode = authCode
	store[clientId] = m

	g.JSON(200, authCode)
}

// request: auth code
// return: token
func Token(g *gin.Context) {
	log.Println("token")
	clientIdStr, ok := g.GetQuery("client_id")
	if !ok {
		g.JSON(400, "error")
		return
	}

	clientId, err := strconv.Atoi(clientIdStr)
	if err != nil {
		g.JSON(400, "error")
		return
	}

	// 需要验证 secret id
	// ...

	authCode := g.Query("auth")
	if store[clientId].AuthCode != authCode {
		g.JSON(400, "error")
		return
	}

	token := "this." + authCode + ".test"

	g.JSON(200, token)
}

// request: token
// return: user info
func UserInfo(g *gin.Context) {

}
