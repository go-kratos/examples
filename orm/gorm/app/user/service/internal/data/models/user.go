package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"column:username"`
	NickName string `gorm:"column:nickname"`
	Password string `gorm:"column:password"`
}

func (u User) TableName() string {
	return "users"
}
