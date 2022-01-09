package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	"street/ent/account"
	"street/ent/profile"
)

func (s *db) FindProfile(ctx context.Context, callSign string) (*ent.Profile, error) {
	p, err := s.Profile.Query().Where(profile.CallSign(callSign)).WithAccount().Only(ctx)
	return p, err
}
func (s *db) FindProfilesByAccountID(ctx context.Context, id uuid.UUID) ([]*ent.Profile, error) {
	ps, err := s.Profile.Query().Where(profile.HasAccountWith(account.ID(id))).All(ctx)
	return ps, err
}

func (s *db) FindProfileByID(ctx context.Context, id uuid.UUID) (*ent.Profile, error) {
	p, err := s.Profile.Query().Where(profile.ID(id)).WithAccount().Only(ctx)
	return p, err
}

func (s *db) CreateProfile(ctx context.Context, callSign, title, category string, accountID uuid.UUID) (*ent.Profile, error) {
	p, err := s.Profile.Create().SetTitle(title).SetCallSign(callSign).SetCategory(category).SetAccountID(accountID).Save(ctx)
	return p, err
}

func (s *db) CallSignExists(ctx context.Context, callSign string) (bool, error) {
	result, err := s.FindProfile(ctx, callSign)
	if ent.IsNotFound(err) {
		err = nil
	}
	return result != nil, err
}

func (s *db) UpdateProfile(ctx context.Context, profileID uuid.UUID, callSign, title, category string) (*ent.Profile, error) {
	p, err := s.Profile.UpdateOneID(profileID).SetTitle(title).SetCallSign(callSign).SetCategory(category).Save(ctx)
	return p, err
}
