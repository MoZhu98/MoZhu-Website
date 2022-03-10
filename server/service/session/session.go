/*
Package session
@Author: MoZhu
@File: session
@Software: GoLand
*/
package session

import "github.com/mozhu98/website/server/dao"

type Service interface {
	CloseSessionBySessionKey(sessionKey string) error
}

type serviceImpl struct {
	sessionDAO dao.SessionDAOIF
}

func NewService(sessionDAO dao.SessionDAOIF) Service {
	return &serviceImpl{
		sessionDAO: sessionDAO,
	}
}
