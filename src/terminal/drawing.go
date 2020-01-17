package terminal

import (
  . "fmt"
  . "math"
)

type Position struct {
  X, Y int32
}

func CSI(format string, a ...interface{}) {
  Printf("\x1b[")
  Printf(format, a...)
}

func MoveTo(position Position) {
  CSI("%d;%dH", position.Y, position.X)
}

func Clear() {
  CSI("2J")
}

func AltScreen() {
  CSI("?1049h")
}

func MainScreen() {
  CSI("?1049l")
}

func PlotChar(position Position, char rune) {
  MoveTo(position)
  Write(string(char))
}

func PlotLine(start, end Position, char rune) {
  x0 := float64(start.X)
  y0 := float64(start.Y)
  x1 := float64(end.X)
  y1 := float64(end.Y)

  dx := Abs(x1 - x0)
  dy := Abs(y1 - y0)

  var sx, sy, err float64
  if x0 < x1 {
    sx = 1
  } else {
    sx = -1
  }
  if y0 < y1 {
    sy = 1
  } else {
    sy = -1
  }

  if dx > dy {
    err =  dx / 2.0
  } else {
    err = -dy / 2.0
  }

  for {
    PlotChar(Position { int32(x0), int32(y0) }, char)
    if x0 == x1 && y0 == y1 {
      break
    }

    e2 := err
    if e2 > -dx {
      err -= dy
      x0 += sx
    }
    if e2 < dy {
      err += dx
      y0 += sy
    }
  }

  MoveTo(start)
  Printf(string(char))
}

func Write(str string) {
  Printf(str)
}

