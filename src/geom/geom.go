package geom

import "math"

func SphericalToCartesian(lon, lat float64) Point3 {
  x := math.Cos(lat) * math.Sin(lon);
  y := math.Sin(lat);
  z := math.Cos(lat) * math.Cos(lon);
  return Point3 { x, y, z }
}
