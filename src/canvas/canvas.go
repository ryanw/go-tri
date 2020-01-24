package canvas
import (
  "fmt"
  . "math"
  . "../terminal"
  . "../geom"
)

type Color uint32

func (c *Color) ToRgb() (uint8, uint8, uint8) {
  return 0, 0, 0
}

func (c *Color) ToRgba() (uint8, uint8, uint8, uint8) {
  return 0, 0, 0, 0
}

func (c Color) ToAnsi() uint16 {
  //a := ((c & 0xff000000) >> 24) / 51;
  r := ((c & 0x00ff0000) >> 16) / 51;
  g := ((c & 0x0000ff00) >> 8) / 51;
  b := ((c & 0x000000ff) >> 0) / 51;
  return uint16(16 + 36 * r + 6 * g + b)
}

type Cell struct {
  Fg Color
  Bg Color
  Sprite rune
}

func (c *Cell) AnsiColor() string {
  return fmt.Sprintf("\x1b[38;5;%dm\x1b[48;5;%dm", c.Fg.ToAnsi(), c.Bg.ToAnsi())
}

type Canvas struct {
  Width int
  Height int
  front []Cell
  back []Cell
}

func NewCanvas(width, height int) Canvas {
  return Canvas {
    Width: width,
    Height: height,
    front: make([]Cell, width * height),
    back: make([]Cell, width * height),
  }
}

func (c *Canvas) IsOutOfBounds(x, y int) bool {
  return x < 0 || y < 0 || x >= c.Width || y >= c.Height
}

func (c *Canvas) Set(x, y int, cell Cell) {
  if c.IsOutOfBounds(x, y) {
    return
  }
  idx := c.positionToIndex(x, y)
  c.back[idx] = cell
}

func (c *Canvas) Get(x, y int) *Cell {
  if c.IsOutOfBounds(x, y) {
    return nil
  }
  return c.GetBack(x, y)
}

func (c *Canvas) GetFront(x, y int) *Cell {
  if c.IsOutOfBounds(x, y) {
    return nil
  }
  idx := c.positionToIndex(x, y)
  return &c.front[idx]
}

func (c *Canvas) GetBack(x, y int) *Cell {
  if c.IsOutOfBounds(x, y) {
    return nil
  }
  idx := c.positionToIndex(x, y)
  return &c.back[idx]
}

func (c *Canvas) Resize(width, height int) {
  c.Width = width
  c.Height = height
  // FIXME Don't wipe out what's already drawn
  c.front = make([]Cell, width * height)
  c.back = make([]Cell, width * height)
}

func (c *Canvas) Clear() {
  c.ClearWithCell(Cell {
    Fg: 0xffffffff, // White
    Bg: 0x00000000, // Transparent
    Sprite: ' ',
  })
}

func (c *Canvas) ClearWithCell(cell Cell) {
  for i := range c.back {
    c.back[i] = cell
  }
}

func (c *Canvas) DrawLine3D(start, end Point3, cell Cell) {
  // Center and scale coordinates
  hw := float64(c.Width) / 2
  hh := float64(c.Height) / 2
  startPos := [2]int{
    int(start[0] * hw + hw),
    int(start[1] * hh + hh),
  }
  endPos := [2]int {
    int(end[0] * hw + hw),
    int(end[1] * hh + hh),
  }
  c.DrawLine(startPos, endPos, cell)
}

func (c *Canvas) DrawLine(start, end [2]int, cell Cell) {
  x0 := float64(start[0])
  y0 := float64(start[1])
  x1 := float64(end[0])
  y1 := float64(end[1])

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
    c.Set(int(x0), int(y0), cell)

    if x0 == x1 && y0 == y1 {
      // The end
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
}

func (c *Canvas) Present(term *Terminal) {
  width, height := c.Width, c.Height
  if term.Width() < width {
    width = term.Width()
  }
  if term.Height() < height {
    height = term.Height()
  }

  cursorX, cursorY := -1, -1
  var cursorFg, cursorBg Color = 0x0, 0x0
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      backCell := c.GetBack(x, y)
      frontCell := c.GetFront(x, y)

      if backCell == frontCell {
        continue
      }

      if x != cursorX || y != cursorY {
        cursorX = x
        cursorY = y
        term.MoveTo(Position { x, y })
      }

      if cursorFg != backCell.Fg || cursorBg != backCell.Bg {
        cursorFg = backCell.Fg
        cursorBg = backCell.Bg
        term.Write(backCell.AnsiColor())
      }

      if backCell.Sprite != frontCell.Sprite {
        frontCell.Sprite = backCell.Sprite
        term.WriteRune(frontCell.Sprite)
        cursorX = x + 1
      }
    }
  }
  term.Flush()
}

func (c *Canvas) positionToIndex(x, y int) int {
  return x + y * c.Width
}

func (c *Canvas) indexToPosition(idx int) (int, int) {
  x := idx % c.Width
  y := idx / c.Width
  return x, y
}
