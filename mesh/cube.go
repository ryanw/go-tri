package mesh

import (
	. "tri/geom"
)

func NewTriangleMeshCube() TriangleMesh {
	return TriangleMesh{
		Transform: NewTransform(),
		Vertices: []Point3{
			Point3{-1, -1, 1},
			Point3{1, -1, 1},
			Point3{1, 1, 1},
			Point3{-1, 1, 1},

			Point3{-1, -1, -1},
			Point3{1, -1, -1},
			Point3{1, 1, -1},
			Point3{-1, 1, -1},
		},
		Triangles: [][3]int{
			// Front
			[3]int{0, 1, 2},
			[3]int{2, 3, 0},

			// Back
			[3]int{4, 5, 6},
			[3]int{6, 7, 4},

			// Left
			[3]int{0, 3, 7},
			[3]int{0, 7, 4},

			// Right
			[3]int{1, 2, 6},
			[3]int{1, 6, 5},

			// Top
			[3]int{0, 1, 5},
			[3]int{5, 4, 0},

			// Bottom
			[3]int{3, 2, 6},
			[3]int{6, 7, 3},
		},
		Normals: []Vector3{
			Vector3{0, 0, 1},
			Vector3{0, 0, 1},

			Vector3{0, 0, -1},
			Vector3{0, 0, -1},

			Vector3{-1, 0, 0},
			Vector3{-1, 0, 0},

			Vector3{1, 0, 0},
			Vector3{1, 0, 0},

			Vector3{0, -1, 0},
			Vector3{0, -1, 0},

			Vector3{0, 1, 0},
			Vector3{0, 1, 0},
		},
		Colors: []uint32{
			0xffff0000,
			0xffff0000,

			0xff00ff00,
			0xff00ff00,

			0xff0000ff,
			0xff0000ff,

			0xffffff00,
			0xffffff00,

			0xffff00ff,
			0xffff00ff,

			0xff00ffff,
			0xff00ffff,
		},
	}

}

func NewLineMeshCube() LineMesh {
	return LineMesh{
		Transform: NewTransform(),
		Vertices: []Point3{
			Point3{-1, -1, -1},
			Point3{1, -1, -1},
			Point3{1, 1, -1},
			Point3{-1, 1, -1},

			Point3{-1, -1, 1},
			Point3{1, -1, 1},
			Point3{1, 1, 1},
			Point3{-1, 1, 1},
		},
		Lines: []Line{
			// Front
			Line{0, 1},
			Line{1, 2},
			Line{2, 3},
			Line{3, 0},

			// Back
			Line{4, 5},
			Line{5, 6},
			Line{6, 7},
			Line{7, 4},

			// Top
			Line{0, 4},
			Line{4, 5},
			Line{5, 1},
			Line{1, 0},

			// Bottom
			Line{2, 6},
			Line{6, 7},
			Line{7, 3},
			Line{3, 2},
		},
	}

}
