package d

import (
	"github.com/google/uuid"
	"street/ent"
)

type ID struct {
	ID uuid.UUID `json:"id" binding:"required,uuid"`
}

type AccountForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Account struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AccountFromEnt(data *ent.Account) *Account {
	return &Account{
		ID:       data.ID.String(),
		Email:    data.Email,
		Password: data.Password,
	}
}
