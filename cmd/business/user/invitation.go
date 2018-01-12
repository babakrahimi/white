package user

import "errors"

type (
	InvitationRepository interface {
		Store(invitation *Invitation) error
		Find(email string) (*Invitation, error)
	}

	CryptoHandler interface {
		GetInvitationToken(email string) (string, error)
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
	}

	InvitationAgent struct {
		Repository    InvitationRepository
		CryptoHandler CryptoHandler
		EmailHandler  EmailHandler
	}
)

func (ia InvitationAgent) InviteUser(email string) error {
	inv, err := ia.Repository.Find(email)
	if inv != nil {
		return errors.New("duplicate invitation error")
	}
	if err != nil {
		return err
	}

	invitation := &Invitation{Email: email}
	if err := ia.Repository.Store(invitation); err != nil {
		return err
	}

	t, err := ia.CryptoHandler.GetInvitationToken(email)
	if err != nil {
		return err
	}

	if err := ia.EmailHandler.SendInvitationEmail(email, t); err != nil {
		return err
	}

	return nil
}
