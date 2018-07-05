package validation

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/gin-gonic/gin"
)

type UserChildren struct {
	BirthDay string `json:"birthday"`
	NickName string `json:"nickname"`
	Sex      int    `json:"sex"`
	DeviceId string `json:"device_id"`
}

type childId struct {
	ChildId int `json:"child_id"`
}

func ChildrenInfoValidation(c *gin.Context) (childId, bool) {
	var req childId
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}

func AddChildInfoValidation(c *gin.Context) (UserChildren, bool) {
	var req UserChildren
	err := c.BindJSON(&req)
	if err != nil || req.BirthDay == "" || req.NickName == "" {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	if !(req.Sex == 0 || req.Sex == 1) {
		response.BadRequest(gin.H{"error": "性別が不正です。"}, c)
		return req, false
	}
	return req, true
}
