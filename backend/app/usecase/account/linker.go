package account

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
)

type Linker struct {
	keyGen               keygen.KeyGenerator
	userRepo             repo.User
	accountsRelationRepo repo.AccountsRelation
}

func (l Linker) IsAccountLinked(ssoUser entity.SSOUser) (bool, error) {
	return l.accountsRelationRepo.IsSSOUserExist(ssoUser)
}

func (l Linker) LinkAccount(ssoUser entity.SSOUser) error {
	user, err := l.ensureUserExist(ssoUser)
	if err != nil {
		return err
	}
	return l.accountsRelationRepo.CreateRelation(ssoUser, user)
}

func (l Linker) ensureUserExist(ssoUser entity.SSOUser) (entity.User, error) {
	isEmailExist, err := l.userRepo.IsEmailExist(ssoUser.Email)
	if err != nil {
		return entity.User{}, err
	}
	userID, err := l.generateUnassignedUserID()
	if err != nil {
		return entity.User{}, err
	}

	if isEmailExist {
		return l.assignUserID(ssoUser.Email, userID)
	}
	return l.createUser(userID, ssoUser.Name, ssoUser.Email)
}

func (l Linker) generateUnassignedUserID() (string, error) {
	newKey, err := l.keyGen.NewKey()
	return string(newKey), err
}

func (l Linker) createUser(id string, name string, email string) (entity.User, error) {
	user := entity.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	err := l.userRepo.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (l Linker) assignUserID(userEmail string, userID string) (entity.User, error) {
	return l.userRepo.UpdateUserID(userEmail, userID)
}

func NewLinker(
	keyGen keygen.KeyGenerator,
	userRepo repo.User,
	accountRelationRepo repo.AccountsRelation,
) Linker {
	return Linker{
		keyGen:               keyGen,
		userRepo:             userRepo,
		accountsRelationRepo: accountRelationRepo,
	}
}
