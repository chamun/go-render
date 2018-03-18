package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Sphere struct {
	c     Vector
	r     float64
	color color.Color
}

type Scene struct {
	spheres []Sphere
}

const imgW, imgH = 500, 500

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}
var blue = color.RGBA{0, 0, 255, 255}
var white = color.RGBA{255, 255, 255, 255}

var scene = Scene{[]Sphere{
	Sphere{Vector{0, 1, -3}, 1, red},
	Sphere{Vector{2, 0, -4}, 1, blue},
	Sphere{Vector{-2, 0, -4}, 1, green}}}

func abort(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func canvasToViewPort(x, y float64) Vector {
	return Vector{(x - imgW/2) / imgW, -(y - imgH/2) / imgH, 1}
}

func intersectRaySphere(o, d Vector, s Sphere) (float64, float64) {
	oc := Minus(o, s.c)
	k1 := d.Dot(d)
	k2 := 2 * oc.Dot(d)
	k3 := oc.Dot(oc) - s.r*s.r

	discriminant := k2*k2 - 4*k1*k3
	if discriminant < 0 {
		return math.Inf(1), math.Inf(1)
	}

	t1 := (-k2 + math.Sqrt(discriminant)) / (2 * k1)
	t2 := (-k2 - math.Sqrt(discriminant)) / (2 * k1)
	return t1, t2
}

func traceRay(o, d Vector, tmin, tmax float64) color.Color {
	closest_t := math.Inf(1)
	closest_sphere := Sphere{color: white}
	for _, sphere := range scene.spheres {
		t1, t2 := intersectRaySphere(o, d, sphere)
		if t1 >= tmin && t1 <= tmax && t1 < closest_t {
			closest_t = t1
			closest_sphere = sphere
		}
		if t2 >= tmin && t2 <= tmax && t2 < closest_t {
			closest_t = t2
			closest_sphere = sphere
		}
	}
	return closest_sphere.color
}

func main() {
	f, err := os.Create("out/image.png")
	if err != nil {
		abort("Could not create the image file")
	}
	defer f.Close()

	r := image.Rectangle{image.Pt(0, 0), image.Pt(imgW, imgH)}
	img := image.NewRGBA(r)

	O := Vector{0, 0, 0}
	for x := 0; x < imgW; x++ {
		for y := 0; y < imgH; y++ {
			D := canvasToViewPort(float64(x), float64(y))
			color := traceRay(O, D, 1, math.Inf(1))
			img.Set(x, y, color)
		}
	}

	png.Encode(f, img)
}
