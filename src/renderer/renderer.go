package renderer

import (
  . "../geom"
  . "../mesh"
  . "../canvas"
)
type Renderer struct {
  Camera Camera
}

func (r *Renderer) RenderLineMesh(canvas *Canvas, mesh *LineMesh) {
  camera := &r.Camera

  mvp := camera.Projection
  mvp = mvp.Multiply(camera.Transform.Matrix())
  mvp = mvp.Multiply(mesh.Transform.Matrix())

  for _, line := range mesh.Lines {
    start := mvp.TransformPoint3(mesh.Vertices[line[0]])
    end := mvp.TransformPoint3(mesh.Vertices[line[1]])

    canvas.DrawLine3D(start, end, Cell {
      Fg: 0xff00ff00,
      Bg: 0xffaa0000,
      Sprite: ' ',
    })
  }
}

func (r *Renderer) RenderTriangleMesh(canvas *Canvas, mesh *TriangleMesh) {
  camera := &r.Camera

  mvp := camera.Projection
  mvp = mvp.Multiply(camera.Transform.Matrix())
  mvp = mvp.Multiply(mesh.Transform.Matrix())

  for i, tri := range mesh.Triangles {
    p0 := mvp.TransformPoint3(mesh.Vertices[tri[0]])
    p1 := mvp.TransformPoint3(mesh.Vertices[tri[1]])
    p2 := mvp.TransformPoint3(mesh.Vertices[tri[2]])

    triangle := TriangleFloat {
      [2]float64{ p0[0], p0[1] },
      [2]float64{ p1[0], p1[1] },
      [2]float64{ p2[0], p2[1] },
    }

    color := Color(mesh.Colors[i])

    canvas.DrawVectorTriangle(triangle, Cell {
      Fg: color,
      Bg: color,
      Sprite: '.',
    })
  }
}
