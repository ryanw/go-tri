package mesh

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	. "tri/geom"
)

func NewMeshFromObjPath(path string) (TriangleMesh, error) {
	mesh := TriangleMesh{}

	f, err := os.Open(path)
	if err != nil {
		return mesh, err
	}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		cat := strings.Split(line, " ")[0]
		switch cat {
		case "v": // Vertex
			mesh.Vertices = append(mesh.Vertices, parseVertex(line))
		case "f": // Face
			mesh.Triangles = append(mesh.Triangles, parseFace(line))
			mesh.Colors = append(mesh.Colors, 0xffaaaaaa)
		case "vt": // Vertex Texture
		case "vn": // Vertex Normal
		}
	}

	return mesh, nil
}

func parseVertex(line string) Point3 {
	p := Point3{}

	tokens := strings.Split(line, " ")

	for i := 0; i < 3; i++ {
		p[i], _ = strconv.ParseFloat(strings.TrimSpace(tokens[i+1]), 64)
		p[i] *= -1
	}
	/*
		if len(tokens) > 4 {
			w, _ := strconv.ParseFloat(tokens[4], 64)
			p[0] /= w
			p[1] /= w
			p[2] /= w
		}
	*/

	return p
}

func parseFace(line string) [3]int {
	tri := [3]int{}

	tokens := strings.Split(line, " ")

	for i := 0; i < 3; i++ {
		indexes := strings.Split(strings.TrimSpace(tokens[i+1]), "/")
		idx, _ := strconv.ParseInt(indexes[0], 10, 64)

		tri[i] = int(idx - 1)
	}

	return tri
}
