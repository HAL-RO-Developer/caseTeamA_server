package main

import (
	"github.com/HAL-RO-Developer/caseTeamA_server/model"
)

func main() {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.User{})
	db.DropTableIfExists(&model.UserChild{})
	db.DropTableIfExists(&model.Device{})
	db.DropTableIfExists(&model.Bocco{})
	db.DropTableIfExists(&model.Book{})
	db.DropTableIfExists(&model.Genre{})
	db.DropTableIfExists(&model.Question{})
	db.DropTableIfExists(&model.Record{})
	db.DropTableIfExists(&model.Tag{})
	db.DropTableIfExists(&model.CustomMessage{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserChild{})
	db.AutoMigrate(&model.Device{})
	db.AutoMigrate(&model.Bocco{})
	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Genre{})
	db.AutoMigrate(&model.Question{})
	db.AutoMigrate(&model.Record{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.CustomMessage{})
}
