package canvas

import (
	"fmt"
	. "tri/geom"
)

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

func (dst Cell) Blend(src Cell) Cell {
	dst.Fg = dst.Fg.Blend(src.Fg)
	dst.Bg = dst.Bg.Blend(src.Bg)
	dst.Sprite = src.Sprite
	if src.Depth > dst.Depth {
		dst.Depth = src.Depth
	}

	return dst
}
