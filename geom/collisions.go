package geom

func (t Triangle3) IntersectsBox2(box Box2) bool {
	// Approximate hit using the bounding box
	return t.IntoBox2().IntersectsBox2(box)
}

func (a Box2) IntersectsBox2(b Box2) bool {
	// Simple AABB hit test
	return a.MinX() <= b.MaxX() && a.MaxX() >= b.MinX() && a.MinY() <= b.MaxY() && a.MaxY() >= b.MinY()
}
