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
	defer mw.Destroy()

	err = mw.SetImageCompressionQuality(80)
	if err != nil {
		fmt.Printf("set compress quality failed: %v", err)
	}

	err = mw.SetImageFormat("webp")
	if err != nil {
		fmt.Printf("set format failed: %v", err)
	}

	bytes := mw.GetImageBlob()
	fmt.Println(len(bytes))
	err = mw.WriteImage("./.develop/c")
	if err != nil {
		panic(err)
	}

}
