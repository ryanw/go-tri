package geom

import "math"

type Vector3 [3]float64
type Point3 [3]float64
type Vector4 [4]float64
type Point4 [4]float64
type Matrix4 [16]float64
type Line3 [2]Point3

func NewMatrix4Identity() Matrix4 {
  return Matrix4 {
    1, 0, 0, 0,
    0, 1, 0, 0,
    0, 0, 1, 0,
    0, 0, 0, 1,
  }
}

func NewMatrix4FromColumns(cols [4]Vector4) Matrix4 {
  return Matrix4 {
    cols[0][0], cols[1][0], cols[2][0], cols[3][0],
    cols[0][1], cols[1][1], cols[2][1], cols[3][1],
    cols[0][2], cols[1][2], cols[2][2], cols[3][2],
    cols[0][3], cols[1][3], cols[2][3], cols[3][3],
  }
}

func NewMatrix4Perspective(aspect, fov, near, far float64) Matrix4 {
  fovRad := fov * (math.Pi / 180)
  f := 1.0 / math.Tan(fovRad / 2.0)
  r := 1.0 / (near - far)
  return Matrix4{
    f / aspect, 0,                    0,                      0,
    0,          f,                    0,                      0,
    0,          0,     (near + far) * r,     near * far * r * 2,
    0,          0,                   -1,                      0,
  }
}

func NewMatrix4Rotation(x, y, z float64) Matrix4 {

  cosx, sinx := math.Cos(x), math.Sin(x)
  cosy, siny := math.Cos(y), math.Sin(y)
  cosz, sinz := math.Cos(z), math.Sin(z)

  rotx := Matrix4 {
    1,    0,     0, 0,
    0, cosx, -sinx, 0,
    0, sinx,  cosx, 0,
    0,    0,     0, 1,
  }

  roty := Matrix4 {
     cosy, 0, siny, 0,
        0, 1,    0, 0,
    -siny, 0, cosy, 0,
        0, 0,    0, 1,
  }

  rotz := Matrix4 {
    cosz, -sinz, 0, 0,
    sinz,  cosz, 0, 0,
       0,     0, 1, 0,
       0,     0, 0, 1,
  }

  return rotx.Multiply(roty.Multiply(rotz))
}

func NewMatrix4Translation(x, y, z float64) Matrix4 {
  return Matrix4 {
    1, 0, 0, x,
    0, 1, 0, y,
    0, 0, 1, z,
    0, 0, 0, 1,
  }
}

func NewMatrix4Scaling(x, y, z float64) Matrix4 {
  return Matrix4 {
    x, 0, 0, 0,
    0, y, 0, 0,
    0, 0, z, 0,
    0, 0, 0, 1,
  }
}

func (self Matrix4) Row(axis int64) Vector4 {
  o := 4 * axis
  return Vector4 {
    self[o + 0],
    self[o + 1],
    self[o + 2],
    self[o + 3],
  }
}

func (self Matrix4) Rows() [4]Vector4 {
  return [4]Vector4 {
    self.Row(0),
    self.Row(1),
    self.Row(2),
    self.Row(3),
  }
}

func (self Matrix4) Column(axis int64) Vector4 {
  return Vector4 {
    self[axis + 0],
    self[axis + 4],
    self[axis + 8],
    self[axis + 12],
  }
}

func (self Matrix4) Columns() [4]Vector4 {
  return [4]Vector4 {
    self.Column(0),
    self.Column(1),
    self.Column(2),
    self.Column(3),
  }
}

func (self Matrix4) Multiply(other Matrix4) Matrix4 {
  otherCols := other.Columns()
  cols := [4]Vector4 {
    self.MultiplyVector4(otherCols[0]),
    self.MultiplyVector4(otherCols[1]),
    self.MultiplyVector4(otherCols[2]),
    self.MultiplyVector4(otherCols[3]),
  }
  return NewMatrix4FromColumns(cols);
}

func (self Matrix4) MultiplyVector4(vec Vector4) Vector4 {
  cols := self.Columns()

  x := cols[0].Scale(vec[0])
  y := cols[1].Scale(vec[1])
  z := cols[2].Scale(vec[2])
  w := cols[3].Scale(vec[3])

  return Vector4 {
    (x[0] + y[0] + z[0] + w[0]),
    (x[1] + y[1] + z[1] + w[1]),
    (x[2] + y[2] + z[2] + w[2]),
    (x[3] + y[3] + z[3] + w[3]),
  }
}

func (self Matrix4) TransformPoint3(point Point3) Point3 {
  vec := self.MultiplyVector4(Vector4 { point[0], point[1], point[2], 1 })

  return Point3 {
    vec[0] / vec[3],
    vec[1] / vec[3],
    vec[2] / vec[3],
  }
}

func (self Vector4) Scale(scale float64) Vector4 {
  return Vector4 {
    self[0] * scale,
    self[1] * scale,
    self[2] * scale,
    self[3] * scale,
  }
}
