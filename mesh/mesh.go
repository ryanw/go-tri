package mesh

import (
	. "tri/geom"
)

type Line [2]int

type LineMesh struct {
	Transform Transform
	Vertices  []Point3
	Lines     []Line
}

type TriangleMesh struct {
	Transform Transform
	Vertices  []Point3
	Triangles [][3]int
	Normals   []Vector3
	Colors    []uint32
}
