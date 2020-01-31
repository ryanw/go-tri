package geom

import (
	"math"
	"testing"
)

func assertValuesEqual(t *testing.T, actual []float64, expected []float64) {
	if len(actual) != len(expected) {
		t.Errorf("Lens differ: %#v != %#v", actual, expected)
		return
	}
	for i, n := range actual {
		if math.Abs(n-expected[i]) > 0.001 {
			t.Errorf("Value at %d is wrong  %#v != %#v", i, actual, expected)
			break
		}
	}
}

func assertMatrix4Equal(t *testing.T, actual Matrix4, expected Matrix4) {
	assertValuesEqual(t, actual[:], expected[:])
}

func assertVector4Equal(t *testing.T, actual Vector4, expected Vector4) {
	assertValuesEqual(t, actual[:], expected[:])
}

func assertPoint3Equal(t *testing.T, actual Point3, expected Point3) {
	assertValuesEqual(t, actual[:], expected[:])
}
