package securityCheck

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

// QQ小程序内容安全检测

type accessTokenManager struct {
	Token     string
	CreateAt  time.Time
	ExpiresIn time.Duration
}

var (
	QQAppSecret string
	QQAppID     string

	accessToken = &accessTokenManager{}

	imgSecCheckURL    = "https://api.q.qq.com/api/json/security/ImgSecCheck?access_token="
	msgSecCheckURL    = "https://api.q.qq.com/api/json/security/MsgSecCheck?access_token="
	accessTokenGetURL = "https://api.q.qq.com/api/getToken?grant_type=client_credential&appid=%s&secret=%s"
)

func (t *accessTokenManager) loadToken() error {
	resp, err := http.Get(fmt.Sprintf(accessTokenGetURL, QQAppID, QQAppSecret))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var obj struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int32  `json:"expires_in"`
		ErrCode     int32  `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
	}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		return err
	}

	t.Token = obj.AccessToken
	t.CreateAt = time.Now().UTC()
	t.ExpiresIn = time.Duration(obj.ExpiresIn)

	return nil
}

func (t *accessTokenManager) check() error {
	now := time.Now()
	if t.CreateAt.Add(t.ExpiresIn).Sub(now) <= 0 {
		err := t.loadToken()
		if err != nil {
			return err
		}
	}
	return nil
}

func QQSecInit() {
	QQAppID = viper.GetString("QQ_APPID")
	QQAppSecret = viper.GetString("QQ_APP_SECRET")

	accessToken.loadToken()

	imgSecCheckURL += accessToken.Token
	msgSecCheckURL += accessToken.Token
}

type imgCheckReq struct {
	AppID string `json:"appid"`
}

type msgCheckReq struct {
	AppID   string `json:"appid"`
	Content string `json:"content"`
}

type checkResponse struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

// 消息文本检测
func MsgSecCheck(content string) error {
	data, err := json.Marshal(msgCheckReq{
		AppID:   QQAppID,
		Content: content,
	})
	if err != nil {
		return err
	}

	// fmt.Println(string(data))

	resp, err := http.Post(msgSecCheckURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("QQ msg security check err", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// fmt.Println(string(body))

	var rp checkResponse
	if err := json.Unmarshal(body, &rp); err != nil {
		return err
	}

	// fmt.Println(rp)
	if rp.ErrCode != 0 {
		log.Printf(fmt.Sprintf("msg security check failed. code: %d; msg: %s.", rp.ErrCode, rp.ErrMsg))
		return nil
	}

	return nil
}

// 图片检测
func ImgSecCheck(image string) error {
	return nil
}

// func main() {
// 	QQSecInit()

// 	// msg := "特3456书yuuo莞6543李zxcz蒜7782法fgnv级"
// 	msg := "完47知qwez到"
// 	if err := MsgSecCheck(msg); err != nil {
// 		panic(err)
// 	}
// }
