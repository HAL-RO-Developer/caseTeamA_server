package controller

import (
	"time"

	"strconv"

	"github.com/HAL-RO-Developer/caseTeamA_server/controller/response"
	"github.com/HAL-RO-Developer/caseTeamA_server/model"
	"github.com/HAL-RO-Developer/caseTeamA_server/service"
	"github.com/gin-gonic/gin"
)

var Record = recordimpl{}

type recordimpl struct {
}

type recordGraphDate struct {
	Date    time.Time `json:"date" sql:"not null;type:date"`
	NumAns  int       `json:"num_ans"`
	NumCorr int       `json:"num_corr"`
}

type recordGraphGenre struct {
	Genre       string `json:"genre"`
	BookId      int    `json:"book_id"`
	NumProblems int    `json:"num_problems"`
	NumAns      int    `json:"num_ans"`
	NumCorr     int    `json:"num_corr"`
}

type recordInfo struct {
	Date      time.Time    `json:"date" sql:"type:date"`
	GenreName string       `json:"genre_name"`
	Detail    []recordData `json:"detail"`
}

type recordData struct {
	Sentence   string `json:"sentence"`
	UserAnswer string `json:"user_answer"`
	Correct    string `json:"correct"`
	Result     bool   `json:"result"`
}

const layout = "2006-01-02"

// 日付けorジャンルごとのグラフ用回答記録取得
// Todo 同一問題を回答した時は最新の問題のみ(日付け別も)
func (r *recordimpl) WorkRecordForGraph(c *gin.Context) {
	var userDate []recordGraphDate
	var userGenre []recordGraphGenre
	var recordDate recordGraphDate
	var recordGenre recordGraphGenre
	var childRecord []model.Record
	num := 0

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	buf := c.Param("child_id")
	childId, err := strconv.Atoi(buf)
	if err != nil {
		response.BadRequest(gin.H{"error": "childIdが不正です。"}, c)
		return
	}
	filter := c.Query("filter")
	if !(filter == "date" || filter == "genre") {
		response.BadRequest(gin.H{"error": "クエリパラメータが不正です。"}, c)
		return
	}
	records, find := service.GetByRecordFromChild(name, childId)
	// 回答情報が見つからなかった時
	if !find {
		response.Json(gin.H{"records": []string{}}, c)
		return
	}

	// 日付けごとの時
	if filter == "date" {
		for i := 0; len(records) > i; i += len(childRecord) {
			recordDate.Date = records[i].AnswerDay
			childRecord, find = service.GetByRecordFromDay(name, childId, recordDate.Date)
			if !find {
				userDate = []recordGraphDate{}
			} else if find {
				recordDate.NumAns = len(childRecord)
				for j := 0; len(childRecord) > j; j++ {
					if childRecord[j].Correct {
						recordDate.NumCorr++
					}
				}
				userDate = append(userDate, recordDate)
				recordDate = recordGraphDate{}
			}
			num++
		}
		response.Json(gin.H{"records": userDate}, c)
		return
	}
	number := service.GetGenreNumber()
	for i := 1; i < number; i++ {
		childRecord, find = service.GetByRecordFromGenre(name, childId, i)
		if !find {
			userGenre = []recordGraphGenre{}
		}
		buf := service.GetGenreData(i)
		recordGenre.Genre = buf[0].GenreName
		recordGenre.BookId = childRecord[i].BookId

		recordGenre.NumProblems = service.GetByQuestion(childRecord[i].BookId)
		if find {
			recordGenre.NumAns = len(childRecord)
			for j := 0; len(childRecord) > j; j++ {
				if childRecord[j].Correct {
					recordGenre.NumCorr++
				}
			}
			userGenre = append(userGenre, recordGenre)
			recordGenre = recordGraphGenre{}
		}
	}
	response.Json(gin.H{"records": userGenre}, c)
}

// 回答情報詳細取得
func (r *recordimpl) WorkRecordForDetail(c *gin.Context) {
	var childDetail recordInfo
	var detail recordData
	var tagData []model.Tag
	var correctId string
	var correct *model.Tag
	var childRecord []model.Record
	var find bool

	name, ok := authorizationCheck(c)
	if !ok {
		response.TokenError(gin.H{"error": "アクセストークンが不正です。"}, c)
		return
	}

	buf := c.Param("child_id")
	childId, err := strconv.Atoi(buf)
	if err != nil {
		response.BadRequest(gin.H{"error": "childIdが不正です。"}, c)
		return
	}

	day := c.DefaultQuery("date", "")
	date, err := time.Parse(layout, day)
	hoge := c.DefaultQuery("genre", "")
	genre, error := strconv.Atoi(hoge)
	if day == "" {
		if error != nil {
			response.BadRequest(gin.H{"error": "クエリパラメータが不正です。"}, c)
			return
		}
	} else if hoge == "" {
		if err != nil {
			response.BadRequest(gin.H{"error": "クエリパラメータが不正です。"}, c)
			return
		}
	}

	if day != "" && hoge == "" {
		childRecord, find = service.GetByRecordFromDay(name, childId, date)
		if find {
			childDetail.Date = date
		}
	} else if day == "" && hoge != "" {
		childRecord, find = service.GetByRecordFromGenre(name, childId, genre)
		if find {
			childDetail.GenreName = service.GetGenreName(genre)
		}
	} else {
		childRecord, find = service.GetByRecordFromGenreDate(name, childId, date, genre)
		if find {
			childDetail.Date = date
			childDetail.GenreName = service.GetGenreName(genre)
		}
	}

	if len(childRecord) == 0 {
		childDetail.Detail = []recordData{}
	} else {
		for i := 0; i < len(childRecord); i++ {
			tagData = service.GetTagDataFromBookId(childRecord[i].BookId, childRecord[i].QuestionNo)
			detail.Sentence = tagData[0].Sentence
			detail.UserAnswer = childRecord[i].UserAnswer
			correctId = service.GetByCorrect(childRecord[i].BookId, childRecord[i].QuestionNo)
			if correctId == "" {
				response.BadRequest(gin.H{"error": "問題が見つかりませんでした。"}, c)
				return
			}
			correct = service.GetTagDataFromTagId(correctId)
			detail.Correct = correct.Answer
			detail.Result = childRecord[i].Correct
			childDetail.Detail = append(childDetail.Detail, detail)
			detail = recordData{}
		}
	}

	response.Json(gin.H{"records": childDetail}, c)
}
