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
	. "tri/scene"
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
				Translation: Vector3{0, -12, 10},
				Rotation:    Vector3{0.25, 0, 0},
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

	term.AltScreen()
	term.HideCursor()
	term.RawMode()
	term.EnableMouse()
	term.Clear()

	mouseX, mouseY := -1.0, -1.0

	// User input events
	go func() {
		for {
			event := term.NextEvent()
			switch event.EventType {
			case KeyEvent:
				velocity := 1.5
				switch event.Key {
				case 'c':
					term.Clear()
				case 'w':
					renderer.Camera.Translate(0, 0, -velocity)
				case 's':
					renderer.Camera.Translate(0, 0, velocity)
				case 'a':
					renderer.Camera.Translate(-velocity, 0, 0)
				case 'd':
					renderer.Camera.Translate(velocity, 0, 0)
				case 'e':
					renderer.Camera.Translate(0, -velocity, 0)
				case 'q':
					renderer.Camera.Translate(0, velocity, 0)
				case ',':
					renderer.Camera.Transform.Rotation[1] += 0.01 * m.Pi
				case '.':
					renderer.Camera.Transform.Rotation[1] -= 0.01 * m.Pi
				case 'z':
					renderer.Camera.Transform.Rotation[0] += 0.01 * m.Pi
				case 'x':
					renderer.Camera.Transform.Rotation[0] -= 0.01 * m.Pi
				case '\r', '\n':
					scaleX := &renderer.Camera.Transform.Scaling[0]
					if *scaleX == 0.5 {
						*scaleX = 1.0
					} else {
						*scaleX = 0.5
					}
				}

			case MouseEvent:
				switch event.MouseAction {
				case MouseMove:
					x := float64(event.MouseX)
					y := float64(event.MouseY)
					vx := x / float64(width)
					vy := y / float64(height)
					if mouseX > -1.0 {
						renderer.Camera.Transform.Rotation[1] += (mouseX - vx) * m.Pi
					}
					if mouseY > -1.0 {
						renderer.Camera.Transform.Rotation[0] -= (mouseY - vy) * m.Pi
					}
					mouseX = vx
					mouseY = vy

				case MouseDown:
					x := float64(event.MouseX)
					y := float64(event.MouseY)
					mouseX = x / float64(width)
					mouseY = y / float64(height)

				case MouseUp:
					mouseX, mouseY = -1.0, -1.0
				}
			}
		}
	}()

	// Create a scene
	cube := NewTriangleMeshCube()
	cube.Transform = Transform{
		Translation: Vector3{0, -10, 0},
		Rotation:    Vector3{0, 0, 0},
		Scaling:     Vector3{1, 1, 1},
	}

	scene := NewScene()
	cubeId := scene.Add(cube)

	chunkSize := 8
	go func() {
		for x := -2; x < 2; x++ {
			for y := -2; y < 2; y++ {
				time.Sleep(100 * time.Millisecond)
				terrain := NewTerrainMesh(x*chunkSize, y*chunkSize, chunkSize, chunkSize, 0.1)
				terrain.Transform = Transform{
					Translation: Vector3{float64(x * chunkSize), 0, float64(y * chunkSize)},
					Rotation:    Vector3{0, 0, 0},
					Scaling:     Vector3{1, 5, 1},
				}
				scene.Add(terrain)
			}
		}
	}()

	// Main loop
	t := 1.0
	for {
		dt := 1.0 / float64(framerate)

		f := 0.5 * m.Pi

		cubeMesh := scene.Mesh(cubeId)
		cubeMesh.Transform.Rotation[0] += f * dt
		cubeMesh.Transform.Rotation[1] += f * dt
		if len(scene.Meshes) > 7 {
			terrainMesh := scene.Mesh(6)
			terrainMesh.Transform.Translation[1] = 2 * -m.Abs(m.Sin(3*t))
		}

		canvas.Lock()
		canvas.ClearWithCell(Cell{
			Fg:     0x0,
			Bg:     0xff442c7d,
			Depth:  1000000,
			Sprite: ' ',
		})
		renderer.RenderDrawable(&canvas, &scene)
		canvas.Present(&term)
		canvas.Unlock()

		time.Sleep((1000 / framerate) * time.Millisecond)
		t += dt
	}
}
