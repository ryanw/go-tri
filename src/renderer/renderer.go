package renderer

import (
  . "../geom"
  . "../mesh"
  . "../canvas"
)
type Renderer struct {
  Camera Camera
}

func (r *Renderer) RenderLineMesh(canvas *Canvas, mesh *Mesh) {
  camera := &r.Camera

  mvp := camera.Projection
  mvp = mvp.Multiply(camera.Transform.Matrix())
  mvp = mvp.Multiply(mesh.Transform.Matrix())

  for _, line := range mesh.Lines {
    start := mvp.TransformPoint3(mesh.Vertices[line[0]])
    end := mvp.TransformPoint3(mesh.Vertices[line[1]])

    canvas.DrawLine3D(start, end, Cell {
      Fg: 0xff00ff00,
      Bg: 0xff333333,
      Sprite: 'X',
    })
  }
}
