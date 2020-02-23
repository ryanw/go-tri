package geom

type Color uint32

func ColorFromRgba(r, g, b, a float32) Color {
	ra := uint32(a*255) << 24
	rr := uint32(r*255) << 16
	rg := uint32(g*255) << 8
	rb := uint32(b*255) << 0
	return Color(ra + rr + rg + rb)
}

func (c Color) ToRgb() (float32, float32, float32) {
	r, g, b, _ := c.ToRgba()
	return r, g, b
}

func (c Color) ToRgba() (float32, float32, float32, float32) {
	var r, g, b, a float32
	a = float32((c&0xff000000)>>24) / 255.0
	r = float32((c&0x00ff0000)>>16) / 255.0
	g = float32((c&0x0000ff00)>>8) / 255.0
	b = float32((c&0x000000ff)>>0) / 255.0
	return r, g, b, a
}

func (c Color) ToAnsi() uint16 {
	//a := ((c & 0xff000000) >> 24) / 51
	r := ((c & 0x00ff0000) >> 16) / 51
	g := ((c & 0x0000ff00) >> 8) / 51
	b := ((c & 0x000000ff) >> 0) / 51
	return uint16(16 + 36*r + 6*g + b)
}

func (dst Color) Blend(src Color) Color {
	var r, g, b, a float32
	dr, dg, db, da := dst.ToRgba()
	sr, sg, sb, sa := src.ToRgba()

	r = (sr * sa) + (dr * (1.0 - sa))
	g = (sg * sa) + (dg * (1.0 - sa))
	b = (sb * sa) + (db * (1.0 - sa))
	a = (sa * sa) + (da * (1.0 - sa))

	return ColorFromRgba(r, g, b, a)
}
