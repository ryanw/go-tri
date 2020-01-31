package geom

import "math"

type Line2 [2]Point2
type Line3 [2]Point3
type Triangle2 [3]Point2
type Triangle3 [3]Point3
type Plane3 struct {
	Point  Point3
	Normal Vector3
}

func (t Triangle3) Centroid() Point3 {
	x := (t[0][0] + t[1][0] + t[2][0]) / 3
	y := (t[0][1] + t[1][1] + t[2][1]) / 3
	z := (t[0][2] + t[1][2] + t[2][2]) / 3

	return Point3{x, y, z}
}

func (p Plane3) Distance() float64 {
	return p.Point.ToVector3().Dot(p.Normal)
}

func (p Point3) DistanceToPlane3(plane Plane3) float64 {
	normal := plane.Normal.Normalize()
	distance := normal.Dot(plane.Point.ToVector3())
	return normal.X()*p.X() + normal.Y()*p.Y() + normal.Z()*p.Z() - distance
}

func (l Line3) IntersectsPlane3(p Plane3) Point3 {
	normal := p.Normal.Normalize()
	distance := p.Distance()
	start, end := l[0].ToVector3(), l[1].ToVector3()
	startDist := start.Dot(normal)
	endDist := end.Dot(normal)
	t := (distance - startDist) / (endDist - startDist)

	startToEnd := end.Sub(start)
	toIntersect := startToEnd.Scale(t)

	intersect := start.Add(toIntersect).ToPoint3()
	return intersect
}

func SphericalToCartesian(lon, lat float64) Point3 {
	x := math.Cos(lat) * math.Sin(lon)
	y := math.Sin(lat)
	z := math.Cos(lat) * math.Cos(lon)
	return Point3{x, y, z}
}
