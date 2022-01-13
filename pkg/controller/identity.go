package controller

import "street/ent"

type Identity struct {
	account *ent.Account
	profile *ent.Profile
}

func (i *Identity) Profile() *ent.Profile {
	return i.profile
}

func (i *Identity) Account() *ent.Account {
	return i.account
}

func (i *Identity) HasAccount() bool {
	return i.account != nil
}

func (i *Identity) HasProfile() bool {
	return i.profile != nil
}
