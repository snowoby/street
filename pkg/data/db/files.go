package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	ea "street/ent/account"
	ef "street/ent/file"
)

type file struct {
	client *ent.FileClient
}

func (f *file) Create(ctx context.Context, filename, path string, mime string, size int, accountID uuid.UUID) (*ent.File, error) {
	return f.client.Create().SetFilename(filename).SetMime(mime).SetSize(size).SetAccountID(accountID).SetPath(path).Save(ctx)
}
func (f *file) CreateFinished(ctx context.Context, filename, path string, mime string, size int, accountID uuid.UUID) (*ent.File, error) {
	return f.client.Create().SetFilename(filename).SetMime(mime).SetSize(size).SetAccountID(accountID).SetPath(path).SetStatus("uploaded").Save(ctx)
}

func (f *file) Get(ctx context.Context, fileID uuid.UUID) (*ent.File, error) {
	return f.client.Get(ctx, fileID)
}

func (f *file) GetWithAccount(ctx context.Context, fileID uuid.UUID) (*ent.File, error) {
	return f.client.Query().Where(ef.ID(fileID)).WithAccount().First(ctx)
}

func (f *file) NotUploadedExists(ctx context.Context, fileID uuid.UUID, accountID uuid.UUID) (bool, error) {
	return f.client.Query().Where(ef.ID(fileID), ef.HasAccountWith(ea.ID(accountID)), ef.Status("created")).Exist(ctx)
}

func (f *file) UploadedExists(ctx context.Context, fileID uuid.UUID, accountID uuid.UUID) (bool, error) {
	return f.client.Query().Where(ef.ID(fileID), ef.HasAccountWith(ea.ID(accountID)), ef.Status("uploaded")).Exist(ctx)
}

func (f *file) UpdateStatus(ctx context.Context, fileID uuid.UUID, status string) (*ent.File, error) {
	return f.client.UpdateOneID(fileID).SetStatus(status).Save(ctx)
}

func (f *file) Delete(ctx context.Context, fileID uuid.UUID) error {
	return f.client.DeleteOneID(fileID).Exec(ctx)
}

func (f *file) IsOwner(ctx context.Context, accountID uuid.UUID, fileID uuid.UUID) (bool, error) {
	return f.client.Query().Where(ef.HasAccountWith(ea.ID(accountID)), ef.ID(fileID)).Exist(ctx)
}
