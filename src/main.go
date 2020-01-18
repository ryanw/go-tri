package main

import (
  "time"
  m "math"
  . "./terminal"
  . "./geom"
)

const framerate = 30
// FIXME calculate from terminal size
const Width = 80.0
const Height = 48.0


type Camera struct {
  projection Matrix4
  transform Matrix4
}

type Mesh struct {
  transform Matrix4
  vertices []Point3
  lines []Line
}

type Line [2]int64

func (self Mesh) Draw(term *Terminal, camera Camera, char rune) {
  mvp := camera.projection.Multiply(camera.transform).Multiply(self.transform)

  hw := Width / 2.0
  hh := Height / 2.0

  for _, line := range self.lines {
    start := mvp.TransformPoint3(self.vertices[line[0]])
    end := mvp.TransformPoint3(self.vertices[line[1]])

    x1 := start[0]
    y1 := start[1]
    x2 := end[0]
    y2 := end[1]

    // Center and scale to display size
    x1 *= hw
    x2 *= hw
    y1 *= hh
    y2 *= hh

    x1 += hw
    x2 += hw
    y1 += hh
    y2 += hh

    term.PlotLine(
      Position { int(x1), int(y1) },
      Position { int(x2), int(y2) },
      char,
    )
  }
}

func main() {
  term := NewTerminal(Width, Height)

  term.AltScreen()
  term.Clear()

  camera := Camera {
    projection: NewMatrix4Perspective(Width / Height, 45, 0.1, 1000.0),
    transform: NewMatrix4Identity(),
  }

  cube := Mesh {
    transform: NewMatrix4Translation(0.0, 0.0, -5.0),
    vertices: []Point3{
      Point3 {-1, -1, -1},
      Point3 { 1, -1, -1},
      Point3 { 1,  1, -1},
      Point3 {-1,  1, -1},

      Point3 {-1, -1,  1},
      Point3 { 1, -1,  1},
      Point3 { 1,  1,  1},
      Point3 {-1,  1,  1},
    },
    lines: []Line{
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

  t := 1.0
  for {
    dt := 1.0 / float64(framerate)

    term.Draw(func() {

      // Clear old cube
      cube.Draw(&term, camera, ' ')

      // Move the cube
      cube.transform = cube.transform.Multiply(
        NewMatrix4Rotation(0.25 * m.Pi * dt, 0.5 * m.Pi * dt, 0.0),
      )

      // Draw new cube
      cube.Draw(&term, camera, '.')
    })

    time.Sleep((1000 / framerate) * time.Millisecond)
    t += dt
  }

  term.MainScreen()
}

