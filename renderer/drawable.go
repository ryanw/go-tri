package renderer

import "tri/geom"

type Polygon struct {
	Shape geom.Triangle3
	Color uint32
}

type Drawable interface {
	DrawTriangles(ch chan<- Polygon)
}
