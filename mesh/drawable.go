package mesh

import (
	. "tri/geom"
	. "tri/renderer"
)

func (m *TriangleMesh) DrawTriangles(ch chan<- Polygon) {
	// Move to world space
	model := m.Transform.Matrix()
	for i, triIndexes := range m.Triangles {
		triangle := model.TransformTriangle3(Triangle3{
			m.Vertices[triIndexes[0]],
			m.Vertices[triIndexes[1]],
			m.Vertices[triIndexes[2]],
		})
		color := m.Colors[i]
		ch <- Polygon{Shape: triangle, Color: color}
	}
}
