package canvas
import (
  "fmt"
  "sort"
  . "math"
  . "../terminal"
  . "../geom"
)

type Color uint32
type Triangle [3][2]int;
type TriangleFloat [3][2]float64;

func (t *Triangle) ToFloat() TriangleFloat {
  return TriangleFloat {
    [2]float64 {
      float64(t[0][0]),
      float64(t[0][1]),
    },
    [2]float64 {
      float64(t[1][0]),
      float64(t[1][1]),
    },
    [2]float64 {
      float64(t[2][0]),
      float64(t[2][1]),
    },
  }
}

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
  startF := [2]float64 { float64(start[0]), float64(start[1]) }
  endF := [2]float64 { float64(end[0]), float64(end[1]) }
  c.DrawLineFloat(startF, endF, cell)
}

func (c *Canvas) DrawLineFloat(start, end [2]float64, cell Cell) {
  x0 := start[0]
  y0 := start[1]
  x1 := end[0]
  y1 := end[1]

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

func (c *Canvas) fillFlatBottomTriangle(tri TriangleFloat, cell Cell) {
  slope0 := (tri[1][0] - tri[0][0]) / (tri[1][1] - tri[0][1])
  slope1 := (tri[2][0] - tri[0][0]) / (tri[2][1] - tri[0][1])

  x0 := tri[0][0]
  x1 := tri[0][0]

  for y := tri[0][1]; y <= tri[1][1]; y++ {
    c.DrawLine([2]int { int(x0), int(y) }, [2]int { int(x1) , int(y) }, cell)
    x0 += slope0
    x1 += slope1
  }
}

func (c *Canvas) fillFlatTopTriangle(tri TriangleFloat, cell Cell) {
  slope0 := (tri[2][0] - tri[0][0]) / (tri[2][1] - tri[0][1])
  slope1 := (tri[2][0] - tri[1][0]) / (tri[2][1] - tri[1][1])

  x0 := tri[2][0]
  x1 := tri[2][0]

  for y := tri[2][1]; y > tri[0][1]; y-- {
    c.DrawLine([2]int { int(x0), int(y) }, [2]int { int(x1) , int(y) }, cell)
    x0 -= slope0
    x1 -= slope1
  }
}

func (c *Canvas) DrawVectorTriangle(tri TriangleFloat, cell Cell) {
  // Center and scale coordinates
  hw := float64(c.Width) / 2
  hh := float64(c.Height) / 2
  c.DrawTriangle(
    Triangle {
      [2]int{
        int(tri[0][0] * hw + hw),
        int(tri[0][1] * hh + hh),
      },
      [2]int{
        int(tri[1][0] * hw + hw),
        int(tri[1][1] * hh + hh),
      },
      [2]int{
        int(tri[2][0] * hw + hw),
        int(tri[2][1] * hh + hh),
      },
    },

    cell,
  )
}

func (c *Canvas) DrawTriangle(tri Triangle, cell Cell) {
  // Sort by Y axis
  sort.Slice(tri[:], func(i, j int) bool {
    return tri[i][1] < tri[j][1]
  })

  floatTri := tri.ToFloat();

  if floatTri[1][1] == floatTri[2][1] {
    c.fillFlatBottomTriangle(floatTri, cell)

  } else if floatTri[0][1] == floatTri[1][1] {
    c.fillFlatTopTriangle(floatTri, cell)

  } else {
    midVert := [2]float64 {
      floatTri[0][0] + (floatTri[1][1] - floatTri[0][1]) / (floatTri[2][1] - floatTri[0][1]) * (floatTri[2][0] - floatTri[0][0]),
      floatTri[1][1],
    }

    c.fillFlatTopTriangle(TriangleFloat { floatTri[1], midVert, floatTri[2] }, cell)
    c.fillFlatBottomTriangle(TriangleFloat { floatTri[0], floatTri[1], midVert }, cell)
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
        // Write pixel colour
        term.Write(backCell.AnsiColor())
      }

      if frontCell.Fg != backCell.Fg || frontCell.Bg != backCell.Bg || backCell.Sprite != frontCell.Sprite {
        *frontCell = *backCell
        // Write pixel ascii
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
