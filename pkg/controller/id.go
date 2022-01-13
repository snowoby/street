package controller

import "github.com/google/uuid"

type ID struct {
	ID uuid.UUID `uri:"id" binding:"uuid"`
}
