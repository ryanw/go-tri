package geom

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
