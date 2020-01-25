package geom

import "math"

type Vector2 [2]float64
type Vector3 [3]float64
type Vector4 [4]float64

type Matrix4 [16]float64

type Line2 [2]Point2
type Line3 [2]Point3
type Triangle2 [3]Point2
type Triangle3 [3]Point3


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

func SphericalToCartesian(lon, lat float64) Point3 {
  x := math.Cos(lat) * math.Sin(lon);
  y := math.Sin(lat);
  z := math.Cos(lat) * math.Cos(lon);
  return Point3 { x, y, z }
}
