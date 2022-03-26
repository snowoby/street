package d

import "github.com/google/uuid"

type IDCount struct {
	ID    uuid.UUID `json:"id"`
	Count int       `json:"count"`
}
