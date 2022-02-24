package factory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
	"gopkg.in/gographics/imagick.v3/imagick"
	"street/ent"
)

func imageWebpCompress(b []byte, quality uint) ([]byte, error) {

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImageBlob(b)
	if err != nil {
		return nil, fmt.Errorf("read imagick failed: %v: %w", err, asynq.SkipRetry)
	}

	err = mw.SetImageCompressionQuality(quality)
	if err != nil {
		return nil, fmt.Errorf("set compress quality failed: %v: %w", err, asynq.SkipRetry)
	}

	err = mw.SetFormat("webp")
	if err != nil {
		return nil, fmt.Errorf("set format failed: %v: %w", err, asynq.SkipRetry)
	}

	nb := mw.GetImageBlob()
	return nb, nil

}

func imageCrop(b []byte, max uint) ([]byte, error) {

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImageBlob(b)
	if err != nil {
		return nil, fmt.Errorf("read imagick failed: %v: %w", err, asynq.SkipRetry)
	}
	height := mw.GetImageHeight()
	size := height
	width := mw.GetImageWidth()
	if width < size {
		size = width
	}
	if size > max {
		size = max
	}

	err = mw.ResizeImage(size, size, imagick.FILTER_UNDEFINED)
	if err != nil {
		return nil, fmt.Errorf("crop failed: %v: %w", err, asynq.SkipRetry)
	}

	nb := mw.GetImageBlob()
	return nb, nil

}

func (s *service) HandleImageCompressTask(_ context.Context, t *asynq.Task) error {
	var p ent.File
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	b, err := s.storageService.Get(p.Path, p.ID.String(), "original")
	if err != nil {
		return fmt.Errorf("get from s3 failed: %v: %w", err, asynq.SkipRetry)
	}
	nb, err := imageWebpCompress(b, 80)
	if err != nil {
		return fmt.Errorf("compress failed: %v: %w", err, asynq.SkipRetry)
	}
	newFilename := p.Filename + ".webp"
	_, err = s.storageService.PutSingle(bytes.NewReader(nb), p.Path, p.ID, newFilename, "compressed", "image/webp")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}
	return nil

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

	b, err = imageWebpCompress(b, 80)
	if err != nil {
		return err
	}

	nb, err := imageCrop(b, 512)
	if err != nil {
		return err
	}

	newFilename := p.Filename + ".webp"
	_, err = s.storageService.PutSingle(bytes.NewReader(nb), p.Path, p.ID, newFilename, "compressed", "image/webp")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}

	err = s.storageService.Delete(p.Path, p.ID, "original")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}
	return nil

}
