package data

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	et "street/ent/token"

	"street/ent"
	"time"
)

type token struct {
	client *ent.TokenClient
}

func (s *token) Find(ctx context.Context, tokenBody, t string, validOnly bool) (*ent.Token, error) {
	query := s.client.Query().Where(et.Body(tokenBody)).Where(et.Type(t))
	if validOnly {
		query = query.Where(et.ExpireGT(time.Now()))
	}
	return query.WithAccount().Only(ctx)
}

func (s *token) Create(ctx context.Context, accountID uuid.UUID, tokenBody, tokenType string, lifelong time.Duration) (*ent.Token, error) {
	t, err := s.client.Create().
		SetAccountID(accountID).
		SetBody(tokenBody).
		SetType(tokenType).
		SetExpire(time.Now().Add(lifelong)).
		Save(ctx)
	return t, err
}
