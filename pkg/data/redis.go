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

func (r *FileRedis) Get(ctx context.Context, fileID string) (string, error) {
	return r.client.Get(ctx, fileCreateKey(fileID)).Result()
}

func (r *FileRedis) Part(ctx context.Context, fileID string, partID int, partUploadObj interface{}) error {
	return r.client.ZAdd(ctx, filePartKey(fileID), &redis.Z{Score: float64(partID), Member: partUploadObj}).Err()
}

func (r *FileRedis) Complete(ctx context.Context, fileID string) error {
	parts, err := r.client.ZRange(ctx, filePartKey(fileID), 0, -1).Result()
	if err != nil {
		return err
	}

	fmt.Println(parts)
	start, err := r.client.Get(ctx, fileCreateKey(fileID)).Result()
	if err != nil {
		return err
	}
	fmt.Println(start)
	return err
}

func (r *FileRedis) Finish(ctx context.Context, fileID string) error {
	err := r.client.Del(ctx, fileCreateKey(fileID)).Err()
	if err != nil {
		fmt.Println(err)
	}
	err = r.client.Del(ctx, filePartKey(fileID)).Err()
	if err != nil {
		fmt.Println(err)
	}

	return err
}
