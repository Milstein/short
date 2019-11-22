package repo

import (
	"errors"
	"short/app/entity"
)

var _ User = (*UserFake)(nil)

type UserFake struct {
	users map[string]entity.User
}

func (u UserFake) IsEmailExist(email string) (bool, error) {
	_, ok := u.users[email]
	return ok, nil
}

func (u UserFake) GetByEmail(email string) (entity.User, error) {
	user, ok := u.users[email]
	if !ok {
		return entity.User{}, errors.New("can't find email")
	}
	return user, nil
}

func (u UserFake) CreateUser(user entity.User) error {
	u.users[user.Email] = user
	return nil
}

func (u UserFake) UpdateUserID(email string, userID string) (entity.User, error) {
	user, ok := u.users[email]
	if !ok {
		return entity.User{}, errors.New("can't find email")
	}
	user.ID = userID
	return user, nil
}

func NewUserFake(users map[string]entity.User) UserFake {
	return UserFake{
		users: users,
	}
}
