package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA_server/model"
)

var User = user{}

type user struct {
}

var db = model.GetDBConn()

func (u *user) Store(user model.User) model.User {
	db.Create(&user)
	return user
}

func (u *user) ExisByName(name string) bool {
	var users []model.User
	db.Where("name = ?", name).Find(&users)
	return len(users) != 0
}

func (u *user) Login(name, pass string) (*model.User, bool) {
	var users []model.User
	db.Where("name = ?", name).Find(&users)
	if len(users) == 0 {
		return nil, false
	}
	return &users[0], users[0].Password == pass
}

// ユーザー情報削除
func DeleteUser(name string) bool {
	var user model.User
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return false
	}
	db.Delete(&user)
	return true
}

// 子ども情報取得
func GetChildInfo(name string) ([]model.UserChild, bool) {
	var children []model.UserChild
	err := db.Where("name = ?", name).Find(&children).Error
	if err != nil {
		return nil, false
	}
	return children, true
}

// 最初に見つけたこどもID情報削除
func DeleteChildFirst(name string) bool {
	var child model.UserChild
	err := db.Where("name = ?", name).First(&child).Error
	if err != nil {
		return false
	}
	db.Delete(child)
	return true
}

// こどもID情報削除
func DeleteChild(name string, childId int) bool {
	var child model.UserChild
	err := db.Where("name = ? and child_id = ?", name, childId).First(&child).Error
	if err != nil {
		return false
	}
	db.Delete(&child)
	return true
}

func bGetDateUser(date time.Time) []model.User {
	var user []model.User
	//day := date.Format("2006-01-02")
	db.Where("created_at = ?", date).Find(&user)
	return user
}
