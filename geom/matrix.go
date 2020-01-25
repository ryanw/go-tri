package geom

import "math"

func NewMatrix4Identity() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func NewMatrix4FromColumns(cols [4]Vector4) Matrix4 {
	return Matrix4{
		cols[0][0], cols[1][0], cols[2][0], cols[3][0],
		cols[0][1], cols[1][1], cols[2][1], cols[3][1],
		cols[0][2], cols[1][2], cols[2][2], cols[3][2],
		cols[0][3], cols[1][3], cols[2][3], cols[3][3],
	}
}

func NewMatrix4Perspective(aspect, fov, near, far float64) Matrix4 {
	fovRad := fov * (math.Pi / 180)
	f := 1.0 / math.Tan(fovRad/2.0)
	r := 1.0 / (near - far)
	return Matrix4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (near + far) * r, near * far * r * 2,
		0, 0, -1, 0,
	}
}

func NewMatrix4Rotation(x, y, z float64) Matrix4 {

	cosx, sinx := math.Cos(x), math.Sin(x)
	cosy, siny := math.Cos(y), math.Sin(y)
	cosz, sinz := math.Cos(z), math.Sin(z)

	rotx := Matrix4{
		1, 0, 0, 0,
		0, cosx, -sinx, 0,
		0, sinx, cosx, 0,
		0, 0, 0, 1,
	}

	roty := Matrix4{
		cosy, 0, siny, 0,
		0, 1, 0, 0,
		-siny, 0, cosy, 0,
		0, 0, 0, 1,
	}

	rotz := Matrix4{
		cosz, -sinz, 0, 0,
		sinz, cosz, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	return rotx.Multiply(roty.Multiply(rotz))
}

func NewMatrix4Translation(x, y, z float64) Matrix4 {
	return Matrix4{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

func NewMatrix4Scaling(x, y, z float64) Matrix4 {
	return Matrix4{
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1,
	}
}

func (m Matrix4) Row(axis int64) Vector4 {
	o := 4 * axis
	return Vector4{
		m[o+0],
		m[o+1],
		m[o+2],
		m[o+3],
	}
}

func (m Matrix4) Rows() [4]Vector4 {
	return [4]Vector4{
		m.Row(0),
		m.Row(1),
		m.Row(2),
		m.Row(3),
	}
}

func (m Matrix4) Column(axis int64) Vector4 {
	return Vector4{
		m[axis+0],
		m[axis+4],
		m[axis+8],
		m[axis+12],
	}
}

func (m Matrix4) Columns() [4]Vector4 {
	return [4]Vector4{
		m.Column(0),
		m.Column(1),
		m.Column(2),
		m.Column(3),
	}
}

func (m Matrix4) Multiply(other Matrix4) Matrix4 {
	otherCols := other.Columns()
	cols := [4]Vector4{
		m.MultiplyVector4(otherCols[0]),
		m.MultiplyVector4(otherCols[1]),
		m.MultiplyVector4(otherCols[2]),
		m.MultiplyVector4(otherCols[3]),
	}
	return NewMatrix4FromColumns(cols)
}

func (m Matrix4) MultiplyVector4(vec Vector4) Vector4 {
	cols := m.Columns()

	x := cols[0].Scale(vec[0])
	y := cols[1].Scale(vec[1])
	z := cols[2].Scale(vec[2])
	w := cols[3].Scale(vec[3])

	return Vector4{
		(x[0] + y[0] + z[0] + w[0]),
		(x[1] + y[1] + z[1] + w[1]),
		(x[2] + y[2] + z[2] + w[2]),
		(x[3] + y[3] + z[3] + w[3]),
	}
}

func (m Matrix4) TransformVector3(vector Vector3) Vector3 {
	vec := m.MultiplyVector4(Vector4{vector[0], vector[1], vector[2], 0})

	return Vector3{
		vec[0],
		vec[1],
		vec[2],
	}
}

func (m Matrix4) TransformPoint3(point Point3) Point3 {
	vec := m.MultiplyVector4(Vector4{point[0], point[1], point[2], 1})

	return Point3{
		vec[0] / vec[3],
		vec[1] / vec[3],
		vec[2] / vec[3],
	}
}

func (m Matrix4) TransformTriangle3(tri Triangle3) Triangle3 {
	return Triangle3{
		m.TransformPoint3(tri[0]),
		m.TransformPoint3(tri[1]),
		m.TransformPoint3(tri[2]),
	}
}

func (m Vector4) Scale(scale float64) Vector4 {
	return Vector4{
		m[0] * scale,
		m[1] * scale,
		m[2] * scale,
		m[3] * scale,
	}
}
