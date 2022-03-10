/*
Package dao
@Author: MoZhu
@File: session
@Software: GoLand
*/
package dao

import "time"

type Session struct {
	ID        string    `sql:"unique_index"`
	UserName  string    // 用户名
	UserId    string    // 用户ID
	ClientIP  string    // 客户端IP
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	ExpiresAt time.Time // 过期时间
}

func (Session) TableName() string {
	return "sessions"
}

func NewSessionDAO(client *MySQLClient) SessionDAOIF {
	return &sessionDAO{
		client: client,
	}
}

type SessionDAOIF interface {
	Delete(where string, args []interface{}) error
	GetDB() Client
}

type sessionDAO struct {
	client Client
}

func (s *sessionDAO) Delete(where string, args []interface{}) error {
	tx := s.client.DB().Begin()
	if tx.Where(where, args...).Delete(&Session{}); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return tx.Commit().Error
}

func (s *sessionDAO) GetDB() Client {
	return s.client
}
