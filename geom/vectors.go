package geom

import (
	. "math"
)

// A 2D Vector
type Vector2 [2]float64

// A 3D Vector
type Vector3 [3]float64

// A 4D Vector
type Vector4 [4]float64

// Get the value of the X axis
func (v Vector2) X() float64 {
	return v[0]
}

// Get the value of the Y axis
func (v Vector2) Y() float64 {
	return v[1]
}

// Get the value of the X axis
func (v Vector3) X() float64 {
	return v[0]
}

// Get the value of the Y axis
func (v Vector3) Y() float64 {
	return v[1]
}

// Get the value of the Y axis
func (v Vector3) Z() float64 {
	return v[2]
}

// Get the value of the Z axis
func (v Vector4) X() float64 {
	return v[0]
}

// Get the value of the Y axis
func (v Vector4) Y() float64 {
	return v[1]
}

// Get the value of the Z axis
func (v Vector4) Z() float64 {
	return v[2]
}

// Get the value of the W axis
func (v Vector4) W() float64 {
	return v[2]
}

func (v Vector3) ToPoint3() Point3 {
	return Point3(v)
}

func (p Point3) ToVector3() Vector3 {
	return Vector3(p)
}

// Scale the vector so the magnitude is 1.0
func (v Vector3) Normalize() Vector3 {
	mag := v.Magnitude()
	return Vector3{
		v.X() / mag,
		v.Y() / mag,
		v.Z() / mag,
	}
}

func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{
		v1[0] - v2[0],
		v1[1] - v2[1],
		v1[2] - v2[2],
	}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		v1[0] + v2[0],
		v1[1] + v2[1],
		v1[2] + v2[2],
	}
}

// Scale the vector by an amount
func (v Vector3) Scale(scale float64) Vector3 {
	return Vector3{
		v.X() * scale,
		v.Y() * scale,
		v.Z() * scale,
	}
}

// Get the dot product between two vectors
func (left Vector3) Dot(right Vector3) float64 {
	dot := 0.0

	for i := range left {
		dot += left[i] * right[i]
	}

	return dot
}

// Get the magnitude (length) of the vector
func (v Vector3) Magnitude() float64 {
	return Sqrt(Pow(v.X(), 2) + Pow(v.Y(), 2) + Pow(v.Z(), 2))
}

// Convert a Point3 into a Vector3
func (p *Point3) Vector3FromPoint3() Vector3 {
	return Vector3{p[0], p[1], p[2]}
}

// Convert a vector storing RGB into a single 32bit integer: 0x00RRGGBB
func (v Vector3) ToColor() Color {
	r := uint32(v.X()*0xff) << 16
	g := uint32(v.Y()*0xff) << 8
	b := uint32(v.Z()*0xff) << 0

	return Color(r + g + b)
}

// Convert a 32bit uint (0x00RRGGBB) into a vector storing RGB
func Vector3FromColor(i uint32) Vector3 {
	r := float64((i&0xff0000)>>16) / 0xff
	g := float64((i&0x00ff00)>>8) / 0xff
	b := float64((i&0x0000ff)>>0) / 0xff
	return Vector3{r, g, b}
}

// Scale the vector by an amount
func (m Vector4) Scale(scale float64) Vector4 {
	return Vector4{
		m[0] * scale,
		m[1] * scale,
		m[2] * scale,
		m[3] * scale,
	}
}
