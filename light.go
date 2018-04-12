package main

import "math"

type Light interface {
	// ComputeLight calculates the total brightness of point given its normal
	// vector.
	ComputeLight(point, normal Vector) float64
}

type DirectionalLight struct {
	Intensity float64
	Direction Vector
}

func (light DirectionalLight) ComputeLight(_, normal Vector) float64 {
	cos := normal.Cos(light.Direction)
	return math.Max(light.Intensity*cos, 0)
}

type PointLight struct {
	Intensity float64
	Position  Vector
}

func (light PointLight) ComputeLight(point, normal Vector) float64 {
	cos := normal.Cos(light.Position.Minus(point))
	return math.Max(light.Intensity*cos, 0)
}

type AmbientLight struct {
	Intensity float64
}

func (light AmbientLight) ComputeLight(_, _ Vector) float64 {
	return light.Intensity
}
