package controller

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA_server/service"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

var Reader = readerimpl{}

type readerimpl struct {
}

type reader struct {
	deviceId string
	Status   string
}

// リーダーで読み取った情報の送信
func (r *readerimpl) SendTag(c *gin.Context) {
	req, ok := validation.ReaderValidation(c)
	if !ok {
		response.BadRequest(gin.H{"error": "不正なリクエストです。"}, c)
		return
	}

	device, find := service.GetDeviceInfoFromDeviceId(req.DeviceId)

	if find {
		_, result := service.SendUserAnswer(req.DeviceId, req.Uuid)
		if result < 0 {
			response.BadRequest(gin.H{"error": "回答データを保存できませんでした。"}, c)
			return
		}

		// oldRecord, find := service.GetByRecordFromChild(device[0].Name, device[0].ChildId)
		boccoInfo, find := service.ExisByBoccoAPI(device[0].Name)
		message, find := service.GetMessageInfoFromTrue(device[0].Name, device[0].ChildId, result)
		if !find {
			if result == 0 {
				response.Json(gin.H{"success": false}, c)
			} else {
				response.Json(gin.H{"success": true}, c)
			}
			return
		}

		boccoToken, _ := service.GetBoccoToken(boccoInfo[0].Email, boccoInfo[0].Key, boccoInfo[0].Pass)
		roomId, _ := service.GetRoomId(boccoToken)
		uuid := uuid.Must(uuid.NewV4()).String()
		service.SendMessage(uuid, roomId, boccoToken, message[0].Message)
		if result == 0 {
			response.Json(gin.H{"success": false}, c)
			return
		}
	}
	response.Json(gin.H{"success": true}, c)
}
