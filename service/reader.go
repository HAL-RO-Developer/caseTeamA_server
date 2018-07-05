package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA_server/model"
	"github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
)

// ユーザーの回答データ送信
func SendUserAnswer(deviceId string, tagUuid string) (model.Record, int) {
	var correct bool
	var record model.Record
	deviceInfo, find := GetDeviceInfoFromDeviceId(deviceId)
	if !find {
		return model.Record{}, -1
	}

	tagInfo := GetTagDataFromUuid(tagUuid)
	if tagInfo == nil {
		return model.Record{}, -2
	}

	// 問題タグ時
	if tagInfo.Sentence != "" {
		bocco, find := GetDeviceInfoFromDeviceId(deviceId)
		if !find {
			return model.Record{}, -3
		}
		boccoInfo, find := ExisByBoccoAPI(bocco[0].Name)
		if find {
			boccoToken, _ := GetBoccoToken(boccoInfo[0].Email, boccoInfo[0].Key, boccoInfo[0].Pass)
			roomId, _ := GetRoomId(boccoToken)
			uuid := uuid.Must(uuid.NewV4()).String()
			SendMessage(uuid, roomId, boccoToken, tagInfo.Sentence)
		}
	} else {
		genreId := GetBookData(tagInfo.BookId)
		correctId := GetByCorrect(tagInfo.BookId, tagInfo.QuestionNo)
		if correctId == "" {
			return model.Record{}, -4
		}
		if correctId == tagUuid {
			correct = true
		} else {
			correct = false
		}
		record = model.Record{
			Name:       deviceInfo[0].Name,
			ChildId:    deviceInfo[0].ChildId,
			AnswerDay:  time.Now(),
			BookId:     tagInfo.BookId,
			QuestionNo: tagInfo.QuestionNo,
			GenreId:    genreId[0].GenreId,
			UserAnswer: tagInfo.Answer,
			Correct:    correct,
		}
		err := db.Create(&record).Error
		if err != nil {
			return model.Record{}, -5
		}

		if !correct {
			return record, 0
		}
	}

	return record, 1
}

// チャレンジ判定
