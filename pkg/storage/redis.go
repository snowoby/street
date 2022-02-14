package storage

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

type redisService struct {
	client *redis.Client
}

func newRedis(client *redis.Client) *redisService {
	return &redisService{client: client}
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

func (r *redisService) Create(ctx context.Context, fileID string, fileCreateObj interface{}) error {
	data, err := json.Marshal(fileCreateObj)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, fileCreateKey(fileID), data, FilePartAge).Err()
}

func (r *redisService) Get(ctx context.Context, fileID string) (string, error) {
	return r.client.Get(ctx, fileCreateKey(fileID)).Result()
}

func (r *redisService) Part(ctx context.Context, fileID string, partUploadObj interface{}) error {
	data, err := json.Marshal(partUploadObj)
	if err != nil {
		return err
	}
	r.client.Expire(ctx, filePartKey(fileID), FilePartAge)
	return r.client.RPush(ctx, filePartKey(fileID), data).Err()
}

func (r *redisService) GetParts(ctx context.Context, fileID string) ([]string, error) {
	parts, err := r.client.LRange(ctx, filePartKey(fileID), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return parts, err
}

func (r *redisService) Finish(ctx context.Context, fileID string) error {
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
