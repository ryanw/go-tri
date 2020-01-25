package terminal

import (
	. "../geom"
	. "math"
)

type Position struct {
	X, Y int
}

func (t *Terminal) CSI(format string, a ...interface{}) {
	t.Write("\x1b[")
	t.Write(format, a...)
}

func (t *Terminal) Left(amount int) {
	if amount == 1 {
		t.CSI("D")
	} else if amount > 1 {
		t.CSI("%dD", amount)
	} else if amount < 0 {
		t.Right(-amount)
	}
}

func (t *Terminal) Right(amount int) {
	if amount == 1 {
		t.CSI("C")
	} else if amount > 1 {
		t.CSI("%dC", amount)
	} else if amount < 0 {
		t.Left(-amount)
	}
}

func (t *Terminal) Down(amount int) {
	if amount == 1 {
		t.CSI("B")
	} else if amount > 1 {
		t.CSI("%dB", amount)
	} else if amount < 0 {
		t.Up(-amount)
	}
}

func (t *Terminal) Up(amount int) {
	if amount == 1 {
		t.CSI("A")
	} else if amount > 1 {
		t.CSI("%dA", amount)
	} else if amount < 0 {
		t.Down(-amount)
	}
}

func (t *Terminal) MoveTo(position Position) {
	t.CSI("%d;%dH", position.Y+1, position.X+1)
}

func (t *Terminal) Clear() {
	t.CSI("2J")
}

func (t *Terminal) AltScreen() {
	t.CSI("?1049h")
}

func (t *Terminal) MainScreen() {
	t.CSI("?1049l")
}

func (t *Terminal) ShowCursor() {
	t.CSI("?25h")
}

func (t *Terminal) HideCursor() {
	t.CSI("?25l")
}

func (t *Terminal) Color(code uint8) {
	t.CSI("%dm", int(code)+30)
}

func (t *Terminal) PlotChar(position Position, char rune) {
	if position.X < 0 || position.Y < 0 || position.X >= t.width || position.Y >= t.height {
		return
	}
	t.MoveTo(position)
	t.Write(string(char))
}

// Draw vector line using coordinates between -1.0 and +1.0
// It scales to the screen size
func (t *Terminal) DrawLine(start, end Point3, char rune) {
	// Center and scale coordinates
	hw := float64(t.width) / 2
	hh := float64(t.height) / 2
	startPos := Position{
		int(start[0]*hw + hw),
		int(start[1]*hh + hh),
	}
	endPos := Position{
		int(end[0]*hw + hw),
		int(end[1]*hh + hh),
	}
	t.PlotLine(startPos, endPos, char)
}

func (t *Terminal) ClearLine(start, end Position) {
	t.PlotLine(start, end, ' ')
}

// Draw a line using absolute character positions
func (t *Terminal) PlotLine(start, end Position, char rune) {
	x0 := float64(start.X)
	y0 := float64(start.Y)
	x1 := float64(end.X)
	y1 := float64(end.Y)
	px := int(x0)
	py := int(y0)

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
		err = dx / 2.0
	} else {
		err = -dy / 2.0
	}

	// Move to start of line
	t.MoveTo(Position{int(x0), int(y0)})
	color := uint8(0)
	for {
		symbol := char
		if symbol != ' ' {
			// Draw dots at the ends
			if (int(x0) == start.X && int(y0) == start.Y) || (int(x0) == end.X && int(y0) == end.Y) {
				color = 5
				t.Color(color)
				symbol = '♦'
			} else if color == 5 {
				color = 3
				t.Color(color)
			}
		}
		if x0 >= 0 && y0 >= 0 && int(x0) < t.width && int(y0) < t.height {
			if px < 0 || py < 0 || int(px) >= t.width || int(py) >= t.height {
				// Prev location was off screen, so we have to move the cursor
				t.MoveTo(Position{int(x0), int(y0)})
			}
			t.WriteRune(symbol)
		}

		if x0 == x1 && y0 == y1 {
			// The end
			break
		}

		px = int(x0)
		py = int(y0)

		e2 := err
		if e2 > -dx {
			err -= dy
			x0 += sx
			t.Right(int(x0) - px - 1)
		} else {
			t.Left(1)
		}
		if e2 < dy {
			err += dx
			y0 += sy
			t.Down(int(y0) - py)
		}
	}
}
