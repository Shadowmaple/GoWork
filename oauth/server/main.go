package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v4/errors"
	"gopkg.in/oauth2.v4/generates"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/models"
	"gopkg.in/oauth2.v4/server"
	"gopkg.in/oauth2.v4/store"
)

func main() {
	manager := manage.NewDefaultManager()
	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	r := gin.New()

	r.POST("/login", func(g *gin.Context) {
		username := g.PostForm("username")
		psw := g.PostForm("password")
		fmt.Println("== ", username, psw)

		// 外部重定向
		g.Redirect(http.StatusPermanentRedirect, "/auth")
	})

	r.POST("/auth", func(g *gin.Context) {
		err := srv.HandleAuthorizeRequest(g.Writer, g.Request)
		if err != nil {
			g.Error(err)
		}
		// g.String(200, "OK")
	})

	r.POST("/token", func(g *gin.Context) {
		err := srv.HandleTokenRequest(g.Writer, g.Request)
		if err != nil {
			g.Error(err)
		}
		// g.String(200, "ok")
	})

	r.POST("/apply", func(g *gin.Context) {
		// ...
	})

	r.Run(":9096")
}
