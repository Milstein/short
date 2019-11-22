package repo

import (
	"errors"
	"short/app/entity"
)

var _ AccountsRelation = (*AccountsRelationFake)(nil)

type Relation struct {
	ssoUserID string
	userID    string
}

type AccountsRelationFake struct {
	relations []Relation
}

func (a AccountsRelationFake) IsSSOUserExist(ssoUser entity.SSOUser) (bool, error) {
	for _, relation := range a.relations {
		if relation.ssoUserID == ssoUser.ID {
			return true, nil
		}
	}
	return false, nil
}

func (a AccountsRelationFake) IsRelationExist(ssoUser entity.SSOUser, user entity.User) bool {
	for _, relation := range a.relations {
		if relation.ssoUserID == ssoUser.ID && relation.userID == user.ID {
			return true
		}
	}
	return false
}

func (a *AccountsRelationFake) CreateRelation(ssoUser entity.SSOUser, user entity.User) error {
	isExist := a.IsRelationExist(ssoUser, user)
	if isExist {
		return errors.New("relation exists")
	}
	a.relations = append(a.relations, Relation{
		ssoUserID: ssoUser.ID,
		userID:    user.ID,
	})
	return nil
}

func NewAccountsRelationFake(relations []Relation) AccountsRelationFake {
	return AccountsRelationFake{
		relations: relations,
	}
}
