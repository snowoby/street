package factory

import (
	"encoding/json"
	"fmt"
	"street/ent"

	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
)

const (
	ResizeWidth  = 1024
	ResizeHeight = 1024

	ThumbnailWidth  = 1024
	ThumbnailHeight = 1024

	Quality = 85
)

func (s *service) HandleImageCompressTask(_ context.Context, t *asynq.Task) error {
	var p ent.File
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	b, err := s.storageService.Get(p.Path, p.ID.String(), "original")

	if err != nil {
		return fmt.Errorf("get from s3 failed: %v", err)
	}

	return ImageCompress(&p, b, s.storageService)
}

func (s *service) HandleAvatarCompressTask(_ context.Context, t *asynq.Task) error {
	var p ent.File
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	b, err := s.storageService.Get(p.Path, p.ID.String(), "original")
	if err != nil {
		return fmt.Errorf("get from s3 failed: %v: %w", err, asynq.SkipRetry)
	}
	err = AvatarCompress(&p, b, s.storageService)
	if err != nil {
		return fmt.Errorf("compress failed: %v", err)
	}
	err = s.storageService.Delete(p.Path, p.ID, "original")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}
	return nil

}
