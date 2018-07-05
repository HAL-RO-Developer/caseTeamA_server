package controller

import (
	"time"

	"strconv"

	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA_server/model"
	"github.com/HAL-RO-Developer/caseTeamA_server/service"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
)

var User = userimpl{}

type userimpl struct {
	ChildId  int       `json:"child_id"`
	BirthDay time.Time `json:"birthday"`
	NickName string    `json:"nickname"`
	Sex      int       `json:"sex"`
}

func (u *userimpl) Create(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		response.BadRequest(gin.H{"error": "ユーザー名またはパスワードが未入力です。"}, c)
		return
	}
	if service.User.ExisByName(user.Name) {
		response.BadRequest(gin.H{"error": "登録済みのユーザー名です。"}, c)
	} else {
		user = service.User.Store(user)
		response.Json(gin.H{"success": "ユーザー登録を行いました。"}, c)
	}
}

// ユーザー削除(チームA)
func (u *userimpl) UserDeleteForWork(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	works, find := service.ExisByRecord(name)
	if find {
		for i := 0; i < len(works); i++ {
			//service.DeleteUserAnswer(works[i].DeviceId)
		}
	}

	if service.DeleteUser(name) {
		response.Json(gin.H{"success": "ユーザー情報を削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "ユーザー情報が見つかりませんでした。"}, c)
}

// 子供情報取得
func (u *userimpl) GetChildren(c *gin.Context) {
	var children []userimpl
	var child userimpl

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	buf, find := service.FindByUserName(name)
	if !find {
		response.Json(gin.H{"children": []userimpl{}}, c)
		return
	}

	for i := 0; i < len(buf); i++ {
		child.ChildId = buf[i].ChildId
		child.BirthDay = buf[i].BirthDay
		child.NickName = buf[i].NickName
		child.Sex = buf[i].Sex
		children = append(children, child)
	}
	response.Json(gin.H{"children": children}, c)
}

// 子供情報追加
func (u *userimpl) Child(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.AddChildInfoValidation(c)
	if !ok {
		return
	}

	childId, success := service.AddChild(name, req)
	if !success {
		response.BadRequest(gin.H{"error": "登録失敗"}, c)
		return
	}
	response.Json(gin.H{"child_id": childId}, c)
}

// 指定された子ども情報の削除
func (u *userimpl) DeleteChild(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	childId, err := strconv.Atoi(c.Param("child_id"))
	if err != nil {
		response.BadRequest(gin.H{"error": "リクエストが不正です。"}, c)
		return
	}

	devices, _ := service.GetDeviceIdFromChildId(name, childId)
	for i := 0; i < len(devices); i++ {
		service.DeleteDeviceId(name, devices[i].DeviceId)
	}

	workMessages, find := service.GetWorkMessageFromNameChild(name, childId)
	if find {
		for i := 0; i < len(workMessages); i++ {
			service.DeleteWorkMessageFromChild(name, childId)
		}
	}

	success := service.DeleteChild(name, childId)
	if !success {
		response.BadRequest(gin.H{"error": "子どもIDが見つかりませんでした。"}, c)
		return
	}
	response.Json(gin.H{"success": "削除しました。"}, c)
}

//	トークンの検証
func authorizationCheck(c *gin.Context) (string, bool) {
	token := c.GetHeader("Authorization")

	userInfo, err := jwt.Decode(token)
	if err != nil {
		return "error", false
	}

	_, ok := service.User.Login(userInfo["name"], userInfo["pass"])
	if !ok {
		return "error", false
	}

	return userInfo["name"], true
}
