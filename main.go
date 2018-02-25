package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const imgW, imgH = 100, 100

func abort(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	f, err := os.Create("out/image.png")
	if err != nil {
		abort("Could not create the image file")
	}
	defer f.Close()

	r := image.Rectangle{image.Pt(0, 0), image.Pt(imgW, imgH)}
	img := image.NewRGBA(r)
	for i := 0; i < imgW; i++ {
		for j := 0; j < imgW; j++ {
			img.Set(i, j, color.RGBA{255, 0, 0, 255})
		}
	}

	png.Encode(f, img)
}
