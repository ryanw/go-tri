package scene

import (
	. "tri/mesh"
	. "tri/renderer"
)

type Scene struct {
	Meshes []TriangleMesh
}

func NewScene() Scene {
	return NewSceneWith([]TriangleMesh{})
}

func NewSceneWith(meshes []TriangleMesh) Scene {
	return Scene{meshes}
}

func (s *Scene) Mesh(idx int) *TriangleMesh {
	if idx < 0 || idx >= len(s.Meshes) {
		return nil
	}
	return &s.Meshes[idx]
}

func (s *Scene) Add(mesh TriangleMesh) int {
	s.Meshes = append(s.Meshes, mesh)
	return len(s.Meshes) - 1
}

func (s *Scene) DrawTriangles(ch chan<- Polygon) {
	for _, mesh := range s.Meshes {
		mesh.DrawTriangles(ch)
	}
}
