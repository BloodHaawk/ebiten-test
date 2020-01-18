package collision

import "math"

type collisionBox interface {
	PosX() float64
	PosY() float64
	VX() float64
	VY() float64
	SizeX() float64
	SizeY() float64
}

type vertex struct {
	x, y float64
}

func segmentCollision(v1a, v1b, v2a, v2b vertex) bool {
	s1 := vertex{v1b.x - v1a.x, v1b.y - v1a.y}
	s2 := vertex{v2b.x - v2a.x, v2b.y - v2a.y}

	det := -s1.x*s2.y + s1.y*s2.x

	if math.Abs(det) < 1e-16 {
		return false
	}

	// Intersect the two lines
	t := vertex{
		1 / det * ((v2a.x-v1a.x)*-s2.y + (v2a.y-v1a.y)*s2.x),
		1 / det * ((v2a.x-v1a.x)*-s1.y + (v2a.y-v1a.y)*s1.x),
	}

	// Check if the intersection is between the end-points
	return (0 <= t.x && t.x <= 1 && 0 <= t.y && t.y <= 1)
}

func vertexInside(v, qva, qvb, qvc, qvd vertex) bool {

	quad := [4]vertex{qva, qvb, qvc, qvd}
	edgeIndices := [4][2]int{[2]int{0, 1}, [2]int{1, 2}, [2]int{2, 3}, [2]int{3, 0}}

	for i := range edgeIndices {
		edge := [2]vertex{quad[edgeIndices[i][0]], quad[edgeIndices[i][1]]}
		oppEdge := [2]vertex{quad[(edgeIndices[i][0]+2)%4], quad[(edgeIndices[i][1]+2)%4]}

		if segmentCollision(v, oppEdge[0], edge[0], edge[1]) || segmentCollision(v, oppEdge[1], edge[0], edge[1]) {
			return false
		}
	}

	return true
}

func quadrangleCollision(v1a, v1b, v1c, v1d, v2a, v2b, v2c, v2d vertex) bool {
	quad1 := [4]vertex{v1a, v1b, v1c, v1d}
	quad2 := [4]vertex{v2a, v2b, v2c, v2d}
	edgeIndices := [4][2]int{[2]int{0, 1}, [2]int{1, 2}, [2]int{2, 3}, [2]int{3, 0}}

	for i := range quad1 {
		if vertexInside(quad1[i], v2a, v2b, v2c, v2d) {
			return true
		}
	}
	for i := range quad2 {
		if vertexInside(quad2[i], v1a, v1b, v1c, v1d) {
			return true
		}
	}
	for i := range edgeIndices {
		for j := range edgeIndices {
			if segmentCollision(quad1[edgeIndices[i][0]], quad1[edgeIndices[i][1]], quad2[edgeIndices[j][0]], quad2[edgeIndices[j][1]]) {
				return true
			}
		}
	}
	return false
}

func staticCollision(c1, c2 collisionBox) bool {
	x1In := c2.PosX() <= c1.PosX() && c1.PosX() < c2.PosX()+c2.SizeX()
	x2In := c1.PosX() <= c2.PosX() && c2.PosX() < c1.PosX()+c1.SizeX()
	y1In := c2.PosY() <= c1.PosY() && c1.PosY() < c2.PosY()+c2.SizeY()
	y2In := c1.PosY() <= c2.PosY() && c2.PosY() < c1.PosY()+c1.SizeY()

	return (x1In || x2In) && (y1In || y2In)
}

func dynamicCollision(c1, c2 collisionBox) bool {
	var quad1, quad2 [4]vertex
	if c1.VX()*c1.VY() < 0 {
		quad1 = [4]vertex{
			vertex{c1.PosX(), c1.PosY()},
			vertex{c1.PosX() + c1.SizeX(), c1.PosY() + c1.SizeY()},
			vertex{c1.PosX() + c1.SizeX() + c1.VX(), c1.PosY() + c1.SizeY() + c1.VY()},
			vertex{c1.PosX() + c1.VX(), c1.PosY() + c1.VY()},
		}
	} else {
		quad1 = [4]vertex{
			vertex{c1.PosX(), c1.PosY() + c1.SizeY()},
			vertex{c1.PosX() + c1.SizeX(), c1.PosY()},
			vertex{c1.PosX() + c1.SizeX() + c1.VX(), c1.PosY() + c1.VY()},
			vertex{c1.PosX() + c1.VX(), c1.PosY() + c1.SizeY() + c1.VY()},
		}
	}
	if c2.VX()*c2.VY() < 0 {
		quad2 = [4]vertex{
			vertex{c2.PosX(), c2.PosY()},
			vertex{c2.PosX() + c2.SizeX(), c2.PosY() + c2.SizeY()},
			vertex{c2.PosX() + c2.SizeX() + c2.VX(), c2.PosY() + c2.SizeY() + c2.VY()},
			vertex{c2.PosX() + c2.VX(), c2.PosY() + c2.VY()},
		}
	} else {
		quad2 = [4]vertex{
			vertex{c2.PosX(), c2.PosY() + c2.SizeY()},
			vertex{c2.PosX() + c2.SizeX(), c2.PosY()},
			vertex{c2.PosX() + c2.SizeX() + c2.VX(), c2.PosY() + c2.VY()},
			vertex{c2.PosX() + c2.VX(), c2.PosY() + c2.SizeY() + c2.VY()},
		}

	}
	return quadrangleCollision(quad1[0], quad1[1], quad1[2], quad1[3], quad2[0], quad2[1], quad2[2], quad2[3])
}

// Collision returns true if two collisionBoxes are colliding
func Collision(c1, c2 collisionBox) bool {
	return staticCollision(c1, c2) || dynamicCollision(c1, c2)
}
