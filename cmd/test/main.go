package main

import (
	"fmt"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImage(".develop/57129082_p0.png")
	if err != nil {
		panic(err)
	}
	// width, height := utils.ResizeCalculator(mw.GetImageWidth(), mw.GetImageHeight(), 512)
	// mw.ResizeImage(width, height, imagick.FILTER_UNDEFINED)
	mw.SetImageCompressionQuality(85)
	if err != nil {
		panic(err)
	}
	// mw.StripImage()
	mw.SetSamplingFactors([]float64{4, 2, 0})
	// err = mw.GaussianBlurImage(0.05, 0.05)
	fmt.Printf("%v", err)
	err = mw.SetFormat("webp")
	if err != nil {
		panic(err)
	}
	err = mw.WriteImage("./.develop/c.webp")
	if err != nil {
		panic(err)
	}

}
