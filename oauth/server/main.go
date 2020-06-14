package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/oauth2.v4"
	"gopkg.in/oauth2.v4/errors"
	"gopkg.in/oauth2.v4/generates"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/models"
	"gopkg.in/oauth2.v4/server"
	"gopkg.in/oauth2.v4/store"
)

var (
	authCodeExp     = time.Hour * 10
	accessTokenExp  = time.Hour * 10
	refreshTokenExp = time.Hour * 10
	jwtKey          = "oauth"
)

func main() {

	manager := manage.NewDefaultManager()

	// 授权码模式下token配置
	manager.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp:    accessTokenExp,
		RefreshTokenExp:   refreshTokenExp,
		IsGenerateRefresh: true,
	})
	manager.SetAuthorizeCodeExp(authCodeExp)

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("store.db"))
	// manager.MapTokenStorage(tokenStore)
	// token generate
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(jwtKey), jwt.SigningMethodHS512))

	// client store
	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	clientStore.Set("test", &models.Client{
		ID:     "test",
		Secret: "2",
		Domain: "http://localhost:9094",
	})

	srv := server.NewDefaultServer(manager)
	srv.SetAllowedGrantType(oauth2.AuthorizationCode, oauth2.Refreshing)
	srv.SetAllowGetAccessRequest(false)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	// UserAuthorizationHandler get user id from request authorization
	srv.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		// return "", errors.ErrAccessDenied
		return "user", nil
	})

	// get client info (clientID and clientSecret)
	srv.SetClientInfoHandler(func(r *http.Request) (clientID, clientSecret string, err error) {
		clientID = r.FormValue("client_id")
		clientSecret = r.FormValue("client_secret")
		fmt.Println("client info: ", clientID, clientSecret)
		return
	})

	// ClientAuthorizedHandler check the client allows to use this authorization grant type
	srv.SetClientAuthorizedHandler(func(clientID string, grant oauth2.GrantType) (allowed bool, err error) {
		return true, nil
	})

	router := gin.Default()

	// 请求 auth code
	router.POST("/auth", func(g *gin.Context) {
		// 发送auth code

		err := srv.HandleAuthorizeRequest(g.Writer, g.Request)
		if err != nil {
			http.Error(g.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("request auth finished")
	})

	// 请求token
	// Params:
	//   grant_type: authorization_code
	//   response_type: token
	//   redirect_uri:
	// Forms:
	//   client_id
	//   client_secret
	//   code:
	router.POST("/token", func(g *gin.Context) {
		// err := srv.HandleTokenRequest(g.Writer, g.Request)
		// if err != nil {
		// 	http.Error(g.Writer, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		grantType, ok := g.GetQuery("grant_type")
		if !ok {
			g.String(http.StatusBadRequest, "grant_type is required")
			return
		} else if grantType != "authorization_code" {
			g.String(http.StatusBadRequest, "auth code grant")
			return
		}

		code, ok := g.GetPostForm("code")
		if !ok {
			g.String(http.StatusBadRequest, "code")
			return
		}

		clientID, clientSecret, err := srv.ClientInfoHandler(g.Request)
		if err != nil {
			g.String(http.StatusBadRequest, "client info required")
			return
		}

		tgr := &oauth2.TokenGenerateRequest{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Request:      g.Request,
			Code:         code,
		}

		ti, err := srv.GetAccessToken(g, oauth2.GrantType(grantType), tgr)
		if err != nil {
			g.String(500, "error")
			return
		}

		g.JSON(200, ti)
		fmt.Println("request token finished")
	})

	// Params:
	//   grant_type: refresh_token
	//   scope:
	// Forms:
	//   client_id
	//   client_secret
	//   refresh_token:
	router.POST("/refresh", func(g *gin.Context) {
		grantType, ok := g.GetQuery("grant_type")
		if !ok {
			g.String(http.StatusBadRequest, "grant_type is required")
			return
		} else if grantType != "refresh_token" {
			g.String(http.StatusBadRequest, "refresh token")
			return
		}

		refreshToken, ok := g.GetPostForm("refresh_token")
		if !ok {
			g.String(http.StatusBadRequest, "refresh token")
			return
		}

		clientID, clientSecret, err := srv.ClientInfoHandler(g.Request)
		if err != nil {
			g.String(http.StatusBadRequest, "client info required")
			return
		}

		tgr := &oauth2.TokenGenerateRequest{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Request:      g.Request,
			Refresh:      refreshToken,
		}

		ti, err := srv.GetAccessToken(g, oauth2.GrantType(grantType), tgr)
		if err != nil {
			g.String(500, "error")
			return
		}

		g.JSON(200, ti)

		fmt.Println("refresh token finished")
	})

	// 客户端备案
	router.POST("/store", func(g *gin.Context) {
		var rq struct {
			ClientId string `json:"client_id"`
			Domain   string `json:"domain"`
		}
		if err := g.BindJSON(&rq); err != nil {
			http.Error(g.Writer, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(rq)
		_, err := clientStore.GetByID(g, rq.ClientId)
		if err == nil {
			// 找到，已存在
			http.Error(g.Writer, err.Error(), 400)
			return
		}
		secret := generateUUID()
		clientStore.Set(rq.ClientId, &models.Client{
			ID:     rq.ClientId,
			Secret: secret,
			Domain: rq.Domain,
		})
		type rp struct {
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		}
		g.JSON(http.StatusOK, rp{
			ClientId:     rq.ClientId,
			ClientSecret: secret,
		})

		log.Println("request store finished")
	})

	log.Println("Server is running at 9096 port.")
	// log.Fatal(http.ListenAndServe(":9096", nil))
	router.Run(":9096")
}

func generateSecret(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func generateUUID() string {
	u := uuid.NewV4()
	return u.String()
}
