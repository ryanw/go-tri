package mesh

import (
	"math/rand"
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


	heights := map[[2]int]float64{}
	randomHeight := func(p Point3) Point3 {
		x := int(p.X())
		y := int(p.Z())
		key := [2]int{x, y}
		val := heights[key]
		if val == 0.0 {
			val = rand.Float64() * 4.0
			heights[key] = val
		}
		p[1] = val
		return p
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			fx := float64(x) - float64(w)/2.0
			fz := float64(y) - float64(h)/2.0

			idx := len(mesh.Vertices)
			r := uint32(rand.Float64()*float64(0xff)) << 16
			g := uint32(rand.Float64()*float64(0xff)) << 8
			b := uint32(rand.Float64()*float64(0xff)) << 0
			color := 0xff000000 + r + g + b
			mesh.Vertices = append(
				mesh.Vertices,
				randomHeight(Point3{fx + 0, 0, fz + 1}),
				randomHeight(Point3{fx + 1, 0, fz + 1}),
				randomHeight(Point3{fx + 0, 0, fz + 0}),
				randomHeight(Point3{fx + 1, 0, fz + 0}),
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
