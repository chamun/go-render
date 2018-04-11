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
	color Color
}

// NormalAt returns the normal vector on point p as a unit vector
func (s *Sphere) NormalAt(p Vector) Vector {
	return p.Minus(s.c).Normalize()
}

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

// Ray is the line that passes through point O in direction D
type Ray struct {
	O Vector
	D Vector
}

// At returns the point at parameter t of the ray
func (r *Ray) PointAt(t float64) Vector {
	return r.O.Add(r.D.Mult(t))
}

// IntersectSphere calculates where ray r hits the sphere s. It returns t1
// and t2, which are the lenghts at which the ray intersects with the sphere.
//
// There are three situations:
//   1. The ray enters and exits the sphere: t1 != t2
//   2. The ray is tangent to the sphere: t1 = t2
//   3. The ray dos not hit the sphere: t1 = t1 = +infinity
func (r *Ray) IntersectSphere(s Sphere) (float64, float64) {
	oc := r.O.Minus(s.c)
	k1 := r.D.Dot(r.D)
	k2 := 2 * oc.Dot(r.D)
	k3 := oc.Dot(oc) - s.r*s.r

	discriminant := k2*k2 - 4*k1*k3
	if discriminant < 0 {
		return math.Inf(1), math.Inf(1)
	}

	t1 := (-k2 + math.Sqrt(discriminant)) / (2 * k1)
	t2 := (-k2 - math.Sqrt(discriminant)) / (2 * k1)
	return t1, t2
}

type Scene struct {
	spheres []Sphere
	bgColor Color
	lights  []Light
}

// Render renders the scene to Canvas c
func (scene *Scene) Render(c Canvas) {
	O := Vector{0, 0, 0}
	for x := 0; x < c.Width(); x++ {
		for y := 0; y < c.Height(); y++ {
			D := canvasToViewPort(float64(x), float64(y), c)
			ray := Ray{O, D}
			color := scene.traceRay(ray, 1, math.Inf(1))
			c.Set(x, y, color)
		}
	}
}

// traceRay returns the color of the object r hits given that the intersection
// point lies on interval [tmin, tmax] of r. If there is no intersection, it
// returns the scene's background color.
func (scene *Scene) traceRay(r Ray, tmin, tmax float64) Color {
	closest_t := math.Inf(1)
	var closest_sphere *Sphere

	for i, sphere := range scene.spheres {
		t1, t2 := r.IntersectSphere(sphere)
		if t1 >= tmin && t1 <= tmax && t1 < closest_t {
			closest_t = t1
			closest_sphere = &scene.spheres[i]
		}
		if t2 >= tmin && t2 <= tmax && t2 < closest_t {
			closest_t = t2
			closest_sphere = &scene.spheres[i]
		}
	}

	if closest_sphere == nil {
		return scene.bgColor
	}

	p := r.PointAt(closest_t)
	n := closest_sphere.NormalAt(p)

	brightness := 0.0
	for _, light := range scene.lights {
		brightness += light.ComputeLight(p, n)
	}

	return closest_sphere.color.Mult(brightness)
}

// canvasToViewPort converts canvas coordinates to viewport coordinates.
func canvasToViewPort(x, y float64, c Canvas) Vector {
	fw, fh := float64(c.Width()), float64(c.Height())
	return Vector{(x - fw/2) / fw, (y - fh/2) / fh, 1}
}
