package factory

type profile struct {
	maxWidth  int
	maxHeight int
	name      string
	quality   int
	keepRatio bool
	format    string
	mime      string
}

var AvatarProfiles = []profile{
	{maxWidth: 256, maxHeight: 256, name: "small", quality: 90, keepRatio: false, format: "webp", mime: "image/webp"},
	{maxWidth: 512, maxHeight: 512, name: "large", quality: 90, keepRatio: false, format: "webp", mime: "image/webp"},
}

var ImageProfiles = []profile{
	{maxWidth: 1024, maxHeight: 1024, name: "thumbnail", quality: 90, keepRatio: true, format: "webp", mime: "image/webp"},
	// {maxWidth: 1024, maxHeight: 1024, name: "medium", quality: 85, keepRatio: false, format: "webp", mime: "image/webp"},
	{maxWidth: 0, maxHeight: 0, name: "compress", quality: 90, keepRatio: true, format: "webp", mime: "image/webp"},
}

const (
	AvatarWidth  = 512
	AvatarHeight = 512

	ThumbnailWidth  = 1024
	ThumbnailHeight = 1024

	Quality = 85
)

// func (s *service) HandleImageCompressTask(_ context.Context, t *asynq.Task) error {
// 	var p ent.File
// 	if err := json.Unmarshal(t.Payload(), &p); err != nil {
// 		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
// 	}
// 	b, err := s.storageService.Get(p.Path, p.ID.String(), "original")

// 	if err != nil {
// 		return fmt.Errorf("get from s3 failed: %v", err)
// 	}

// 	return ImageCompress(&p, b, s.storageService)
// }

// func (s *service) HandleAvatarCompressTask(_ context.Context, t *asynq.Task) error {
// 	var p ent.File
// 	if err := json.Unmarshal(t.Payload(), &p); err != nil {
// 		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
// 	}
// 	b, err := s.storageService.Get(p.Path, p.ID.String(), "original")
// 	if err != nil {
// 		return fmt.Errorf("get from s3 failed: %v: %w", err, asynq.SkipRetry)
// 	}
// 	err = AvatarCompress(&p, b, s.storageService)
// 	if err != nil {
// 		return fmt.Errorf("compress failed: %v", err)
// 	}
// 	err = s.storageService.Delete(p.Path, p.ID, "original")
// 	if err != nil {
// 		return fmt.Errorf("upload failed: %v: %w", err, asynq.SkipRetry)
// 	}
// 	return nil

// }
