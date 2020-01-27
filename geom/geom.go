package geom

import "math"

type Line2 [2]Point2
type Line3 [2]Point3
type Triangle2 [3]Point2
type Triangle3 [3]Point3

func SphericalToCartesian(lon, lat float64) Point3 {
	x := math.Cos(lat) * math.Sin(lon)
	y := math.Sin(lat)
	z := math.Cos(lat) * math.Cos(lon)
	return Point3{x, y, z}
}
