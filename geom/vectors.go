package geom

import (
  . "math"
)

type Vector2 [2]float64
type Vector3 [3]float64
type Vector4 [4]float64

func (v *Vector2) X() float64 {
  return v[0]
}

func (v *Vector2) Y() float64 {
  return v[1]
}

func (v *Vector3) X() float64 {
  return v[0]
}

func (v *Vector3) Y() float64 {
  return v[1]
}

func (v *Vector3) Z() float64 {
  return v[2]
}

func (v *Vector4) X() float64 {
  return v[0]
}

func (v *Vector4) Y() float64 {
  return v[1]
}

func (v *Vector4) Z() float64 {
  return v[2]
}

func (v *Vector4) W() float64 {
  return v[2]
}


func (v Vector3) Normalize() Vector3 {
  mag := v.Magnitude()
  return Vector3 {
    v.X() / mag,
    v.Y() / mag,
    v.Z() / mag,
  }
}

func (v *Vector3) Scale(scale float64) Vector3 {
  return Vector3 {
    v.X() * scale,
    v.Y() * scale,
    v.Z() * scale,
  }
}

func (left Vector3) Dot(right Vector3) float64 {
  dot := 0.0

  for i := range left {
    dot += left[i] * right[i]
  }

  return dot
}

func (v Vector3) Magnitude() float64 {
  return Sqrt(Pow(v.X(), 2) + Pow(v.Y(), 2) + Pow(v.Z(), 2))
}

func (p *Point3) Vector3FromPoint3() Vector3 {
  return Vector3 { p[0], p[1], p[2] }
}

func (v *Vector3) ToColor() Color {
  r := uint32(v.X() * 0xff) << 16;
  g := uint32(v.Y() * 0xff) << 8;
  b := uint32(v.Z() * 0xff) << 0;

  return Color(r + g + b)
}

func Vector3FromColor(i uint32) Vector3 {
  r := float64((i & 0xff0000) >> 16) / 0xff
  g := float64((i & 0x00ff00) >> 8)  / 0xff
  b := float64((i & 0x0000ff) >> 0)  / 0xff
  return Vector3 { r, g, b }
}
