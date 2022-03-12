package factory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"street/ent"
	"street/pkg/d"
	"street/pkg/utils"

	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	ResizeWidth  = 1024
	ResizeHeight = 1024

	ThumbnailWidth  = 1024
	ThumbnailHeight = 1024

	Quality = 80
)

func (s *service) HandleImageCompressTask(_ context.Context, t *asynq.Task) error {
	var p ent.File
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	b, err := s.storageService.Get(p.Path, p.ID.String(), "original")
	if err != nil {
		return fmt.Errorf("get from s3 failed: %v: %w", err, asynq.SkipRetry)
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err = mw.ReadImageBlob(b)
	if err != nil {
		return fmt.Errorf("read blob failed: %v: %w", err, asynq.SkipRetry)
	}

	// thumbnail
	{
		mw := mw.Clone()
		defer mw.Destroy()
		if mw.GetImageWidth() > ResizeWidth || mw.GetImageHeight() > ResizeHeight {
			width, height := utils.ResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), ResizeWidth)
			err := mw.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
			if err != nil {
				return fmt.Errorf("resize failed: %v: %w", err, asynq.SkipRetry)
			}
		}
		err := mw.SetImageCompressionQuality(Quality)
		if err != nil {
			return fmt.Errorf("set compress quality failed: %v: %w", err, asynq.SkipRetry)
		}

		err = mw.SetFormat("webp")
		if err != nil {
			return fmt.Errorf("set format failed: %v: %w", err, asynq.SkipRetry)
		}

		filename := fmt.Sprintf("%s_%s.%s", p.Filename, "thumbnail", "webp")
		_, err = s.storageService.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, filename, "thumbnail", "image/webp")
		if err != nil {
			return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
		}
	}

	// compressed
	{
		mw := mw.Clone()
		defer mw.Destroy()
		if mw.GetImageWidth() > ResizeWidth || mw.GetImageHeight() > ResizeHeight {
			width, height := utils.ResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), ResizeWidth)
			err := mw.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
			if err != nil {
				return fmt.Errorf("resize failed: %v: %w", err, asynq.SkipRetry)
			}
		}
		err := mw.SetImageCompressionQuality(Quality)
		if err != nil {
			return fmt.Errorf("set compress quality failed: %v: %w", err, asynq.SkipRetry)
		}

		err = mw.SetFormat("webp")
		if err != nil {
			return fmt.Errorf("set format failed: %v: %w", err, asynq.SkipRetry)
		}

		filename := fmt.Sprintf("%s_%s.%s", p.Filename, "thumbnail", "webp")
		_, err = s.storageService.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, filename, d.StringThumbnail, "image/webp")
		if err != nil {
			return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
		}
	}

	// original size compress
	{
		mw := mw.Clone()
		defer mw.Destroy()

		err := mw.SetImageCompressionQuality(Quality)
		if err != nil {
			return fmt.Errorf("set compress quality failed: %v: %w", err, asynq.SkipRetry)
		}

		err = mw.SetFormat("webp")
		if err != nil {
			return fmt.Errorf("set format failed: %v: %w", err, asynq.SkipRetry)
		}

		filename := fmt.Sprintf("%s_%s.%s", p.Filename, "compressed", "webp")
		_, err = s.storageService.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, filename, d.StringCompressed, "image/webp")
		if err != nil {
			return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
		}
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

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err = mw.ReadImageBlob(b)

	if err != nil {
		return fmt.Errorf("read image blob failed: %v: %w", err, asynq.SkipRetry)
	}

	err = mw.SetImageCompressionQuality(Quality)
	if err != nil {
		return fmt.Errorf("set compress quality failed: %v: %w", err, asynq.SkipRetry)
	}

	if mw.GetImageWidth() > ResizeWidth || mw.GetImageHeight() > ResizeHeight {

		width, height := utils.ResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), ResizeWidth)

		err := mw.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
		if err != nil {
			return fmt.Errorf("resize failed: %v: %w", err, asynq.SkipRetry)
		}
	}
	if err != nil {
		return fmt.Errorf("crop failed: %v: %w", err, asynq.SkipRetry)
	}

	newFilename := p.Filename + ".webp"
	_, err = s.storageService.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, newFilename, "compressed", "image/webp")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}

	err = s.storageService.Delete(p.Path, p.ID, "original")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}
	return nil

}
