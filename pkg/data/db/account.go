package db

import (
	"context"
	"street/ent"
	ea "street/ent/account"
)

type account struct {
	client *ent.AccountClient
}

func (s *account) Find(ctx context.Context, email string) (*ent.Account, error) {
	user, err := s.client.Query().Where(ea.Email(email)).Only(ctx)
	return user, err
}

func (s *account) Create(ctx context.Context, email string, encryptedPassword string) (*ent.Account, error) {
	user, err := s.client.Create().SetEmail(email).SetPassword(encryptedPassword).Save(ctx)
	return user, err
}

func (s *account) EmailExists(ctx context.Context, email string) (bool, error) {
	return s.client.Query().Where(ea.Email(email)).Exist(ctx)
}
