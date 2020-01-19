package mesh

import (
  "math"
  . "../geom"
  . "../terminal"
)

type Line [2]int64

type Mesh struct {
  Transform Transform
  Vertices []Point3
  Lines []Line
}

func (self *Mesh) Draw(term *Terminal, camera Camera, char rune) {
  mvp := camera.Projection
  mvp = mvp.Multiply(camera.Transform.Matrix())
  mvp = mvp.Multiply(self.Transform.Matrix())

  for _, line := range self.Lines {
    start := mvp.TransformPoint3(self.Vertices[line[0]])
    end := mvp.TransformPoint3(self.Vertices[line[1]])

    term.DrawLine(start, end, char)
  }
}


func NewMeshCube() Mesh {
  return Mesh {
    Transform: NewTransform(),
    Vertices: []Point3{
      Point3 {-1, -1, -1},
      Point3 { 1, -1, -1},
      Point3 { 1,  1, -1},
      Point3 {-1,  1, -1},

      Point3 {-1, -1,  1},
      Point3 { 1, -1,  1},
      Point3 { 1,  1,  1},
      Point3 {-1,  1,  1},
    },
    Lines: []Line{
      // Front
      Line {0, 1},
      Line {1, 2},
      Line {2, 3},
      Line {3, 0},

      // Back
      Line {4, 5},
      Line {5, 6},
      Line {6, 7},
      Line {7, 4},

      // Top
      Line {0, 4},
      Line {4, 5},
      Line {5, 1},
      Line {1, 0},

      // Bottom
      Line {2, 6},
      Line {6, 7},
      Line {7, 3},
      Line {3, 2},
    },
  }

}

func NewMeshSphere() Mesh {
  xSegments := 5.0
  ySegments := 6.0

  mesh := Mesh {
    Transform: NewTransform(),
    Vertices: []Point3{
    },
    Lines: []Line{
    },
  }

  for y := 0.0; y <= ySegments; y++ {
    lat := math.Pi * (y + 2) / ySegments
    for x := 0.0; x <= xSegments; x++ {
      lng := 2.0 * math.Pi * x / xSegments
      point := SphericalToCartesian(lng, lat)
      mesh.Vertices = append(mesh.Vertices, point)

      if len(mesh.Vertices) > 1 {
        // Horizontal line
        if x > 0 {
          mesh.Lines = append(mesh.Lines, Line { int64(len(mesh.Vertices) - 2), int64(len(mesh.Vertices) - 1) })
        }

        // Vertical line
        if y > 0 {
          mesh.Lines = append(mesh.Lines, Line { int64(len(mesh.Vertices)) - int64(xSegments) - 2, int64(len(mesh.Vertices) - 1) })
        }
      }
    }

  }

  return mesh
}
