package geom

import (
	. "math"
)

type Point2 [2]float64
type Point3 [3]float64
type Point4 [4]float64

func (p *Point2) X() float64 {
	return p[0]
}

func (p *Point2) Y() float64 {
	return p[1]
}

func (p *Point3) X() float64 {
	return p[0]
}

func (p *Point3) Y() float64 {
	return p[1]
}

func (p *Point3) Z() float64 {
	return p[2]
}

func (p Point3) Ceil() Point3 {
	return Point3{
		Ceil(p[0]),
		Ceil(p[1]),
		Ceil(p[2]),
	}
}
func (p Point2) Ceil() Point2 {
	return Point2{
		Ceil(p[0]),
		Ceil(p[1]),
	}
}

func (p Point3) Floor() Point3 {
	return Point3{
		Floor(p[0]),
		Floor(p[1]),
		Floor(p[2]),
	}
}
func (p Point2) Floor() Point2 {
	return Point2{
		Floor(p[0]),
		Floor(p[1]),
	}
}

func (p Point3) Round() Point3 {
	return Point3{
		Round(p[0]),
		Round(p[1]),
		Round(p[2]),
	}
}
func (p Point2) Round() Point2 {
	return Point2{
		Round(p[0]),
		Round(p[1]),
	}
}

func (p Point3) Normalize() Point3 {
	mag := p.Magnitude()
	return Point3{
		p.X() / mag,
		p.Y() / mag,
		p.Z() / mag,
	}
}

func (p Point3) Magnitude() float64 {
	return Sqrt(Pow(p.X(), 2) + Pow(p.Y(), 2) + Pow(p.Z(), 2))
}

func (v *Vector3) Point3FromVector3() Point3 {
	return Point3{v[0], v[1], v[2]}
}
