package mesh

import (
	"github.com/aquilax/go-perlin"
	. "tri/geom"
)

func NewTerrainMesh(w, h int) TriangleMesh {
	mesh := TriangleMesh{
		Transform: NewTransform(),
		Vertices:  []Point3{},
		Triangles: [][3]int{},
		Normals:   []Vector3{},
		Colors:    []uint32{},
	}

	noise := perlin.NewPerlin(2, 2, 3, 666)
	randomHeight := func(p Point3) Point3 {
		scale := 0.1
		p[1] = noise.Noise2D(scale*p.X(), scale*p.Z()) * 40.0
		return p
	}
	colors := []uint32{
		0xffaaaa00,
		0xff228800,
		0xff22aa00,
		0xff00cc00,
		0xff779966,
		0xff779966,
		0xff779966,
		0xff779966,
		0xff779966,
	}
	randomColor := func(p Point3) uint32 {
		scale := 0.2
		val := 0.5 + noise.Noise2D(scale*p.X(), scale*p.Z())
		// FIXME why does it go outside 0.0 - 1.0
		if val > 1.0 {
			val = 1.0
		}
		if val < 0.0 {
			val = 0.0
		}
		numColors := float64(len(colors) - 1)
		shade := uint32(numColors - val*numColors)
		return colors[shade]
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			fx := float64(x) - float64(w)/2.0
			fz := float64(y) - float64(h)/2.0

			idx := len(mesh.Vertices)

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

			tri := len(mesh.Triangles)
			color1 := randomColor(mesh.Triangle(tri - 2).Centroid())
			color2 := randomColor(mesh.Triangle(tri - 1).Centroid())

			mesh.Colors = append(mesh.Colors, color1, color2)
		}
	}

	return mesh
}
