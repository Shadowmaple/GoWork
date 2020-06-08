package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v4"
	"gopkg.in/oauth2.v4/errors"
	"gopkg.in/oauth2.v4/generates"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/models"
	"gopkg.in/oauth2.v4/server"
	"gopkg.in/oauth2.v4/store"
)

func main() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("store.db"))

	// client store
	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	// token generate
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("oauth"), jwt.SigningMethodHS512))

	clientStore.Set("test", &models.Client{
		ID:     "test",
		Secret: "2",
		Domain: "http://localhost:9094",
	})

	srv := server.NewDefaultServer(manager)
	srv.SetAllowedGrantType(oauth2.AuthorizationCode)
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
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	return "", "", err
		// }
		// fmt.Println(string(body))
		var info struct {
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		}
		// if err := json.Unmarshal(body, &info); err != nil {
		// 	return "", "", err
		// }
		info.ClientId = r.FormValue("client_id")
		info.ClientSecret = r.FormValue("client_secret")
		fmt.Println("client info: ", info.ClientId, info.ClientSecret)
		return info.ClientId, info.ClientSecret, nil
	})

	// ClientAuthorizedHandler check the client allows to use this authorization grant type
	srv.SetClientAuthorizedHandler(func(clientID string, grant oauth2.GrantType) (allowed bool, err error) {
		return true, nil
	})

	// http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	// 保存user id入上下文
	// 	store, err := session.Start(r.Context(), w, r)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	store.Set("loginUserId", r.Form.Get("username"))
	// 	store.Save()
	// 	w.Header().Set("Location", "/auth")
	// })

	// 请求 auth code
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		// 发送auth code

		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Println("request auth finished")
	})

	// 请求token
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		// srv.GetAccessToken()
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println("request token finished")
	})

	// 客户端备案
	http.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(string(body))
		var rp struct {
			ClientId string `json:"client_id"`
			Domain   string `json:"domain"`
		}
		if err := json.Unmarshal(body, &rp); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(rp)
		_, err = clientStore.GetByID(r.Context(), rp.ClientId)
		if err == nil {
			// 找到，已存在
			http.Error(w, err.Error(), 400)
			return
		}
		secret := generateSecret(10)
		clientStore.Set(rp.ClientId, &models.Client{
			ID:     rp.ClientId,
			Secret: secret,
			Domain: rp.Domain,
		})
		w.Write([]byte(secret))

		log.Println("request store finished")
	})

	log.Println("Server is running at 9096 port.")
	log.Fatal(http.ListenAndServe(":9096", nil))
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
