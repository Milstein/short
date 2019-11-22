package repo

import "short/app/entity"

type AccountsRelation interface {
	IsSSOUserExist(ssoUser entity.SSOUser) (bool, error)
	CreateRelation(ssoUser entity.SSOUser, user entity.User) error
}
