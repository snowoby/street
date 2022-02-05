package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	ea "street/ent/account"
	ep "street/ent/profile"
)

type profile struct {
	client *ent.ProfileClient
}

func (profile *profile) Client() *ent.ProfileClient {
	return profile.client
}

func (profile *profile) Find(ctx context.Context, call string) (*ent.Profile, error) {
	p, err := profile.client.Query().Where(ep.Call(call)).WithAccount().Only(ctx)
	return p, err
}
func (profile *profile) FindByAccountID(ctx context.Context, id uuid.UUID) ([]*ent.Profile, error) {
	ps, err := profile.client.Query().Where(ep.HasAccountWith(ea.ID(id))).All(ctx)
	return ps, err
}

func (profile *profile) FindByID(ctx context.Context, id uuid.UUID) (*ent.Profile, error) {
	p, err := profile.client.Query().Where(ep.ID(id)).WithAccount().Only(ctx)
	return p, err
}

func (profile *profile) Create(ctx context.Context, callSign, title, category, avatar string, accountID uuid.UUID) (*ent.Profile, error) {
	p, err := profile.client.Create().SetTitle(title).SetCall(callSign).SetCategory(category).SetAccountID(accountID).SetAvatar(avatar).Save(ctx)
	return p, err
}

func (profile *profile) CallExists(ctx context.Context, call string) (bool, error) {
	return profile.client.Query().Where(ep.Call(call)).Exist(ctx)
}

func (profile *profile) Update(ctx context.Context, profileID uuid.UUID, title, call, category, avatar string) (*ent.Profile, error) {
	p, err := profile.client.UpdateOneID(profileID).SetTitle(title).SetCall(call).SetCategory(category).SetAvatar(avatar).Save(ctx)
	return p, err
}

func (profile *profile) IsOwner(ctx context.Context, ownerID uuid.UUID, objectID uuid.UUID) (bool, error) {
	ok, err := profile.client.Query().Where(ep.And(ep.ID(objectID), ep.HasAccountWith(ea.ID(ownerID)))).Exist(ctx)
	return ok, err
	//return ownerID == objectID, nil
}
