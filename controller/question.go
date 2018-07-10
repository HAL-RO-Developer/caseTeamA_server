package controller

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA_server/service"
	"github.com/gin-gonic/gin"
)

var Question = questionimpl{}

type questionimpl struct {
}

type GenreInfo struct {
	GenreId   int    `json:"genre_id"`
	GenreName string `json:"genre_name"`
}

func (q *questionimpl) CreateQuestion(c *gin.Context) {
	req, success := validation.QuestionValidation(c)
	if !success {
		return
	}

	success = service.BookGenerate(req.BookId, req.GenreId)
	if !success {
		response.BadRequest(gin.H{"error": "本登録に失敗しました。"}, c)
		return
	}

	success = service.QuestionGenerate(req.BookId, req.QuestionNo, req.Sentence[0].TagId, req.Correct)
	if !success {
		response.BadRequest(gin.H{"error": "問題登録に失敗しました。"}, c)
		return
	}

	success = service.SentenceGenerate(req.Sentence[0].TagId, "", req.BookId, req.QuestionNo, req.Sentence[0].Text)
	if !success {
		response.BadRequest(gin.H{"error": "問題文登録に失敗しました。"}, c)
		return
	}

	success = service.CorrectGenerate(req)
	if !success {
		response.BadRequest(gin.H{"error": "回答登録に失敗しました。"}, c)
		return
	}
	response.Json(gin.H{"success": "問題情報を登録しました。"}, c)
}

// ジャンル作成
func (q *questionimpl) CreateGenre(c *gin.Context) {
	req, success := validation.GenreValidation(c)
	if !success {
		return
	}
	success = service.GenreGenerate(req)
	if !success {
		response.BadRequest(gin.H{"error": "ジャンル登録に失敗しました。"}, c)
		return
	}
	response.Json(gin.H{"success": "ジャンル情報を登録しました。"}, c)
}

// ジャンル情報取得
func (q *questionimpl) GetGenre(c *gin.Context) {
	var buf GenreInfo
	var res []GenreInfo

	genres := service.GetGenreNum()
	for i := 0; i < len(genres); i++ {
		buf.GenreId = genres[i].GenreId
		buf.GenreName = genres[i].GenreName
		res = append(res, buf)
		buf = GenreInfo{}
	}

	response.Json(gin.H{"genre": res}, c)
}
