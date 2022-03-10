/*
Package dao
@Author: MoZhu
@File: user
@Software: GoLand
*/
package dao

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName  string `json:"user_name"`
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	AvatarURL string `json:"avatar_url"`
	Source    int    `json:"source"`
}

func (u *User) TableName() string {
	return "users"
}
