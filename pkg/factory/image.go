package factory

import (
	"bytes"
	"fmt"
	"street/ent"
	"street/pkg/base3"
	"street/pkg/d"
	"street/pkg/utils"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func ImageCompress(p *ent.File, b []byte, s3Service base3.Prototype) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImageBlob(b)
	if err != nil {
		return fmt.Errorf("read blob failed: %v", err)
	}

	// thumbnail
	{
		mw := mw.Clone()
		defer mw.Destroy()
		if mw.GetImageWidth() > ResizeWidth || mw.GetImageHeight() > ResizeHeight {
			width, height := utils.ResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), ResizeWidth)
			err := mw.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
			if err != nil {
				return fmt.Errorf("resize failed: %v", err)
			}
		}
		err := mw.SetImageCompressionQuality(Quality)
		if err != nil {
			return fmt.Errorf("set compress quality failed: %v", err)
		}

		err = mw.SetImageFormat("webp")
		if err != nil {
			return fmt.Errorf("set format failed: %v", err)
		}
		mw.ResetIterator()
		filename := fmt.Sprintf("%s_%s.%s", p.Filename, "thumbnail", "webp")
		_, err = s3Service.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, filename, d.StringThumbnail, "image/webp")
		if err != nil {
			return fmt.Errorf("upload failed: %v", err)
		}
	}

	// original size compress
	{
		mw := mw.Clone()
		defer mw.Destroy()

		err := mw.SetImageCompressionQuality(Quality)
		if err != nil {
			return fmt.Errorf("set compress quality failed: %v", err)
		}

		err = mw.SetImageFormat("webp")
		if err != nil {
			return fmt.Errorf("set format failed: %v", err)
		}
		mw.ResetIterator()
		filename := fmt.Sprintf("%s_%s.%s", p.Filename, "compressed", "webp")
		_, err = s3Service.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, filename, d.StringCompressed, "image/webp")
		if err != nil {
			return fmt.Errorf("upload failed: %v", err)
		}
	}

	return nil
}

func AvatarCompress(p *ent.File, b []byte, s3Service base3.Prototype) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImageBlob(b)

	if err != nil {
		return fmt.Errorf("read image blob failed: %v", err)
	}

	err = mw.SetImageCompressionQuality(Quality)
	if err != nil {
		return fmt.Errorf("set compress quality failed: %v", err)
	}

	err = mw.SetImageFormat("webp")
	if err != nil {
		return fmt.Errorf("set format failed: %v", err)
	}

	if mw.GetImageWidth() > ResizeWidth || mw.GetImageHeight() > ResizeHeight {

		width, height := utils.ResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), ResizeWidth)

		err := mw.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
		if err != nil {
			return fmt.Errorf("resize failed: %v", err)
		}
	}
	if err != nil {
		return fmt.Errorf("crop failed: %v", err)
	}

	newFilename := p.Filename + ".webp"
	_, err = s3Service.PutSingle(bytes.NewReader(mw.GetImageBlob()), p.Path, p.ID, newFilename, "compressed", "image/webp")
	if err != nil {
		return fmt.Errorf("upload failed: %v", err)
	}
	return nil
}
