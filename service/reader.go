package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA_server/model"
	"github.com/satori/go.uuid"
	_ "github.com/satori/go.uuid"
)

// ユーザーの回答データ送信
func SendUserAnswer(deviceId string, tagUuid string, oldUuid string) (model.Record, int) {
	var correct bool
	var record model.Record
	deviceInfo, find := GetDeviceInfoFromDeviceId(deviceId)
	if !find {
		return model.Record{}, -1
	}

	if oldUuid != "" {
		oldTag := GetTagDataFromUuid(tagUuid)
		if oldTag == nil {
			return model.Record{}, -2
		}
	}

	tagInfo := GetTagDataFromUuid(tagUuid)
	if tagInfo == nil {
		return model.Record{}, -3
	}

	// 問題タグ時
	if tagInfo.Sentence != "" {
		if tagUuid != oldUuid {
			tagInfo.Sentence = "前回の問題を回答してね"
		}
		bocco, find := GetDeviceInfoFromDeviceId(deviceId)
		if !find {
			return model.Record{}, -4
		}
		boccoInfo, find := ExisByBoccoAPI(bocco[0].Name)
		if find {
			boccoToken, _ := GetBoccoToken(boccoInfo[0].Email, boccoInfo[0].Key, boccoInfo[0].Pass)
			roomId, _ := GetRoomId(boccoToken)
			uuid := uuid.Must(uuid.NewV4()).String()
			SendMessage(uuid, roomId, boccoToken, tagInfo.Sentence)
		}
	} else {
		oldTagInfo := GetTagDataFromUuid(oldUuid)
		// 前回uuidが答えの時もしくは今回のuuidが前回の回答の時
		if oldTagInfo.Answer != "" || (oldTagInfo.BookId == tagInfo.BookId && oldTagInfo.QuestionNo == tagInfo.QuestionNo) {
			genreId := GetBookData(tagInfo.BookId)
			correctId := GetByCorrect(tagInfo.BookId, tagInfo.QuestionNo)
			if correctId == "" {
				return model.Record{}, -5
			}

			if tagInfo.Uuid == tagUuid {
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
				return model.Record{}, -6
			}

			if !correct {
				return record, 0
			}
		} else {
			return model.Record{}, 3
		}
	}

	return record, 1
}
