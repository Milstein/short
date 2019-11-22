package account

var _ Account = (*Fake)(nil)

type Fake struct {
	isAccountExist    bool
	isAccountExistErr error
	createAccountErr  error
}

func (a Fake) IsAccountExist(email string) (bool, error) {
	return a.isAccountExist, a.isAccountExistErr
}

func (a Fake) CreateAccount(email string, name string) error {
	return a.createAccountErr
}

func NewFake(
	isAccountExist bool,
	isAccountExistErr error,
	createAccountErr error,
) Fake {
	return Fake{
		isAccountExist:    isAccountExist,
		isAccountExistErr: isAccountExistErr,
		createAccountErr:  createAccountErr,
	}
}
