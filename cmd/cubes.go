package main

import (
	m "math"
	"os"
	"os/signal"
	"syscall"
	"time"
	. "tri/canvas"
	. "tri/geom"
	. "tri/mesh"
	. "tri/renderer"
	. "tri/terminal"
)

const framerate = 30

func main() {
	term := NewTerminal()
	width, height := term.Size()
	canvas := NewCanvas(width, height)
	renderer := Renderer{
		Camera: Camera{
			Projection: NewMatrix4Perspective(float64(width)/float64(height), 45, 0.1, 1000.0),
			Transform: Transform{
				Translation: Vector3{0, 0, 0},
				Rotation:    Vector3{0, 0, 0},
				Scaling:     Vector3{1, 1, 1},
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
				term.NormalMode()
				term.DisableMouse()
				term.Flush()
				os.Exit(0)

			// Terminal was resized
			case syscall.SIGWINCH:
				term.UpdateSize()
				width, height = term.Size()
				canvas.Resize(width, height)
				renderer.Camera.Projection = NewMatrix4Perspective(float64(width)/float64(height), 45, 0.1, 1000.0)
			}
		}
	}()

	cube := NewLineMeshCube()
	cube.Transform = Transform{
		Translation: Vector3{2, 0, -7},
		Rotation:    Vector3{0, 0, 0},
		Scaling:     Vector3{1, 1, 1},
	}

	triCube := NewTriangleMeshCube()
	triCube.Transform = Transform{
		Translation: Vector3{-2, 0, -7},
		Rotation:    Vector3{0, 0, 0},
		Scaling:     Vector3{1, 1, 1},
	}

	plane := NewTriangleMeshPlane(4, 4)
	plane.Transform = Transform{
		Translation: Vector3{0, 2, -10},
		Rotation:    Vector3{0, 0, 0},
		Scaling:     Vector3{3, 1, 3},
	}

	term.AltScreen()
	term.HideCursor()
	term.RawMode()
	term.EnableMouse()
	term.Clear()

	// User input events
	go func() {
		for {
			event := term.NextEvent()
			switch event.EventType {
			case KeyEvent:

			case MouseEvent:
				switch event.MouseAction {
				case MouseMove:
					x := float64(event.MouseX)
					y := float64(event.MouseY)
					vx := x / float64(width)
					vy := y / float64(height)
					triCube.Transform.Rotation[0] = m.Pi * vy * 2
					triCube.Transform.Rotation[1] = m.Pi * -vx * 2

				case MouseDown:
				}
			}
		}
	}()

	// Main loop
	t := 1.0
	for {
		dt := 1.0 / float64(framerate)

		f := 0.5 * m.Pi

		cube.Transform.Rotation[0] += f * dt
		cube.Transform.Rotation[1] += f * dt
		cube.Transform.Translation[2] = -8 - m.Sin(t*0.8)*2
		//triCube.Transform.Rotation[0] += f * dt
		//triCube.Transform.Rotation[1] += f * dt
		triCube.Transform.Translation[2] = -8 - m.Sin(t*0.8)*2
		plane.Transform.Rotation[1] -= f * dt

		canvas.Clear()
		renderer.RenderLineMesh(&canvas, &cube)
		renderer.RenderTriangleMesh(&canvas, &triCube)
		renderer.RenderTriangleMesh(&canvas, &plane)
		canvas.Present(&term)

		time.Sleep((1000 / framerate) * time.Millisecond)
		t += dt
	}
}
