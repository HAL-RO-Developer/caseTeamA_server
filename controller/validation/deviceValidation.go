package validation

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/gin-gonic/gin"
)

type PostDevice struct {
	Pin string `json:"pin"`
}

func DeviceRegistrationCheck(c *gin.Context) (PostDevice, bool) {
	var req PostDevice
	err := c.BindJSON(&req)
	if err != nil || req.Pin == "" {
		response.BadRequest(gin.H{"error": "pinが未入力です。"}, c)
		return req, false
	}
	return req, true
}
