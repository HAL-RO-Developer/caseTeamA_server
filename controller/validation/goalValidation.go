package validation

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Goal struct {
	ChildId  int    `json:"child_id"`
	Content  string `json:"content"`
	Criteria int    `json:"criteria"`
	Deadline string `json:"deadline"`
}

type UpdateGoal struct {
	GoalId   string `json:"goal_id"`
	DeviceId string `json:"device_id"`
}

func GoalRegistrationValidation(c *gin.Context) (Goal, bool) {
	var req Goal
	err := c.BindJSON(&req)
	if err != nil || req.Content == "" || req.Criteria <= 0 {
		response.BadRequest(gin.H{"error": "未入力の項目があります。"}, c)
		return req, false
	}
	return req, true
}

func GoalUpdateValidation(c *gin.Context) (UpdateGoal, bool) {
	var req UpdateGoal
	err := c.BindJSON(&req)
	if err != nil || req.DeviceId == "" {
		response.BadRequest(gin.H{"error": "未入力の項目があります。"}, c)
		return req, false
	}
	return req, true
}
