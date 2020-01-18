package main

import (
  "time"
  "os"
  "os/signal"
  "syscall"
  m "math"
  . "./terminal"
  . "./geom"
  . "./mesh"
)

const framerate = 30

func main() {
  term := NewTerminal()

  width, height := term.Size()

  term.AltScreen()
  term.HideCursor()
  term.Clear()

  camera := Camera {
    Projection: NewMatrix4Perspective(float64(width) / float64(height), 45, 0.1, 1000.0),
    Transform: NewTransform(),
  }

  cube := Mesh {
    Transform: Transform {
      Translation: Vector3 { 0, 0, -5 },
      Rotation: Vector3 { 0, 0, 0 },
      Scaling: Vector3 { 1, 1, 1 },
    },
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

  // Listen for resize
  cr := make(chan os.Signal, 1)
  signal.Notify(cr, syscall.SIGWINCH)
  go func(){
    for _ = range cr {
      term.Clear()
      term.UpdateSize()
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
      cube.Transform.Rotation[0] += 0.25 * m.Pi * dt
      cube.Transform.Rotation[1] += 0.5 * m.Pi * dt
      cube.Transform.Translation[0] = m.Cos(t * 2)
      cube.Transform.Translation[2] = -8 - m.Sin(t * 1.2) * 5

      // Draw new cube
      cube.Draw(&term, camera, '.')
    })

    time.Sleep((1000 / framerate) * time.Millisecond)
    t += dt
  }
}

