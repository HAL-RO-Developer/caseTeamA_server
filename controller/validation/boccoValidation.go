package validation

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Bocco struct {
	Email    string `json:"email"`
	Key      string `json:"key"`
	Password string `json:"pass"`
}

func BoccoValidation(c *gin.Context) (Bocco, bool) {
	var req Bocco
	err := c.BindJSON(&req)
	if err != nil || req.Email == "" || req.Key == "" || req.Password == "" {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}
