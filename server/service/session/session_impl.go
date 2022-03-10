/*
Package session
@Author: MoZhu
@File: session_impl
@Software: GoLand
*/
package session

func (s *serviceImpl) CloseSessionBySessionKey(sessionKey string) error {
	return s.sessionDAO.Delete("id = ?", []interface{}{sessionKey})
}
