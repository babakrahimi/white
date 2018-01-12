package user

import "errors"

var (
	ErrInvalidInvitationToken = errors.New("invalid invitation token")
	ErrExpiredInvitationToken = errors.New("expired invitation token")
	ErrNotInvitedUser         = errors.New("not invited user")
)

type (
	InvitationRepository interface {
		Store(invitation *Invitation) error
		Verify(invitation *Invitation) error
		Find(email string) (*Invitation, error)
	}

	CryptoHandler interface {
		GetInvitationToken(email string) (string, error)
		VerifyInvitationToken(token string) (string, error)
	}

	EmailHandler interface {
		SendInvitationEmail(to, token string) error
	}

	Invitation struct {
		Email   string `json:"email" bson:"email"`
		Visited bool   `json:"visited" bson:"visited"`
	}

	InvitationOperator interface {
		InviteUser(email string) error
		VerifyInvitation(token string) (string, error)
	}

	InvitationAgent struct {
		Repository    InvitationRepository
		CryptoHandler CryptoHandler
		EmailHandler  EmailHandler
	}
)

func (a *InvitationAgent) VerifyInvitation(token string) (string, error) {
	email, err := a.CryptoHandler.VerifyInvitationToken(token)
	if err != nil {
		return "", err
	}

	inv, err := a.Repository.Find(email)
	if err != nil {
		return "", err
	}
	if inv == nil {
		return "", ErrNotInvitedUser
	}

	inv.Visited = true
	if err := a.Repository.Verify(inv); err != nil {
		return "", err
	}

	return email, nil
}

func (a *InvitationAgent) InviteUser(email string) error {
	inv, err := a.Repository.Find(email)
	if inv != nil {
		return errors.New("duplicate invitation error")
	}
	if err != nil {
		return err
	}

	invitation := &Invitation{Email: email}
	if err := a.Repository.Store(invitation); err != nil {
		return err
	}

	t, err := a.CryptoHandler.GetInvitationToken(email)
	if err != nil {
		return err
	}

	if err := a.EmailHandler.SendInvitationEmail(email, t); err != nil {
		return err
	}

	return nil
}
