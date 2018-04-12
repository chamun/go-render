package main

import (
	"image/color"
	"math"
)

var red = Color{255, 0, 0}
var green = Color{0, 255, 0}
var blue = Color{0, 0, 255}
var yellow = Color{255, 255, 0}
var white = Color{255, 255, 255}

// Color is a vector that implements the color.Color interface
type Color Vector

// RGBA clamps the color components to 255 and delegates the conversion to
// color.RGBA.RGBA()
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	r := uint8(math.Min(c.X, 255))
	g := uint8(math.Min(c.Y, 255))
	b := uint8(math.Min(c.Z, 255))
	return color.RGBA{r, g, b, 255}.RGBA()
}

// Mult returns the color represented by a multiplied by b
func (a Color) Mult(b float64) Color {
	return Color(Vector(a).Mult(b))
}
