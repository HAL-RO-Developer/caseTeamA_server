package validation

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/gin-gonic/gin"
)

type Question struct {
	BookId     int       `json:"book_id"`
	QuestionNo int       `json:"question_no"`
	Sentence   []TagInfo   `json:"sentence"`
	Answer     []TagInfo `json:"answer"`
	Correct    string    `json:"correct"`
	GenreId    int       `json:"genre_id"`
}

type TagInfo struct {
	TagId string `json:"tag_id"`
	Text  string `json:"text"`
}

func QuestionValidation(c *gin.Context) (Question, bool) {
	var req Question
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}

func GenreValidation(c *gin.Context) (string, bool) {
	var req string
	err := c.BindJSON(&req)
	if err != nil {
		response.BadRequest(gin.H{"error": "入力されていないデータがあります。"}, c)
		return req, false
	}
	return req, true
}
