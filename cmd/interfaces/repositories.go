package interfaces

import (
	"github.com/megaminx/white/cmd/business"
	"github.com/megaminx/white/cmd/core"
)



type UserRepo struct {
}

type EmployeeRepo struct {
}

func (er *EmployeeRepo) Store(employee *core.Employee) error {
	return nil
}

func (ur *UserRepo) Store(user *business.User) error {

	return nil
}

func (ur *UserRepo) FindUserByEmail(email string) (*business.User, error) {
	return nil, nil
}

func (ur *UserRepo) SetUserStateToVerified(email string) error {
	return nil
}
