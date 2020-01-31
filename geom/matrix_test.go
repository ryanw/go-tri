package geom

import (
	"math"
	"testing"
)

func TestNewMatrix4Identity(t *testing.T) {
	mat := NewMatrix4Identity()
	expected := Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	assertMatrix4Equal(t, mat, expected)
}

func TestMatrix4Inverse(t *testing.T) {
	mat := Matrix4{
		3, 7, 2, 3,
		3, 1, 3, 5,
		5, 4, 2, 0,
		8, 5, 1, 1,
	}
	expected := Matrix4{
		-0.2564, 0.0769, -0.0512, 0.3846,
		0.4102, -0.2371, -0.0320, -0.0448,
		-0.1794, 0.2820, 1.3333, -0.8717,
		0.1794, 0.2884, -0.7628, 0.3012,
	}

	result, err := mat.Inverse()

	if err != nil {
		t.Errorf("Matrix failed to invert - %v", err)
	}

	assertMatrix4Equal(t, result, expected)

	// Inverse again should return the original
	result, err = result.Inverse()

	if err != nil {
		t.Errorf("Matrix failed to invert - %v", err)
	}

	assertMatrix4Equal(t, result, mat)
}

func TestMatrix4InverseTranslation(t *testing.T) {
	mat := Matrix4{
		1, 0, 0, 5,
		0, 1, 0, 4,
		0, 0, 1, 3,
		0, 0, 0, 1,
	}
	expected := Matrix4{
		1, 0, 0, -5,
		0, 1, 0, -4,
		0, 0, 1, -3,
		0, 0, 0, 1,
	}

	result, err := mat.Inverse()

	if err != nil {
		t.Errorf("Matrix failed to invert - %v", err)
	}

	assertMatrix4Equal(t, result, expected)

	// Inverse again should return the original
	result, err = result.Inverse()

	if err != nil {
		t.Errorf("Matrix failed to invert - %v", err)
	}

	assertMatrix4Equal(t, result, mat)
}

func TestMultiply(t *testing.T) {
	mat1 := Matrix4{
		1, 0, 2, 0,
		0, 1, 0, 7,
		4, 0, 3, 1,
		0, 2, 0, 2,
	}
	mat2 := Matrix4{
		3, 7, 2, 3,
		3, 1, 3, 5,
		5, 4, 2, 0,
		8, 5, 1, 1,
	}

	result := mat1.Multiply(mat2)
	expected := Matrix4{
		13, 15, 6, 3,
		59, 36, 10, 12,
		35, 45, 15, 13,
		22, 12, 8, 12,
	}

	assertMatrix4Equal(t, result, expected)
}

func TestTransformPoint3(t *testing.T) {
	mat := Matrix4{
		1, 0, 2, 0,
		0, 1, 0, 7,
		4, 0, 3, 1,
		0, 2, 0, 2,
	}
	point := Point3{4, 5, 6}
	result := mat.TransformPoint3(point)

	assertPoint3Equal(t, result, Point3{1.3333, 1.0000, 2.9170})
}

func TestMultiplyVector4(t *testing.T) {
	mat := Matrix4{
		1, 0, 2, 0,
		0, 1, 0, 7,
		4, 0, 3, 1,
		0, 2, 0, 2,
	}
	vec := Vector4{4, 5, 6, 1}
	result := mat.MultiplyVector4(vec)

	assertVector4Equal(t, result, Vector4{16.0, 12.0, 35.0, 12.0})
}

func TestCreateRotation(t *testing.T) {
	mat := NewMatrix4Rotation(0.0, math.Pi*0.25, 0.0)
	expected := Matrix4{
		0.707, 0.0, 0.707, 0.0,
		0.0, 1.0, 0.0, 0.0,
		-0.707, 0.0, 0.707, 0.0,
		0.0, 0.0, 0.0, 1.0,
	}

	assertMatrix4Equal(t, mat, expected)
}

func TestRotatePoint(t *testing.T) {
	point := Point3{4, 5, 6}
	mat := NewMatrix4Rotation(0.0, math.Pi*0.25, 0.0)
	result := mat.TransformPoint3(point)
	expected := Point3{7.071, 5.00, 1.414}

	assertPoint3Equal(t, result, expected)
}

func TestCreatePerspective(t *testing.T) {
	mat := NewMatrix4Perspective(16.0/9.0, 45.0, 0.1, 10.0)
	expected := Matrix4{
		1.3579, 0, 0, 0,
		0, 2.4142, 0, 0,
		0, 0, -1.0202, -0.2020,
		0, 0, -1, 0,
	}

	assertMatrix4Equal(t, mat, expected)
}

func TestPerspectiveTransformVector4(t *testing.T) {
	mat := NewMatrix4Perspective(16.0/9.0, 45.0, 0.1, 10.0)
	vec := Vector4{4, 5, 6, 1}
	result := mat.MultiplyVector4(vec)
	expected := Vector4{5.432, 12.071, -6.323, -6.000}

	assertVector4Equal(t, result, expected)
}

func TestPerspectiveTransformPoint3(t *testing.T) {
	mat := NewMatrix4Perspective(16.0/9.0, 45.0, 0.01, 1000.0)
	point := Point3{4, 5, -6}
	result := mat.TransformPoint3(point)
	expected := Point3{0.905, 2.011, 0.996}
	assertPoint3Equal(t, result, expected)
}
