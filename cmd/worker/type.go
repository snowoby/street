package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"golang.org/x/net/context"
	"gopkg.in/gographics/imagick.v3/imagick"
	"street/ent"
	"street/pkg/data"
	"time"
)

func HandleImageCompressTask(_ context.Context, t *asynq.Task, store *data.Store) error {
	var p ent.File
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	b, err := store.Storage.Get(p.Path, p.ID.String(), "original")
	if err != nil {
		return fmt.Errorf("get from s3 failed: %v: %w", err, asynq.SkipRetry)
	}

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err = mw.ReadImageBlob(b)
	if err != nil {
		return fmt.Errorf("read imagick failed: %v: %w", err, asynq.SkipRetry)
	}

	err = mw.SetImageCompressionQuality(90)
	if err != nil {
		return fmt.Errorf("set compress quality failed: %v: %w", err, asynq.SkipRetry)
	}

	err = mw.SetFormat("webp")
	if err != nil {
		return fmt.Errorf("set format failed: %v: %w", err, asynq.SkipRetry)
	}

	nb := mw.GetImageBlob()
	newFilename := p.Filename + ".webp"
	_, err = store.Storage.PutSingle(bytes.NewReader(nb), p.Path, p.ID, newFilename, "compressed", "image/webp")
	if err != nil {
		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
	}
	time.Sleep(time.Second * 100)
	return nil

}
