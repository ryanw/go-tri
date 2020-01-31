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

	viewProj := camera.ViewProjection()
	mvp := viewProj.Multiply(mesh.Transform.Matrix())

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

func clipTriangle(plane Plane3, tri Triangle3) []Triangle3 {
	inside := []Point3{}
	outside := []Point3{}

	for _, point := range tri {
		if point.DistanceToPlane3(plane) <= 0 {
			inside = append(inside, point)
			// TODO actually clip the triangle. For now just remove it
			return []Triangle3{}
		} else {
			outside = append(outside, point)
		}
	}

	return []Triangle3{tri}
}

func (r *Renderer) RenderTriangleMesh(canvas *Canvas, mesh *TriangleMesh) {
	camera := &r.Camera
	lightDir := Vector3{0.4, -0.7, -0.3}.Normalize()

	proj := camera.Projection
	view := camera.View()
	model := mesh.Transform.Matrix()
	nearPlane := Plane3{
		Point:  Point3{0.0, 0.0, -0.1},
		Normal: Vector3{0.0, 0.0, -1.0},
	}

	for i, triIndexes := range mesh.Triangles {
		triangle := Triangle3{
			mesh.Vertices[triIndexes[0]],
			mesh.Vertices[triIndexes[1]],
			mesh.Vertices[triIndexes[2]],
		}

		// Move to world space
		triangle = model.TransformTriangle3(triangle)

		// Calculate normal in world space
		line1 := triangle[2].ToVector3().Sub(triangle[0].ToVector3())
		line2 := triangle[1].ToVector3().Sub(triangle[0].ToVector3())
		normal := line1.Cross(line2).Normalize()

		// Move to view space
		triangle = view.TransformTriangle3(triangle)

		// Clip triangles outside the view frustrum
		triangles := clipTriangle(nearPlane, triangle)

		for _, triangle = range triangles {
			triangle = proj.TransformTriangle3(triangle)

			ambient := 0.1
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
}

func (r *Renderer) RenderWireTriangleMesh(canvas *Canvas, mesh *TriangleMesh) {
	camera := &r.Camera
	lightDir := Vector3{0.4, -0.7, -0.3}.Normalize()

	proj := camera.Projection
	view := camera.View()
	model := mesh.Transform.Matrix()
	modelView := view.Multiply(model)
	plane := Plane3{
		Point:  Point3{0.0, 0.0, -0.01},
		Normal: Vector3{0.0, 0.0, -1.0},
	}

	for i, triIndexes := range mesh.Triangles {
		triangle := Triangle3{
			mesh.Vertices[triIndexes[0]],
			mesh.Vertices[triIndexes[1]],
			mesh.Vertices[triIndexes[2]],
		}
		triangle = modelView.TransformTriangle3(triangle)

		// Clip triangles outside the view frustrum
		triangles := clipTriangle(plane, triangle)

		for _, tri := range triangles {
			triangle = proj.TransformTriangle3(tri)

			normal := model.TransformVector3(mesh.Normals[i])
			ambient := 0.01
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

			canvas.DrawWireTriangle3(triangle, Cell{
				Fg:     0xff000000,
				Bg:     color.Scale(0.7).ToColor(),
				Sprite: 'x',
			})
		}
	}
}
