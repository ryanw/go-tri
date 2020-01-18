package main

import (
  "time"
  "os"
  "os/signal"
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
  transform Transform
}

type Mesh struct {
  transform Transform
  vertices []Point3
  lines []Line
}

type Line [2]int64

func (self Mesh) Draw(term *Terminal, camera Camera, char rune) {
  mvp := camera.projection.Multiply(camera.transform.Matrix()).Multiply(self.transform.Matrix())

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
  term.HideCursor()
  term.Clear()

  camera := Camera {
    projection: NewMatrix4Perspective(Width / Height, 45, 0.1, 1000.0),
    transform: NewTransform(),
  }

  cube := Mesh {
    transform: Transform {
      Translation: Vector3 { 0, 0, -5 },
      Rotation: Vector3 { 0, 0, 0 },
      Scaling: Vector3 { 1, 1, 1 },
    },
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

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  go func(){
    // Catch Ctrl+C and cleanup the terminal
    for _ = range c {
      term.MainScreen()
      term.ShowCursor()
      term.Flush()
      os.Exit(0)
    }
  }()


  // Main loop
  t := 1.0
  for {
    dt := 1.0 / float64(framerate)

    term.Draw(func() {

      // Clear old cube
      cube.Draw(&term, camera, ' ')

      // Move the cube
      cube.transform.Rotation[0] += 0.25 * m.Pi * dt
      cube.transform.Rotation[1] += 0.5 * m.Pi * dt
      cube.transform.Translation[0] = m.Cos(t * 2)
      cube.transform.Translation[2] = -8 - m.Sin(t * 1.2) * 5

      // Draw new cube
      cube.Draw(&term, camera, '.')
    })

    time.Sleep((1000 / framerate) * time.Millisecond)
    t += dt
  }
}

