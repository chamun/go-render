package main

// Type Vector represents a vector in a three dimensional space.
type Vector struct {
	X, Y, Z float64
}

// Dot returns the dot product between vectors a and b.
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Minus returns the vector represented by b - a.
func Minus(a, b Vector) Vector {
	return Vector{b.X - a.X, b.Y - a.Y, b.Z - a.Z}
}
