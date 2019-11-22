package account

import (
	"short/app/entity"
	"short/app/usecase/repo"

	"github.com/byliuyang/app/fw"
)

var _ Account = (*Persist)(nil)

type Persist struct {
	userRepo repo.User
	timer    fw.Timer
}

func (r Persist) IsAccountExist(email string) (bool, error) {
	return r.userRepo.IsEmailExist(email)
}

func (r Persist) CreateAccount(email string, name string) error {
	now := r.timer.Now()
	user := entity.User{
		Email:     email,
		Name:      name,
		CreatedAt: &now,
	}
	return r.userRepo.CreateUser(user)
}

func NewPersist(userRepo repo.User, timer fw.Timer) Persist {
	return Persist{
		userRepo: userRepo,
		timer:    timer,
	}
}
