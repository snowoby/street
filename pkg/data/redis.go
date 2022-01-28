package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type FileRedis struct {
	client *redis.Client
}

const StringCreate = "create"
const StringPart = "part"
const FilePartAge = time.Minute * 20

/*
	FileID:create:{create_obj}
	FileID:part:id:{upload_obj}
*/

func fileCreateKey(fileID string) string {
	return fmt.Sprintf("%s:%s", fileID, StringCreate)
}

func filePartKey(fileID string) string {
	return fmt.Sprintf("%s:%s", fileID, StringPart)
}

func (r *FileRedis) Create(ctx context.Context, fileID string, fileCreateObj interface{}) error {
	return r.client.Set(ctx, fileCreateKey(fileID), fileCreateObj, FilePartAge).Err()
}

func (r *FileRedis) Exists(ctx context.Context, fileID string) error {
	return r.client.Exists(ctx, fileCreateKey(fileID)).Err()
}

func (r *FileRedis) Part(ctx context.Context, fileID string, partID int, partUploadObj interface{}) error {
	return r.client.ZAdd(ctx, filePartKey(fileID), &redis.Z{Score: float64(partID), Member: partUploadObj}).Err()
}

func (r *FileRedis) Complete(ctx context.Context, fileID string, partID int, partUploadObj interface{}) error {
	return r.client.ZAdd(ctx, filePartKey(fileID), &redis.Z{Score: float64(partID), Member: partUploadObj}).Err()
}
