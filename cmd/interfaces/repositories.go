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
		DB          DBHandler
		StorageName string
	}

	DBHandler interface {
		Store(selector map[string]interface{}, data interface{}, to string) error
		FindOne(selector map[string]interface{}, result interface{}, from string) error
	}
)

func (r *InvitationRepo) Verify(invitation *user.Invitation) error {
	s := make(map[string]interface{})
	s["email"] = invitation.Email
	if err := r.DB.Store(s, invitation, r.StorageName); err != nil {
		return err
	}
	return nil
}

func (r *InvitationRepo) Store(invitation *user.Invitation) error {
	if err := r.DB.Store(nil, invitation, r.StorageName); err != nil {
		return err
	}
	return nil
}

func (r InvitationRepo) Find(email string) (*user.Invitation, error) {
	s := make(map[string]interface{})
	s["email"] = email

	var result user.Invitation
	err := r.DB.FindOne(s, &result, r.StorageName)
	if err == ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &result, nil
}
