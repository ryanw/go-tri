package geom

type Color uint32

func (c *Color) ToRgb() (uint8, uint8, uint8) {
  return 0, 0, 0
}

func (c *Color) ToRgba() (uint8, uint8, uint8, uint8) {
  return 0, 0, 0, 0
}

func (c Color) ToAnsi() uint16 {
  //a := ((c & 0xff000000) >> 24) / 51
  r := ((c & 0x00ff0000) >> 16) / 51
  g := ((c & 0x0000ff00) >> 8) / 51
  b := ((c & 0x000000ff) >> 0) / 51
  return uint16(16 + 36 * r + 6 * g + b)
}
