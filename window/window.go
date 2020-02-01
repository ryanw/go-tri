package window

import (
	. "tri/canvas"
	. "tri/geom"
	. "tri/mesh"
	. "tri/renderer"
	. "tri/terminal"
)

type Window struct {
	Terminal Terminal
	Canvas   Canvas
	Renderer Renderer
}

func New() Window {
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

	return Window{
		Terminal: term,
		Canvas:   canvas,
		Renderer: renderer,
	}
}

func (w *Window) Open() {
	w.Terminal.AltScreen()
	w.Terminal.HideCursor()
	w.Terminal.RawMode()
	w.Terminal.EnableMouse()
	//w.Terminal.DisableCtrlC()
	w.Terminal.Clear()
	w.Terminal.Flush()
}

func (w *Window) Close() {
	w.Terminal.ShowCursor()
	w.Terminal.NormalMode()
	w.Terminal.DisableMouse()
	w.Terminal.EnableCtrlC()
	w.Terminal.MainScreen()
	w.Terminal.Flush()
}

func (w *Window) Draw(dt float64) {
	w.Canvas.Clear()
}

func (w *Window) Present() {
	w.Canvas.Present(&w.Terminal)
}

func (w *Window) Clear() {
	w.Canvas.Clear()
}

func (w *Window) DrawMesh(mesh *TriangleMesh) {
	w.Renderer.RenderTriangleMesh(&w.Canvas, mesh)
}
