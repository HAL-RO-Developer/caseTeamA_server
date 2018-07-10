package main

import "github.com/HAL-RO-Developer/caseTeamA_server/model"

func main() {
	bookGenerate()
	genreGenerate()
	questionGenerate()
	sentenceGenerate()
	correctGenerate()
}

var db = model.GetDBConn()

// ドリル情報作成
func bookGenerate() {
	book := model.Book{
		BookId:  1,
		GenreId: 1,
	}

	book2 := model.Book{
		BookId:  2,
		GenreId: 2,
	}

	db.Create(&book)
	db.Create(&book2)
}

// ジャンル作成
func genreGenerate() {
	genre := model.Genre{
		GenreId:   1,
		GenreName: "算数",
	}

	genre2 := model.Genre{
		GenreId:   2,
		GenreName: "社会",
	}

	db.Create(&genre)
	db.Create(&genre2)
}

// 問題作成
func questionGenerate() {
	question := model.Question{
		BookId:     1,
		QuestionNo: 1,
		Sentence:   "question1_1",
		Correct:    "answer1_1_2",
	}
	question2 := model.Question{
		BookId:     1,
		QuestionNo: 2,
		Sentence:   "question1_2",
		Correct:    "answer1_2_1",
	}
	question3 := model.Question{
		BookId:     2,
		QuestionNo: 1,
		Sentence:   "question2_1",
		Correct:    "answer2_1_3",
	}
	db.Create(&question)
	db.Create(&question2)
	db.Create(&question3)
}

// 問題タグ
func sentenceGenerate() {
	tag := model.Tag{
		TagId:      "question1_1", // 任意のタグ番号
		Uuid:       "sample",      // tagのuuid
		BookId:     1,             // ドリルId
		QuestionNo: 1,             // ドリルの問題番号
		Sentence:   "1+1は?",       // 問題文
	}
	tag2 := model.Tag{
		TagId:      "question1_2",
		Uuid:       "sample_uuid",
		BookId:     1,
		QuestionNo: 2,
		Sentence:   "1x10は?",
	}
	tag3 := model.Tag{
		TagId:      "question2_1",
		Uuid:       "test_uuid",
		BookId:     2,
		QuestionNo: 1,
		Sentence:   "兵庫県の県庁所在地は?",
	}
	db.Create(&tag) // カラム作成
	db.Create(&tag2)
	db.Create(&tag3)
}

// 回答タグ
func correctGenerate() {
	tag := model.Tag{
		TagId:      "answer1_1_1", // 任意のタグ番号
		Uuid:       "test",        // タグのuuid
		BookId:     1,             // ドリルid
		QuestionNo: 1,             // ドリルの問題番号
		Answer:     "1",           // タグの回答
	}

	tag2 := model.Tag{
		TagId:      "answer1_1_2",
		Uuid:       "index",
		BookId:     1,
		QuestionNo: 1,
		Answer:     "2",
	}

	tag3 := model.Tag{
		TagId:      "answer1_1_3",
		Uuid:       "buf",
		BookId:     1,
		QuestionNo: 1,
		Answer:     "3",
	}

	tag4 := model.Tag{
		TagId:      "answer1_2_1",
		Uuid:       "test2",
		BookId:     1,
		QuestionNo: 2,
		Answer:     "10",
	}

	tag5 := model.Tag{
		TagId:      "answer1_2_2",
		Uuid:       "index2",
		BookId:     1,
		QuestionNo: 2,
		Answer:     "2",
	}

	tag6 := model.Tag{
		TagId:      "answer1_2_3",
		Uuid:       "buf2",
		BookId:     1,
		QuestionNo: 2,
		Answer:     "3",
	}

	tag7 := model.Tag{
		TagId:      "answer2_1_1",
		Uuid:       "test3",
		BookId:     2,
		QuestionNo: 1,
		Answer:     "兵庫市",
	}

	tag8 := model.Tag{
		TagId:      "answer2_1_2",
		Uuid:       "index3",
		BookId:     2,
		QuestionNo: 1,
		Answer:     "姫路市",
	}

	tag9 := model.Tag{
		TagId:      "answer2_1_3",
		Uuid:       "buf3",
		BookId:     2,
		QuestionNo: 1,
		Answer:     "神戸市",
	}

	db.Create(&tag)
	db.Create(&tag2)
	db.Create(&tag3)
	db.Create(&tag4)
	db.Create(&tag5)
	db.Create(&tag6)
	db.Create(&tag7)
	db.Create(&tag8)
	db.Create(&tag9)
}
