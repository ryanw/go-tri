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

  // Catch Ctrl+C and cleanup the terminal
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  go func(){
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


  camera := Camera {
    Projection: NewMatrix4Perspective(float64(width) / float64(height), 45, 0.1, 1000.0),
    Transform: NewTransform(),
  }

  cube := NewMeshCube()
  cube.Transform = Transform {
    Translation: Vector3 { 2.5, 0, -7 },
    Rotation: Vector3 { 0, 0, 0 },
    Scaling: Vector3 { 1, 1, 1 },
  }

  sphere := NewMeshSphere()
  sphere.Transform = Transform {
    Translation: Vector3 { -1, 0, -4 },
    Rotation: Vector3 { 0, 0, 0 },
    Scaling: Vector3 { 1, 1, 1 },
  }


  term.AltScreen()
  term.HideCursor()
  term.Clear()

  // Main loop
  t := 1.0
  for {
    dt := 1.0 / float64(framerate)

    term.Draw(func() {
      cube.Draw(&term, camera, ' ')
      sphere.Draw(&term, camera, ' ')

      cube.Transform.Rotation[0] += 0.25 * m.Pi * dt
      cube.Transform.Rotation[1] += 0.5 * m.Pi * dt

      sphere.Transform.Rotation[0] += 0.25 * m.Pi * dt
      sphere.Transform.Rotation[1] += 0.5 * m.Pi * dt

      cube.Draw(&term, camera, '.')
      sphere.Draw(&term, camera, '.')
    })

    time.Sleep((1000 / framerate) * time.Millisecond)
    t += dt
  }
}

