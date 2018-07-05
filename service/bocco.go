package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"encoding/json"

	"net/http/httputil"

	"github.com/HAL-RO-Developer/caseTeamA_server/model"
)

type GetToken struct {
	Token string `json:"access_token"`
}

var tokenURl = "https://api.bocco.me/alpha/sessions"
var roomIdURL = "https://api.bocco.me/alpha/rooms/joined"
var messageURL = "https://api.bocco.me/alpha/rooms"

// API情報
func ExisByBoccoAPI(name string) ([]model.Bocco, bool) {
	var boccos []model.Bocco
	db.Where("name = ?", name).Find(&boccos)
	return boccos, len(boccos) != 0
}

// BoccoAPIのトークン取得
func GetBoccoToken(email string, key string, pass string) (string, bool) {
	// httpリクエスト
	client := &http.Client{}

	values := url.Values{}
	values.Add("apikey", key)
	values.Add("email", email)
	values.Add("password", pass)

	req, err := http.NewRequest("POST", tokenURl, strings.NewReader(values.Encode()))
	if err != nil {
		return "", false
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// リクエスト実行
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}

	// JSONパース
	jsonBytes := ([]byte)(b)
	data := new(GetToken)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		return "", false
	}
	return data.Token, true
}

// ルームID取得
func GetRoomId(token string) (string, bool) {
	var data interface{}
	values := url.Values{}
	values.Add("access_token", token)

	resp, err := http.Get(roomIdURL + "?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", false
	}

	// JSONパース
	if err := json.Unmarshal(b, &data); err != nil {
		return "", false
	}

	roomId := data.([]interface{})[0].(map[string]interface{})["uuid"].(string)

	return roomId, true
}

// メッセージ送信
func SendMessage(uuid string, roomId string, token string, text string) bool {
	// idを生成
	values := url.Values{}
	values.Add("text", text)
	values.Add("unique_id", uuid)
	values.Add("media", "text")
	values.Add("access_token", token)

	req, err := http.NewRequest("POST", messageURL+"/"+roomId+"/messages", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Language", "ja")
	if err != nil {
		return false
	}

	client := new(http.Client)
	resp, err := client.Do(req)

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)

	return true
}

// BoccoAPI情報登録
func RegistrationBoccoInfo(name string, email string, key string, pass string) error {
	bocco := model.Bocco{
		Name:  name,
		Email: email,
		Key:   key,
		Pass:  pass,
	}
	err := db.Create(&bocco).Error
	return err
}

// BOCCOAPI情報削除
func DeleteBoccoInfo(name string) {
	var bocco model.Bocco
	db.Where("name = ?", name).First(&bocco)
	db.Delete(bocco)
}
