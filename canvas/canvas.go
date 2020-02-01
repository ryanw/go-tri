package canvas

import (
	"fmt"
	. "math"
	"sort"
	. "tri/geom"
	. "tri/terminal"
)

type Triangle [3][2]int
type TriangleFloat [3][2]float64

func (t *Triangle) ToFloat() TriangleFloat {
	return TriangleFloat{
		[2]float64{
			float64(t[0][0]),
			float64(t[0][1]),
		},
		[2]float64{
			float64(t[1][0]),
			float64(t[1][1]),
		},
		[2]float64{
			float64(t[2][0]),
			float64(t[2][1]),
		},
	}
}

type Cell struct {
	Fg     Color
	Bg     Color
	Depth  float64
	Sprite rune
}

func (c *Cell) AnsiColor() string {
	return fmt.Sprintf("\x1b[38;5;%dm\x1b[48;5;%dm", c.Fg.ToAnsi(), c.Bg.ToAnsi())
}

func (c *Cell) Ansi24BitColor() string {
	fr := ((c.Fg & 0x00ff0000) >> 16)
	fg := ((c.Fg & 0x0000ff00) >> 8)
	fb := ((c.Fg & 0x000000ff) >> 0)
	br := ((c.Bg & 0x00ff0000) >> 16)
	bg := ((c.Bg & 0x0000ff00) >> 8)
	bb := ((c.Bg & 0x000000ff) >> 0)

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm", fr, fg, fb, br, bg, bb)
}

type Canvas struct {
	Width  int
	Height int
	front  []Cell
	back   []Cell
}

func NewCanvas(width, height int) Canvas {
	return Canvas{
		Width:  width,
		Height: height,
		front:  make([]Cell, width*height),
		back:   make([]Cell, width*height),
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

func (c *Canvas) DepthAt(x, y int) float64 {
	if c.IsOutOfBounds(x, y) {
		return 0.0
	}
	return c.GetBack(x, y).Depth
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
	if idx >= len(c.front) {
		return nil
	}
	return &c.front[idx]
}

func (c *Canvas) GetBack(x, y int) *Cell {
	if c.IsOutOfBounds(x, y) {
		return nil
	}
	idx := c.positionToIndex(x, y)
	if idx >= len(c.back) {
		return nil
	}
	return &c.back[idx]
}

func (c *Canvas) Resize(width, height int) {
	c.Width = width
	c.Height = height
	// FIXME Don't wipe out what's already drawn
	c.front = make([]Cell, width*height)
	c.back = make([]Cell, width*height)
}

func (c *Canvas) Clear() {
	c.ClearWithCell(Cell{
		Fg:     0xffffffff, // White
		Bg:     0x00000000, // Transparent
		Depth:  1000000,
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
		int(start[0]*hw + hw),
		int(start[1]*hh + hh),
	}
	endPos := [2]int{
		int(end[0]*hw + hw),
		int(end[1]*hh + hh),
	}
	c.DrawLine(startPos, endPos, cell)
}

// FIXME coords aren't consistent with naming
// Line coords are cell positions
func (c *Canvas) DrawDeepLine(line Line3, cell Cell) {
	start, end := line[0], line[1]

	x0, y0, z0 := Floor(start[0]), Floor(start[1]), start[2]
	x1, y1, z1 := Floor(end[0]), Floor(end[1]), end[2]

	dx := Abs(x1 - x0)
	dy := Abs(y1 - y0)
	dz := 0.0

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
		dz = (z0 - z1) / dx
		err = dx / 2.0
	} else {
		dz = (z0 - z1) / dy
		err = -dy / 2.0
	}

	for {
		depth := z0
		if c.DepthAt(int(x0), int(y0)) > depth {
			cell.Depth = depth
			c.Set(int(x0), int(y0), cell)
		}

		if int(x0) == int(x1) && int(y0) == int(y1) {
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
		z0 -= dz
	}
}

func (c *Canvas) DrawLine(start, end [2]int, cell Cell) {
	startF := [2]float64{float64(start[0]), float64(start[1])}
	endF := [2]float64{float64(end[0]), float64(end[1])}
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
		err = dx / 2.0
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
		c.DrawLine([2]int{int(x0), int(y)}, [2]int{int(x1), int(y)}, cell)
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
		c.DrawLine([2]int{int(x0), int(y)}, [2]int{int(x1), int(y)}, cell)
		x0 -= slope0
		x1 -= slope1
	}
}

func isTriangleOnScreen(tri Triangle3) bool {
	return tri.IntersectsBox2(Box2{Point2{-1, -1}, Point2{1, 1}})
}
func isTriangleOffScreen(tri Triangle3) bool {
	return !isTriangleOnScreen(tri)
}

// Triangle coords are -1.0 to +1.0
// Vertex order: [top, br, bl]
func (c *Canvas) fillFlatBottomTriangle3(tri Triangle3, cell Cell) {
	if isTriangleOffScreen(tri) {
		return
	}

	// Translate to pixel coords
	p0 := c.Point3ToPoint2(tri[0]).Floor()
	p1 := c.Point3ToPoint2(tri[1]).Floor()
	p2 := c.Point3ToPoint2(tri[2]).Floor()

	dy := p1.Y() - p0.Y()
	slope0 := (p1.X() - p0.X()) / dy
	slope1 := (p2.X() - p0.X()) / dy
	zSlope0 := (tri[1].Z() - tri[0].Z()) / dy
	zSlope1 := (tri[2].Z() - tri[0].Z()) / dy

	x0 := p0.X()
	x1 := p0.X()
	z0 := tri[0].Z()
	z1 := tri[0].Z()

	for y := p0.Y(); y <= p1.Y(); y++ {
		var sx0, sx1 float64
		if slope0 < slope1 {
			sx0 = Floor(x0)
			sx1 = Floor(x1)
		} else {
			sx0 = Floor(x0)
			sx1 = Floor(x1)
		}

		line := Line3{Point3{sx0, y, z0}, Point3{sx1, y, z1}}
		c.DrawDeepLine(line, cell)
		x0 += slope0
		x1 += slope1
		z0 += zSlope0
		z1 += zSlope1
	}
}

// Vertex order: [tl, tr, b]
func (c *Canvas) fillFlatTopTriangle3(tri Triangle3, cell Cell) {
	if isTriangleOffScreen(tri) {
		return
	}

	p0 := c.Point3ToPoint2(tri[0]).Floor()
	p1 := c.Point3ToPoint2(tri[1]).Floor()
	p2 := c.Point3ToPoint2(tri[2]).Floor()

	dy := p2.Y() - p0.Y()
	slope0 := (p2.X() - p0.X()) / dy
	slope1 := (p2.X() - p1.X()) / dy
	zSlope0 := (tri[0].Z() - tri[2].Z()) / dy
	zSlope1 := (tri[1].Z() - tri[2].Z()) / dy

	x0 := float64(p2.X())
	x1 := float64(p2.X())
	z0 := tri[2].Z()
	z1 := tri[2].Z()

	for y := p2.Y(); y >= p0.Y(); y-- {
		var sx0, sx1 float64
		if slope0 > slope1 {
			sx0 = Floor(x0)
			sx1 = Floor(x1)
		} else {
			sx0 = Floor(x0)
			sx1 = Floor(x1)
		}

		line := Line3{Point3{sx0, y, z0}, Point3{sx1, y, z1}}
		c.DrawDeepLine(line, cell)
		x0 -= slope0
		x1 -= slope1
		z0 += zSlope0
		z1 += zSlope1
	}
}

func (c *Canvas) DrawTriangle3(tri Triangle3, cell Cell) {
	// Sort by Y axis
	sort.SliceStable(tri[:], func(i, j int) bool {
		return tri[i].Y() < tri[j].Y()
	})

	if tri[1].Y() == tri[2].Y() {
		c.fillFlatBottomTriangle3(tri, cell)

	} else if tri[0].Y() == tri[1].Y() {
		c.fillFlatTopTriangle3(tri, cell)

	} else {
		dy := (tri[1].Y() - tri[0].Y()) / (tri[2].Y() - tri[0].Y())

		midVert := Point3{
			tri[0].X() + dy*(tri[2].X()-tri[0].X()),
			tri[1].Y(),
			tri[0].Z() + dy*(tri[2].Z()-tri[0].Z()),
		}

		c.fillFlatBottomTriangle3(Triangle3{tri[0], midVert, tri[1]}, cell)
		c.fillFlatTopTriangle3(Triangle3{tri[1], midVert, tri[2]}, cell)
	}
}

func (c *Canvas) DrawWireTriangle3(tri Triangle3, cell Cell) {
	p0 := c.ScreenPoint3ToCellPoint3(tri[0]).Floor()
	p1 := c.ScreenPoint3ToCellPoint3(tri[1]).Floor()
	p2 := c.ScreenPoint3ToCellPoint3(tri[2]).Floor()
	line := Line3{p0, p1}
	c.DrawDeepLine(line, cell)
	line = Line3{p1, p2}
	c.DrawDeepLine(line, cell)
	line = Line3{p2, p0}
	c.DrawDeepLine(line, cell)
}

func (c *Canvas) ScreenPoint3ToCellPoint3(point Point3) Point3 {
	hw := float64(c.Width) / 2
	hh := float64(c.Height) / 2

	x := point.X()*hw + hw
	y := point.Y()*hh + hh

	return Point3{x, y, point.Z()}
}

func (c *Canvas) Point3ToPoint2(point Point3) Point2 {
	hw := float64(c.Width) / 2
	hh := float64(c.Height) / 2

	x := point.X()*hw + hw
	y := point.Y()*hh + hh

	return Point2{x, y}
}

func (c *Canvas) Point2ToCoord(point Point2) [2]int {
	hw := float64(c.Width) / 2
	hh := float64(c.Height) / 2

	x := int(point.X()*hw + hw)
	y := int(point.Y()*hh + hh)

	return [2]int{x, y}
}

func (c *Canvas) Point3ToCoord(point Point3) [2]int {
	hw := float64(c.Width) / 2
	hh := float64(c.Height) / 2

	x := int(point.X()*hw + hw)
	y := int(point.Y()*hh + hh)

	return [2]int{x, y}
}

func (c *Canvas) DrawVectorTriangle(tri TriangleFloat, cell Cell) {
	// Center and scale coordinates
	hw := float64(c.Width) / 2
	hh := float64(c.Height) / 2
	c.DrawTriangle(
		Triangle{
			[2]int{
				int(tri[0][0]*hw + hw),
				int(tri[0][1]*hh + hh),
			},
			[2]int{
				int(tri[1][0]*hw + hw),
				int(tri[1][1]*hh + hh),
			},
			[2]int{
				int(tri[2][0]*hw + hw),
				int(tri[2][1]*hh + hh),
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

	floatTri := tri.ToFloat()

	if floatTri[1][1] == floatTri[2][1] {
		c.fillFlatBottomTriangle(floatTri, cell)

	} else if floatTri[0][1] == floatTri[1][1] {
		c.fillFlatTopTriangle(floatTri, cell)

	} else {
		midVert := [2]float64{
			floatTri[0][0] + (floatTri[1][1]-floatTri[0][1])/(floatTri[2][1]-floatTri[0][1])*(floatTri[2][0]-floatTri[0][0]),
			floatTri[1][1],
		}

		c.fillFlatTopTriangle(TriangleFloat{floatTri[1], midVert, floatTri[2]}, cell)
		c.fillFlatBottomTriangle(TriangleFloat{floatTri[0], floatTri[1], midVert}, cell)
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
	var cursorColor string = ""
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			backCell := c.GetBack(x, y)
			frontCell := c.GetFront(x, y)

			if backCell == nil || frontCell == nil || backCell == frontCell {
				continue
			}

			if x != cursorX || y != cursorY {
				cursorX = x
				cursorY = y
				term.MoveTo(Position{x, y})
			}

			// Send colour only if it has changed
			color := backCell.Ansi24BitColor()
			if color != cursorColor {
				cursorColor = color
				term.Write(color)
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
	return x + y*c.Width
}

func (c *Canvas) indexToPosition(idx int) (int, int) {
	x := idx % c.Width
	y := idx / c.Width
	return x, y
}
