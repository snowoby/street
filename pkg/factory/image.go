package factory

import (
	"fmt"
	"street/ent"
	"street/pkg/base3"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func ImageCompress(p *ent.File, b []byte, s3Service base3.Prototype) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImageBlob(b)
	if err != nil {
		return fmt.Errorf("read blob failed: %v", err)
	}

	for _, profile := range ImageProfiles {
		err := process(mw, &profile)
		if err != nil {
			return fmt.Errorf("process failed: %v", err)
		}

		reader, err := output(mw, &profile)
		if err != nil {
			return fmt.Errorf("output failed: %v", err)
		}

		filename := fmt.Sprintf("%s_%s.%s", p.Filename, profile.name, profile.format)
		_, err = s3Service.PutSingle((reader), p.Path, p.ID, filename, profile.name, profile.mime)
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
		return fmt.Errorf("read blob failed: %v", err)
	}

	for _, profile := range AvatarProfiles {
		err := process(mw, &profile)
		if err != nil {
			return fmt.Errorf("process failed: %v", err)
		}

		reader, err := output(mw, &profile)
		if err != nil {
			return fmt.Errorf("output failed: %v", err)
		}

		filename := fmt.Sprintf("%s_%s.%s", p.Filename, profile.name, profile.format)
		_, err = s3Service.PutSingle((reader), p.Path, p.ID, filename, profile.name, profile.mime)
		if err != nil {
			return fmt.Errorf("upload failed: %v", err)
		}
	}

	return nil
}
