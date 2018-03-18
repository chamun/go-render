package main

import (
	"image/color"
	"math"
)

type Canvas interface {
	Set(x, y int, c color.Color)
	Width() int
	Height() int
}

type Sphere struct {
	c     Vector
	r     float64
	color color.Color
}

type Scene struct {
	spheres []Sphere
}

// Render renders the scene to Canvas c
func (scene *Scene) Render(c Canvas) {
	O := Vector{0, 0, 0}
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			D := canvasToViewPort(float64(x), float64(y), c)
			color := scene.traceRay(O, D, 1, math.Inf(1))
			c.Set(x, y, color)
		}
	}
}

// traceRay returns the color of the object that the ray that passes on point o
// with direction d hits given that the intersection point lies on interval
// [tmin, tmax] of the ray. If there is no intersection, it returns the color
// white.
func (scene *Scene) traceRay(o, d Vector, tmin, tmax float64) color.Color {
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

// canvasToViewPort converts canvas coordinates to viewport coordinates.
func canvasToViewPort(x, y float64, c Canvas) Vector {
	fw, fh := float64(c.Width()), float64(c.Height())
	return Vector{(x - fw/2) / fw, -(y - fh/2) / fh, 1}
}

// intersectRaySphere calculates if the ray that passes on point o with
// direction d hits the sphere s. It returns t1 and t2, which are the lenghts at
// which the ray intersects with the sphere.
//
// There are three situations:
//   1. The ray enters and exits the sphere: t1 != t2
//   2. The ray is tangent to the sphere: t1 = t2
//   3. The ray dos not hit the sphere: t1 = t1 = +infinity
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