package geom

import (
	"testing"
)

func TestLineIntersectsPlane3Hits(t *testing.T) {
	line := Line3{
		Point3{3.0, -10.0, -4.0},
		Point3{3.0, 10.0, -4.0},
	}
	plane := Plane3{
		Point3{0.0, 7.0, 0.0},
		Vector3{0.0, -1.0, 0.0},
	}

	hit, point := line.IntersectsPlane3(plane)

	if !hit {
		t.Error("Line didn't hit plane")
	}

	assertValuesEqual(t, point[:], []float64{3.0, 7.0, -4.0})
}

func TestLineIntersectsPlane3Misses(t *testing.T) {
	line := Line3{
		Point3{3.0, -10.0, -4.0},
		Point3{3.0, 10.0, -4.0},
	}
	plane := Plane3{
		Point3{0.0, -20.0, 0.0},
		Vector3{0.0, -1.0, 0.0},
	}

	hit, _ := line.IntersectsPlane3(plane)

	if hit {
		t.Error("Line shouldn't hit plane")
	}
}
