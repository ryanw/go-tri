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

func (m *TriangleMesh) Triangle(index int) Triangle3 {
	if index >= len(m.Triangles) {
		return Triangle3{}
	}
	tri := m.Triangles[index]
	return Triangle3{
		m.Vertices[tri[0]],
		m.Vertices[tri[1]],
		m.Vertices[tri[2]],
	}
}
