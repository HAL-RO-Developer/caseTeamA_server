package validation

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Tag struct {
	DeviceId string `json:"device_id"` // 読み取り機のdeviceID
	Uuid     string `json:"uuid"`
	OldUuid  string `json:"old_uuid"` // 前回回答データ
}

func ReaderValidation(c *gin.Context) (Tag, bool) {
	var req Tag
	err := c.BindJSON(&req)
	if err != nil || req.DeviceId == "" || req.Uuid == "" {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}

	return req, true
}
