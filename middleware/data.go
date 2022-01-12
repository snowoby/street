package middleware

import "github.com/google/uuid"

type ID struct {
	ID uuid.UUID `binding:"uuid" header:"Profile" uri:"id"`
}
