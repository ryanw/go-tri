package mesh

import (
	"math"
	. "tri/geom"
)

func NewLineMeshSphere() LineMesh {
	xSegments := 8.0
	ySegments := 6.0

	mesh := LineMesh{
		Transform: NewTransform(),
		Vertices:  []Point3{},
		Lines:     []Line{},
	}

	for y := ySegments * -0.5; y <= ySegments*0.5; y++ {
		lat := math.Pi * (y / ySegments)

		for x := xSegments * -0.5; x <= xSegments*0.5; x++ {
			lng := 2.0 * math.Pi * x / xSegments

			point := SphericalToCartesian(lng, lat)
			mesh.Vertices = append(mesh.Vertices, point)

			if len(mesh.Vertices) > 1 {
				// Horizontal line
				if x > xSegments*-0.5 {
					mesh.Lines = append(mesh.Lines, Line{int(len(mesh.Vertices) - 2), int(len(mesh.Vertices) - 1)})
				}

				// Vertical line
				if y > ySegments*-0.5 {
					mesh.Lines = append(mesh.Lines, Line{int(len(mesh.Vertices)) - int(xSegments) - 2, int(len(mesh.Vertices) - 1)})
				}
			}
		}

	}

	return mesh
}
