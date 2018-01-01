package interfaces

import (
	"errors"

	"github.com/megaminx/white/cmd/business/user"
)

var (
	ErrNotFound = errors.New("not found")
)

type (
	InvitationRepo struct {
		DB DBHandler
	}

	DBHandler interface {
		Store(data interface{}) error
		FindOne(conditions map[string]interface{}, result interface{}) error
	}
)

func (r *InvitationRepo) Store(invitation *user.Invitation) error {
	if err := r.DB.Store(invitation); err != nil {
		return err
	}
	return nil
}

func (r InvitationRepo) Find(email string) (*user.Invitation, error) {
	c := make(map[string]interface{})
	c["email"] = email

	var result user.Invitation
	err := r.DB.FindOne(c, &result)
	if err == ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &result, nil
}
