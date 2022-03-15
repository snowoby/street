package factory

import (
	"bytes"
	"fmt"
	"street/ent"
	"street/pkg/utils"

	"gopkg.in/gographics/imagick.v3/imagick"
)

type imageProcessTask struct {
	file         *ent.File
	profile      *profile
	storage      *storageService
	mw           *imagick.MagickWand
	resultReader *bytes.Reader
}

// func (task *imageProcessTask) Run() error {
// 	err := task.process()
// 	if err != nil {
// 		return err
// 	}
// 	err = task.output()
// 	if err != nil {
// 		return err
// 	}

// 	err = task.put()
// 	return err
// }

// func NewImageProcessTask(storage *storageService, file *ent.File, profile *profile, mw *imagick.MagickWand) *imageProcessTask {
// 	return &imageProcessTask{
// 		file:    file,
// 		profile: profile,
// 		storage: storage,
// 		mw:      mw,
// 	}
// }

func process(mw *imagick.MagickWand, profile *profile) error {
	mw.ResetIterator()
	err := mw.SetImageCompressionQuality(uint(profile.quality))
	if err != nil {
		return fmt.Errorf("set compress quality failed: %v", err)
	}

	if profile.maxWidth*profile.maxWidth != 0 {
		width := uint(profile.maxWidth)
		height := uint(profile.maxHeight)
		if profile.keepRatio {
			width, height = utils.KeepRatioResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), uint(profile.maxWidth), uint(profile.maxHeight))
		}
		err = mw.AdaptiveResizeImage(width, height)
		if err != nil {
			return fmt.Errorf("resize failed: %v", err)
		}
	}
	return nil
}

func output(mw *imagick.MagickWand, profile *profile) (*bytes.Reader, error) {
	err := mw.SetImageFormat(profile.format)
	if err != nil {
		return nil, fmt.Errorf("set format failed: %v", err)
	}
	mw.ResetIterator()
	return bytes.NewReader(mw.GetImageBlob()), nil
}
