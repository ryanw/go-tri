package main

import (
  "time"
  "os"
  "os/signal"
  "syscall"
  m "math"
  . "../canvas"
  . "../terminal"
  . "../geom"
  . "../mesh"
  . "../renderer"
)

const framerate = 30

func main() {
  term := NewTerminal()
  width, height := term.Size()
  canvas := NewCanvas(width, height)
  renderer := Renderer {
    Camera: Camera {
      Projection: NewMatrix4Perspective(float64(width) / float64(height), 45, 0.1, 1000.0),
      Transform: Transform {
        Translation: Vector3 { 0, 0, 0 },
        Rotation: Vector3 { 0, 0, 0 },
        Scaling: Vector3 { 1, 1, 1 },
      },
    },
  }



  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  signal.Notify(c, syscall.SIGWINCH)
  go func() {
    for sig := range c {
      switch sig {
      // User pressed Ctrl+C
      case os.Interrupt:
        term.MainScreen()
        term.ShowCursor()
        term.Flush()
        os.Exit(0)

      // Terminal was resized
      case syscall.SIGWINCH:
        term.UpdateSize()
        width, height = term.Size()
        canvas.Resize(width, height)
        renderer.Camera.Projection = NewMatrix4Perspective(float64(width) / float64(height), 45, 0.1, 1000.0)
      }
    }
  }()

  cube := NewLineMeshCube()
  cube.Transform = Transform {
    Translation: Vector3 { 2, 0, -7 },
    Rotation: Vector3 { 0, 0, 0 },
    Scaling: Vector3 { 1, 1, 1 },
  }

  triCube := NewTriangleMeshCube()
  triCube.Transform = Transform {
    Translation: Vector3 { -2, 0, -7 },
    Rotation: Vector3 { 0, 0, 0 },
    Scaling: Vector3 { 1, 1, 1 },
  }

  sphere := NewLineMeshSphere()
  sphere.Transform = Transform {
    Translation: Vector3 { -0.5, 0, -4 },
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

    cube.Transform.Rotation[0] += 0.15 * m.Pi * dt
    cube.Transform.Rotation[1] += 0.5 * m.Pi * dt
    cube.Transform.Translation[2] = -8 - m.Sin(t * 0.8) * 2
    triCube.Transform.Rotation[0] += 0.15 * m.Pi * dt
    triCube.Transform.Rotation[1] += 0.5 * m.Pi * dt
    triCube.Transform.Translation[2] = -8 - m.Sin(t * 0.8) * 2

    sphere.Transform.Rotation[0] += 0.25 * m.Pi * dt
    sphere.Transform.Rotation[1] += 0.5 * m.Pi * dt
    sphere.Transform.Translation[2] = -4 - m.Sin(t * 1.2) * 2

    canvas.Clear()
    renderer.RenderLineMesh(&canvas, &cube)
    renderer.RenderTriangleMesh(&canvas, &triCube)
    canvas.Present(&term)

    time.Sleep((1000 / framerate) * time.Millisecond)
    t += dt
  }
}

