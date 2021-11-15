package vec3

import "math"

// Vector3 struct
type Vector3 struct {
	X, Y, Z float64
}

// Add two Vector3's together
func Add(a, b Vector3) Vector3 {
	return Vector3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Subtract two vectors
func Subtract(a, b Vector3) Vector3 {
	return Vector3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Multiply a Vector
func Mult(a Vector3, b float64) Vector3 {
	return Vector3{a.X * b, a.Y * b, a.Z * b}
}

//  Divide a Vector
func Div(a Vector3, b float64) Vector3 {
	return Vector3{a.X / b, a.Y / b, a.Z / b}
}

// Length of a Vector3
func (a Vector3) Length() float64 {
	return float64(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
}

// Length aquared of a Vector3
func (a Vector3) LengthSqr() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

// Distance of a Vector3
func Distance(a, b Vector3) float64 {
	xDiff := a.X - b.X
	yDiff := b.Y - b.Y
	zDiff := a.Z - b.Z
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff)
}

// DistanceSquared of a Vector3
func DistanceSquared(a, b Vector3) float64 {
	xDiff := a.X - b.X
	yDiff := b.Y - b.Y
	zDiff := a.Z - b.Z
	return xDiff*xDiff + yDiff*yDiff + zDiff*zDiff
}

// Normalize a Vector
func Normalize(a Vector3) Vector3 {
	len := a.Length()
	return Vector3{a.X / len, a.Y / len, a.Z / len}
}

// Dot product of two vectors
func Dot(a, b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross product of two vectors
func Cross(a, b Vector3) Vector3 {
	i := a.Y*b.Z - a.Z*b.Y
	j := a.Z*b.X - a.X*b.Z
	k := a.X*b.Y - a.Y*b.X
	return Vector3{i, j, k}
}

// Provide zero vector
func Zero() Vector3 {
	return Vector3{0, 0, 0}
}

// Create new vector3

func New(x, y, z float64) Vector3 {
	return Vector3{x, y, z}
}
