package renderer

import (
	. "tri/canvas"
	. "tri/geom"
)

type Renderer struct {
	Camera Camera
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

func (r *Renderer) RenderDrawable(canvas *Canvas, mesh Drawable) int {
	count := 0
	camera := &r.Camera
	lightDir := Vector3{0.4, -0.7, -0.3}.Normalize()

	proj := camera.Projection
	view := camera.View()
	nearPlane := Plane3{
		Point:  Point3{0.0, 0.0, -0.1},
		Normal: Vector3{0.0, 0.0, -1.0},
	}

	ch := make(chan Polygon, 100)
	go func() {
		mesh.DrawTriangles(ch)
		close(ch)
	}()

	for {
		poly, ok := <-ch
		if ok == false {
			break
		}

		triangle := poly.Shape
		color := Vector3FromColor(poly.Color)

		// Calculate normal in world space
		line1 := triangle[2].ToVector3().Sub(triangle[0].ToVector3())
		line2 := triangle[1].ToVector3().Sub(triangle[0].ToVector3())
		normal := line1.Cross(line2).Normalize()

		// Move to view space
		triangle = view.TransformTriangle3(triangle)

		// Clip triangles outside the view frustrum
		triangles := clipTriangle(nearPlane, triangle)

		for _, triangle = range triangles {
			count += 1
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
			color = color.Scale(light)

			canvas.DrawTriangle3(triangle, Cell{
				Fg:     color.Scale(0.7).ToColor(),
				Bg:     color.ToColor(),
				Sprite: ' ',
			})
		}
	}

	return count
}
