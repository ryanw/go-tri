package mesh

import (
	. "tri/geom"
)

func NewTriangleMeshPlane(w, h int) TriangleMesh {

	mesh := TriangleMesh{
		Transform: NewTransform(),
		Vertices:  []Point3{},
		Triangles: [][3]int{},
		Normals:   []Vector3{},
		Colors:    []uint32{},
	}

	colors := []uint32{
		0xff555555,
		0xffaaaaaa,
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			fx := float64(x) - float64(w)/2.0
			fy := float64(y) - float64(h)/2.0
			idx := len(mesh.Vertices)
			cidx := (x + y) % len(colors)
			color := colors[cidx]
			mesh.Vertices = append(
				mesh.Vertices,
				Point3{fx + 0, 0, fy + 1},
				Point3{fx + 1, 0, fy + 1},
				Point3{fx + 0, 0, fy + 0},
				Point3{fx + 1, 0, fy + 0},
			)
			mesh.Triangles = append(
				mesh.Triangles,
				[3]int{idx + 0, idx + 1, idx + 2},
				[3]int{idx + 2, idx + 1, idx + 3},
			)
			mesh.Normals = append(
				mesh.Normals,
				Vector3{0, -1, 0},
				Vector3{0, -1, 0},
			)
			mesh.Colors = append(
				mesh.Colors,
				color, color,
			)
		}
	}

	return mesh
}
