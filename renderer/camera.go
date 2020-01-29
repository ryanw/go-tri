package renderer

import . "tri/geom"

type Camera struct {
	Projection Matrix4
	Transform  Transform
}

func (c *Camera) View() Matrix4 {
	m, _ := c.Transform.Matrix().Inverse()
	return m
}

func (c *Camera) ViewProjection() Matrix4 {
	return c.Projection.Multiply(c.View())
}

func (c *Camera) Translate(x, y, z float64) {
	translation := NewMatrix4Translation(x, y, z)
	rot := c.Transform.RotationMatrix()
	invRot, _ := rot.Inverse()

	newPosition := translation.Multiply(invRot).TransformPoint3(Point3{
		c.Transform.Translation.X(),
		c.Transform.Translation.Y(),
		c.Transform.Translation.Z(),
	})

	newPosition = rot.TransformPoint3(newPosition)

	c.Transform.Translation = Vector3{
		newPosition.X(),
		newPosition.Y(),
		newPosition.Z(),
	}
}
