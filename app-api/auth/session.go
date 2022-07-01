package auth

import "github.com/supinf/supinf-mail/app-api/model"

type Session struct {
	APIKey   string
	UserName string
	UserMail string
	Role     model.Role
}

// IsAdmin 管理者かどうかを返します
func (s Session) IsAdmin() bool {
	switch s.Role {
	case model.RoleAdmin:
		return true
	default:
		return false
	}
}
