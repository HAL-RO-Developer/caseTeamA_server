package middleware

import (
	"net/http"

	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA_server/model"
	"github.com/HAL-RO-Developer/caseTeamA_server/service"
	"github.com/gin-gonic/gin"
	"github.com/makki0205/gojwt"
)

func Login(c *gin.Context) {
	var req model.User
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, ok := service.User.Login(req.Name, req.Password)
	if !ok {
		response.BadRequest(gin.H{"error": "ログインエラー"}, c)
		return
	}
	claims := map[string]string{
		"name": user.Name,
		"pass": user.Password,
	}
	token := jwt.Generate(claims)

	response.Json(gin.H{"token": token}, c)
}
