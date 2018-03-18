package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}
var white = color.RGBA{255, 255, 255, 255}

type CanvasAdapter struct {
	draw.Image
}

func (ca *CanvasAdapter) Width() int {
	return ca.Bounds().Dx()
}

func (ca *CanvasAdapter) Height() int {
	return ca.Bounds().Dy()
}

func main() {
	f, err := os.Create("out/image.png")
	if err != nil {
		fmt.Println("Could not create the image file")
		os.Exit(1)
	}
	defer f.Close()

	scene := Scene{[]Sphere{
		Sphere{Vector{0, 1, -3}, 1, red},
		Sphere{Vector{2, 0, -4}, 1, blue},
		Sphere{Vector{-2, 0, -4}, 1, green}},
		white}

	r := image.Rectangle{image.Pt(0, 0), image.Pt(500, 500)}
	img := image.NewRGBA(r)
	scene.Render(&CanvasAdapter{img})

	png.Encode(f, img)
}
