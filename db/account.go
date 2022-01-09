package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	"street/ent/account"
	"street/ent/token"
	"time"
)

func (s *db) FindAccount(ctx context.Context, email string) (*ent.Account, error) {
	user, err := s.Account.Query().Where(account.Email(email)).Only(ctx)
	return user, err
}

// CreateAccount The encryptedPassword here will be stored directly!
func (s *db) CreateAccount(ctx context.Context, email string, encryptedPassword string) (*ent.Account, error) {
	user, err := s.Account.Create().SetEmail(email).SetPassword(encryptedPassword).Save(ctx)
	return user, err
}

func (s *db) EmailExists(ctx context.Context, email string) bool {
	user, _ := s.FindAccount(ctx, email)
	return user != nil
}

func (s *db) FindToken(ctx context.Context, tokenBody, t string, validOnly bool) (*ent.Token, error) {
	query := s.Token.Query().Where(token.Body(tokenBody)).Where(token.Type(t)).WithAccount()
	if validOnly {
		query = query.Where(token.ExpireTimeGT(time.Now()))
	}
	return s.Token.Query().Where(token.Body(tokenBody)).WithAccount().Only(ctx)
}

func (s *db) CreateToken(ctx context.Context, accountID uuid.UUID, tokenBody, tokenType string, lifelong time.Duration) (*ent.Token, error) {
	t, err := s.Token.Create().
		SetAccountID(accountID).
		SetBody(tokenBody).
		SetType(tokenType).
		SetExpireTime(time.Now().Add(lifelong)).
		Save(ctx)
	return t, err
}
