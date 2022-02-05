package data

import (
	"context"
	"github.com/google/uuid"
)

type Owner interface {
	OwnerID(ctx context.Context, objID uuid.UUID) (string, error)
}
