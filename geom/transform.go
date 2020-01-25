package geom

type Transform struct {
	Translation Vector3
	Rotation    Vector3
	Scaling     Vector3
}

func NewTransform() Transform {
	return Transform{
		Translation: Vector3{0, 0, 0},
		Rotation:    Vector3{0, 0, 0},
		Scaling:     Vector3{1, 1, 1},
	}
}

func (t *Transform) Matrix() Matrix4 {
	return NewMatrix4Translation(
		t.Translation[0],
		t.Translation[1],
		t.Translation[2],
	).Multiply(
		NewMatrix4Rotation(
			t.Rotation[0],
			t.Rotation[1],
			t.Rotation[2],
		),
	).Multiply(
		NewMatrix4Scaling(
			t.Scaling[0],
			t.Scaling[1],
			t.Scaling[2],
		),
	)
}
