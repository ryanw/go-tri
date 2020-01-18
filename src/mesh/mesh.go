package mesh

import (
  . "../geom"
  . "../terminal"
)

type Line [2]int64

type Mesh struct {
  Transform Transform
  Vertices []Point3
  Lines []Line
}

func (self Mesh) Draw(term *Terminal, camera Camera, char rune) {
  mvp := camera.Projection
  mvp = mvp.Multiply(camera.Transform.Matrix())
  mvp = mvp.Multiply(self.Transform.Matrix())

  for _, line := range self.Lines {
    start := mvp.TransformPoint3(self.Vertices[line[0]])
    end := mvp.TransformPoint3(self.Vertices[line[1]])

    term.DrawLine(start, end, char)
  }
}

