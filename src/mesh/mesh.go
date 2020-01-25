package mesh

import (
  "math"
  . "../geom"
  . "../terminal"
)

type Line [2]int

type LineMesh struct {
  Transform Transform
  Vertices []Point3
  Lines []Line
}

type TriangleMesh struct {
  Transform Transform
  Vertices []Point3
  Triangles [][3]int
  Normals []Vector3
  Colors []uint32
}

func (self *LineMesh) Draw(term *Terminal, camera Camera, char rune) {
  mvp := camera.Projection
  mvp = mvp.Multiply(camera.Transform.Matrix())
  mvp = mvp.Multiply(self.Transform.Matrix())

  for _, line := range self.Lines {
    start := mvp.TransformPoint3(self.Vertices[line[0]])
    end := mvp.TransformPoint3(self.Vertices[line[1]])

    term.DrawLine(start, end, char)
  }
}

func NewTriangleMeshCube() TriangleMesh {
  return TriangleMesh {
    Transform: NewTransform(),
    Vertices: []Point3{
      Point3 {-1, -1,  1},
      Point3 { 1, -1,  1},
      Point3 { 1,  1,  1},
      Point3 {-1,  1,  1},

      Point3 {-1, -1, -1},
      Point3 { 1, -1, -1},
      Point3 { 1,  1, -1},
      Point3 {-1,  1, -1},
    },
    Triangles: [][3]int{
      // Front
      [3]int { 0, 1, 2 },
      [3]int { 2, 0, 3 },

      // Back
      [3]int { 4, 5, 6 },
      [3]int { 6, 4, 7 },

      // Left
      [3]int { 0, 3, 7 },
      [3]int { 0, 7, 4 },

      // Right
      [3]int { 1, 2, 6 },
      [3]int { 1, 6, 5 },
    },
    Normals: []Vector3 {
      Vector3 {  0, 0,  1 },
      Vector3 {  0, 0,  1 },

      Vector3 {  0, 0, -1 },
      Vector3 {  0, 0, -1 },

      Vector3 { -1, 0,  0 },
      Vector3 { -1, 0,  0 },

      Vector3 {  1, 0,  0 },
      Vector3 {  1, 0,  0 },
    },
    Colors: []uint32 {
      0xffff0000,
      0xffff0000,

      0xff00ff00,
      0xff00ff00,

      0xff0000ff,
      0xff0000ff,

      0xffffff00,
      0xffffff00,
    },
  }

}

func NewLineMeshCube() LineMesh {
  return LineMesh {
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

func NewLineMeshSphere() LineMesh {
  xSegments := 8.0
  ySegments := 6.0

  mesh := LineMesh {
    Transform: NewTransform(),
    Vertices: []Point3{
    },
    Lines: []Line{
    },
  }

  for y := ySegments * -0.5; y <= ySegments * 0.5; y++ {
    lat := math.Pi * (y / ySegments)

    for x := xSegments * -0.5; x <= xSegments * 0.5; x++ {
      lng := 2.0 * math.Pi * x / xSegments

      point := SphericalToCartesian(lng, lat)
      mesh.Vertices = append(mesh.Vertices, point)

      if len(mesh.Vertices) > 1 {
        // Horizontal line
        if x > xSegments * -0.5 {
          mesh.Lines = append(mesh.Lines, Line { int(len(mesh.Vertices) - 2), int(len(mesh.Vertices) - 1) })
        }

        // Vertical line
        if y > ySegments * -0.5 {
          mesh.Lines = append(mesh.Lines, Line { int(len(mesh.Vertices)) - int(xSegments) - 2, int(len(mesh.Vertices) - 1) })
        }
      }
    }

  }

  return mesh
}
