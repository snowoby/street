package data

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	"street/ent/account"
	ep "street/ent/profile"
)

type profile struct {
	client *ent.ProfileClient
}

func (profile *profile) Client() *ent.ProfileClient {
	return profile.client
}

func (profile *profile) FindProfile(ctx context.Context, callSign string) (*ent.Profile, error) {
	p, err := profile.client.Query().Where(ep.CallSign(callSign)).WithAccount().Only(ctx)
	return p, err
}
func (profile *profile) FindProfilesByAccountID(ctx context.Context, id uuid.UUID) ([]*ent.Profile, error) {
	ps, err := profile.client.Query().Where(ep.HasAccountWith(account.ID(id))).All(ctx)
	return ps, err
}

func (profile *profile) FindProfileByID(ctx context.Context, id uuid.UUID) (*ent.Profile, error) {
	p, err := profile.client.Query().Where(ep.ID(id)).WithAccount().Only(ctx)
	return p, err
}

func (profile *profile) CreateProfile(ctx context.Context, callSign, title, category string, accountID uuid.UUID) (*ent.Profile, error) {
	p, err := profile.client.Create().SetTitle(title).SetCallSign(callSign).SetCategory(category).SetAccountID(accountID).Save(ctx)
	return p, err
}

func (profile *profile) CallSignExists(ctx context.Context, callSign string) (bool, error) {
	return profile.client.Query().Where(ep.CallSign(callSign)).Exist(ctx)
}

func (profile *profile) UpdateProfile(ctx context.Context, profileID uuid.UUID, callSign, title, category string) (*ent.Profile, error) {
	p, err := profile.client.UpdateOneID(profileID).SetTitle(title).SetCallSign(callSign).SetCategory(category).Save(ctx)
	return p, err
}

func (profile *profile) IsOwner(ctx context.Context, ownerID uuid.UUID, objectID uuid.UUID) (bool, error) {
	//ok, err := profile.client.Query().Where(ep.And(ep.HasAccountWith(account.ID(ownerID)), ep.ID(objectID))).Exist(ctx)
	return ownerID == objectID, nil
}
