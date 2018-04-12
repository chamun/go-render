package main

import "math"

// Type Vector represents a vector in a three dimensional space.
type Vector struct {
	X, Y, Z float64
}

// Dot returns the dot product between vectors a and b.
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Minus returns the vector represented by a - b.
func (a Vector) Minus(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Cos returns the cosine of the angle between a and b
func (a Vector) Cos(b Vector) float64 {
	return a.Dot(b) / (a.Length() * b.Length())
}

// Length returns the length of vector a
func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

// Add returns the vector represented by a + b
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Mult returns the vector represented by a multiplied by b
func (a Vector) Mult(b float64) Vector {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

// Normalize returns the unit vector of a
func (a Vector) Normalize() Vector {
	return a.Mult(1 / a.Length())
}
