package renderer

import (
	. "tri/canvas"
	. "tri/geom"
	. "tri/mesh"
)

type Renderer struct {
	Camera Camera
}

func (r *Renderer) RenderLineMesh(canvas *Canvas, mesh *LineMesh) {
	camera := &r.Camera

	mvp := camera.Projection
	mvp = mvp.Multiply(camera.Transform.Matrix())
	mvp = mvp.Multiply(mesh.Transform.Matrix())

	for _, line := range mesh.Lines {
		start := mvp.TransformPoint3(mesh.Vertices[line[0]])
		end := mvp.TransformPoint3(mesh.Vertices[line[1]])

		canvas.DrawLine3D(start, end, Cell{
			Fg:     0xff00ff00,
			Bg:     0xffaa0000,
			Sprite: ' ',
		})
	}
}

func (r *Renderer) RenderTriangleMesh(canvas *Canvas, mesh *TriangleMesh) {
	camera := &r.Camera
	lightDir := Vector3{0.4, -0.7, -0.3}.Normalize()
	_ = lightDir

	mvp := camera.Projection
	mvp = mvp.Multiply(camera.Transform.Matrix())
	mvp = mvp.Multiply(mesh.Transform.Matrix())

	for i, tri := range mesh.Triangles {
		triangle := mvp.TransformTriangle3(Triangle3{
			mesh.Vertices[tri[0]],
			mesh.Vertices[tri[1]],
			mesh.Vertices[tri[2]],
		})

		normal := mvp.TransformVector3(mesh.Normals[i])
		ambient := 0.2
		diffuse := normal.Dot(lightDir)
		light := ambient + diffuse
		if light < ambient {
			light = ambient
		}
		if light > 1 {
			light = 1
		}
		color := Vector3FromColor(mesh.Colors[i])
		color = color.Scale(light)

		canvas.DrawTriangle3(triangle, Cell{
			Fg:     color.Scale(0.7).ToColor(),
			Bg:     color.ToColor(),
			Sprite: ' ',
		})
	}
}
