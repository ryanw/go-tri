package geom

type Box2 [2]Point2

func (b Box2) MinX() float64 {
	x0, x1 := b[0].X(), b[1].X()
	if x0 < x1 {
		return x0
	}
	return x1
}
func (b Box2) MaxX() float64 {
	x0, x1 := b[0].X(), b[1].X()
	if x0 > x1 {
		return x0
	}
	return x1
}

func (b Box2) MinY() float64 {
	y0, y1 := b[0].Y(), b[1].Y()
	if y0 < y1 {
		return y0
	}
	return y1
}
func (b Box2) MaxY() float64 {
	y0, y1 := b[0].Y(), b[1].Y()
	if y0 > y1 {
		return y0
	}
	return y1
}

// Returns a Box2 that encloses the triangle.
// Ignores the Z axis.
func (t Triangle3) IntoBox2() Box2 {
	minX, maxX := t[0].X(), t[0].X()
	minY, maxY := t[0].Y(), t[0].Y()

	for _, v := range t {
		if v.X() < minX {
			minX = v.X()
		}
		if v.X() > maxX {
			maxX = v.X()
		}
		if v.Y() < minY {
			minY = v.Y()
		}
		if v.Y() > maxY {
			maxY = v.Y()
		}
	}

	return Box2{
		Point2{minX, minY},
		Point2{maxX, maxY},
	}
}
