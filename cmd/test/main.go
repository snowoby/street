package main

import "gopkg.in/gographics/imagick.v3/imagick"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImage(".develop/8.png")
	if err != nil {
		panic(err)
	}
	err = mw.SetImageCompressionQuality(90)
	if err != nil {
		panic(err)
	}
	err = mw.SetFormat("webp")
	if err != nil {
		panic(err)
	}
	err = mw.WriteImage("./.develop/c.webp")
	if err != nil {
		panic(err)
	}

}
