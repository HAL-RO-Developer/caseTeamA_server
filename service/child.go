package service

import (
	"time"

	"github.com/HAL-RO-Developer/caseTeamA_server/controller/validation"
	"github.com/HAL-RO-Developer/caseTeamA_server/model"
)

// 子ども情報取得
func FindByUserName(name string) ([]model.UserChild, bool) {
	var children []model.UserChild
	db.Where("name = ?", name).Find(&children)
	return children, len(children) != 0
}

// 子どもID検索
func GetByChildInfo(name string, childId int) ([]model.UserChild, bool) {
	var children []model.UserChild
	db.Where("name = ? and child_id = ?", name, childId).Find(&children)
	return children, len(children) != 0
}

// 子供情報追加
func AddChild(name string, info validation.UserChildren) (int, bool) {
	var i int
	find := true

	for i = 1; find == true; i++ {
		_, find = GetByChildInfo(name, i)
	}

	childId := i - 1

	birthday, err := time.Parse("2006-01-02", info.BirthDay)
	if err != nil {
		return 0, false
	}

	child := model.UserChild{
		Name:     name,
		ChildId:  childId,
		BirthDay: birthday,
		NickName: info.NickName,
		Sex:      info.Sex,
	}

	err = db.Create(&child).Error
	if err != nil {
		return 0, false
	}
	return childId, true
}
