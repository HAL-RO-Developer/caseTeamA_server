package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// ユーザー基本情報
type User struct {
	Model
	Name     string `json:"name" binding:"required"`
	Password string `json:"pass" binding:"required"`
}

// ユーザー任意情報
type UserChild struct {
	Model
	Name     string    `json:"name" binding:"required"`
	ChildId  int       `json:"child_id"`
	BirthDay time.Time `json:"birthday" sql:"type:date"`
	NickName string    `json:"nickname"`
	Sex      int       `json:"sex"`
}

// デバイス
type Device struct {
	Model
	Name     string `json:"name"`
	ChildId  int    `json:"child_id"`
	DeviceId string `json:"device_id"`
}

// BOCCOAPI設定
type Bocco struct {
	Model
	Name  string `json:"name"`
	Key   string `json:"key"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
