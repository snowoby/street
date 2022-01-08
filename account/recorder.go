package account

import (
	"context"
	"github.com/google/uuid"
	"street/db"
	"street/ent"
	"street/ent/account"
	"street/ent/token"
	"time"
)

type store struct {
	db.Store
}

const (
	RefreshToken = "refresh"
	AccessToken  = "access"
)

func (s *store) findAccount(ctx context.Context, email string) (*ent.Account, error) {
	user, err := s.DB().Account.Query().Where(account.Email(email)).Only(ctx)
	return user, err
}

// The password here will be stored directly!
func (s *store) createAccount(ctx context.Context, email string, encryptedPassword string) (*ent.Account, error) {
	user, err := s.DB().Account.Create().SetEmail(email).SetPassword(encryptedPassword).Save(ctx)
	return user, err
}

func (s *store) emailExists(ctx context.Context, email string) bool {
	user, _ := s.findAccount(ctx, email)
	return user != nil
}

func (s *store) findToken(ctx context.Context, tokenBody, t string, validOnly bool) (*ent.Token, error) {
	query := s.DB().Token.Query().Where(token.Body(tokenBody)).Where(token.Type(t)).WithAccount()
	if validOnly {
		query = query.Where(token.ExpireAtGT(time.Now()))
	}
	return s.DB().Token.Query().Where(token.Body(tokenBody)).WithAccount().Only(ctx)
}

func (s *store) createToken(ctx context.Context, accountID uuid.UUID, tokenBody, tokenType string, lifelong time.Duration) (*ent.Token, error) {
	t, err := s.DB().Token.Create().
		SetAccountID(accountID).
		SetBody(tokenBody).
		SetType(tokenType).
		SetExpireAt(time.Now().Add(lifelong)).
		Save(ctx)
	return t, err
}
