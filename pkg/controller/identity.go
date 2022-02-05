package controller

import (
	"github.com/google/uuid"
	"street/ent"
	"street/errs"
)

type Identity struct {
	account     *ent.Account
	allProfiles []*ent.Profile
}

//func (i *Identity) Profile() *ent.Profile {
//	return i.profile
//}

func (i *Identity) AllProfiles() []*ent.Profile {
	return i.allProfiles
}

func (i *Identity) HaveProfile(id uuid.UUID) bool {
	for _, profile := range i.allProfiles {
		if profile.ID == id {
			return true
		}
	}
	return false
}

func (i *Identity) HaveProfileX(id uuid.UUID) error {
	if !i.HaveProfile(id) {
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
