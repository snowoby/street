package operator

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"street/ent"
	"street/errs"
)

type Identity struct {
	account *ent.Account
}

//func (i *Identity) Profile() *ent.Profile {
//	return i.profile
//}

func New(account *ent.Account) *Identity {
	return &Identity{
		account: account,
	}
}

func (i *Identity) AllProfiles(ctx context.Context) []*ent.Profile {
	profiles, err := i.account.QueryProfile().All(ctx)
	if err != nil {
		fmt.Println(err)
	}
	return profiles
}

func (i *Identity) HaveProfile(ctx context.Context, id string) bool {
	for _, profile := range i.AllProfiles(ctx) {
		if profile.ID == uuid.MustParse(id) {
			return true
		}
	}
	return false
}

func (i *Identity) HaveProfileX(ctx context.Context, id string) error {
	if !i.HaveProfile(ctx, id) {
		return errs.UnauthorizedError
	}
	return nil
}

func (i *Identity) Account() *ent.Account {
	return i.account
}

//func (i *Identity) HasAccount() bool {
//	return i.account != nil
//}

//func (i *Identity) HasProfile() bool {
//	return i.profile != nil
//}
