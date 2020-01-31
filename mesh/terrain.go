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

	noise := perlin.NewPerlin(2, 2, 1, 123)
	noise2 := perlin.NewPerlin(2, 2, 1, 123)
	randomHeight := func(p Point3) Point3 {
		scale := 0.15
		p[1] = 10 + noise.Noise2D(scale*p.X(), scale*p.Z())*25.0
		return p
	}
	colors := []uint32{
		0xff000044,
		0xff000088,
		0xff0000aa,
		0xff0000ff,
		0xffaaaa00,
		0xffffaa00,
		0xff228800,
		0xff00aa00,
		0xff00aa00,
		0xff00dd00,
		0xff22aa00,
		0xff99aa99,
		0xff999999,
		0xffaaaaaa,
		0xffbbbbbb,
	}
	randomColor := func(p Point3) uint32 {
		scale := 0.2
		val := 0.5 + noise2.Noise2D(scale*p.X(), scale*p.Z())
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
			color := randomColor(Point3{fx, 0, fz})

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
