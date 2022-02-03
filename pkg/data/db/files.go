package db

import (
	"context"
	"github.com/google/uuid"
	"street/ent"
	ef "street/ent/file"
	ep "street/ent/profile"
)

type file struct {
	client *ent.FileClient
}

func (f *file) Create(ctx context.Context, filename, path string, mime string, size int, profileID uuid.UUID) (*ent.File, error) {
	return f.client.Create().SetFilename(filename).SetMime(mime).SetSize(size).SetProfileID(profileID).SetPath(path).Save(ctx)
}
func (f *file) CreateFinished(ctx context.Context, filename, path string, mime string, size int, profileID uuid.UUID) (*ent.File, error) {
	return f.client.Create().SetFilename(filename).SetMime(mime).SetSize(size).SetProfileID(profileID).SetPath(path).SetStatus("uploaded").Save(ctx)
}

func (f *file) Get(ctx context.Context, fileID uuid.UUID) (*ent.File, error) {
	return f.client.Query().Where(ef.ID(fileID)).WithProfile().First(ctx)
}

func (f *file) NotUploadedExists(ctx context.Context, fileID uuid.UUID, profileID uuid.UUID) (bool, error) {
	return f.client.Query().Where(ef.ID(fileID), ef.HasProfileWith(ep.ID(profileID)), ef.Status("created")).Exist(ctx)
}

func (f *file) UploadedExists(ctx context.Context, fileID uuid.UUID, profileID uuid.UUID) (bool, error) {
	return f.client.Query().Where(ef.ID(fileID), ef.HasProfileWith(ep.ID(profileID)), ef.Status("uploaded")).Exist(ctx)
}

func (f *file) UpdateStatus(ctx context.Context, fileID uuid.UUID, status string) (*ent.File, error) {
	return f.client.UpdateOneID(fileID).SetStatus(status).Save(ctx)
}

func (f *file) Delete(ctx context.Context, fileID uuid.UUID) error {
	return f.client.DeleteOneID(fileID).Exec(ctx)
}

func (f *file) IsOwner(ctx context.Context, profileID uuid.UUID, fileID uuid.UUID) (bool, error) {
	return f.client.Query().Where(ef.HasProfileWith(ep.ID(profileID)), ef.ID(fileID)).Exist(ctx)
}
