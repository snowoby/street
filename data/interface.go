package data

import (
	"context"
	"github.com/google/uuid"
)

type Owner interface {
	IsOwner(ctx context.Context, ownerID uuid.UUID, objID uuid.UUID) (bool, error)
}
