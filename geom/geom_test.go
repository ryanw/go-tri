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

	point := line.IntersectsPlane3(plane)

	assertValuesEqual(t, point[:], []float64{3.0, 7.0, -4.0})
}

func TestTriangle3Centroid(t *testing.T) {
	tri := Triangle3{
		Point3{36, 12, 5},
		Point3{46, 40, 16},
		Point3{65, 20, 9},
	}

	result := tri.Centroid()

	assertPoint3Equal(t, result, Point3{49, 24, 10})
}
