package main

import (
	"fmt"

	_ "image/jpeg"
	_ "image/png"

	"github.com/qeesung/image2ascii/convert"
)

func main() {
	imageFilename := "image/lilDocker.png"

	opt := convert.DefaultOptions
	opt.FixedWidth = 88
	opt.FixedHeight = 30
	opt.Colored = true

	converter := convert.NewImageConverter()

	fmt.Println(converter.ImageFile2ASCIIString(imageFilename, &opt))
}
