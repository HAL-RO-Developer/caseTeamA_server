package controller

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA_server/model"
	"github.com/HAL-RO-Developer/caseTeamA_server/service"
	"github.com/gin-gonic/gin"
)

var Device = deviceimpl{}

type deviceimpl struct {
}

type deviceInfo struct {
	ChildId  int      `json:"child_id"`
	Nickname string   `json:"nickname"`
	Devices  []string `json:"child_devices"`
}

// デバイスID発行
func (d *deviceimpl) CreateNewDevice(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	req, ok := validation.ChildrenInfoValidation(c)
	if !ok {
		return
	}
	res, ok := service.CreateDevice(name, req.ChildId)
	if !ok {
		response.BadRequest(gin.H{"error": res}, c)
		return
	}

	timer := time.NewTimer(time.Minute * 10)
	go func() {
		<-timer.C
		_, find := service.GetPin(res)
		if find {
			service.PinRemove(res)
		}
	}()
	response.Json(gin.H{"pin": res}, c)
}

// デバイス一覧取得
func (d *deviceimpl) ListDevice(c *gin.Context) {
	var userDevices []deviceInfo
	var device deviceInfo
	var devices []model.Device
	var childData []model.UserChild

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	_, find := service.GetDeviceId(name)
	if !find {
		userDevices = []deviceInfo{}
		response.Json(gin.H{"devices": userDevices}, c)
		return
	}
	children, _ := service.GetChildInfo(name)
	/* 子どもIDの数繰り返し */
	for i := 0; i < len(children); i++ {
		devices, find = service.GetDeviceIdFromChildId(name, children[i].ChildId)
		if !find {
			device.ChildId = children[i].ChildId
			childData, _ = service.GetByChildInfo(name, children[i].ChildId)
			device.Nickname = childData[0].NickName
			device.Devices = []string{}
		} else {
			device.ChildId = children[i].ChildId
			childData, _ = service.GetByChildInfo(name, children[i].ChildId)
			device.Nickname = childData[0].NickName
			/* デバイスIDの数繰り返し */
			for j := 0; j < len(devices); j++ {
				device.Devices = append(device.Devices, devices[j].DeviceId)
			}
			userDevices = append(userDevices, device)
			device.Devices = nil
		}
	}
	response.Json(gin.H{"devices": userDevices}, c)
}

// デバイスID削除
func (d *deviceimpl) DeleteDevice(c *gin.Context) {
	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	deviceId := c.Param("device_id")

	if service.DeleteDeviceId(name, deviceId) {
		response.Json(gin.H{"success": "デバイスIDを削除しました。"}, c)
		return
	}
	response.BadRequest(gin.H{"error": "デバイスIDが見つかりません。"}, c)
}

// デバイスとIDの紐付け
func (d *deviceimpl) DeviceRegistration(c *gin.Context) {
	req, ok := validation.DeviceRegistrationCheck(c)
	if !ok {
		return
	}

	device_id, success := service.RegistrationDevice(req.Pin)
	if !success {
		response.BadRequest(gin.H{"error": device_id}, c)
		return
	}
	response.Json(gin.H{"device_id": device_id}, c)
}
