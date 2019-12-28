package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCollisionBox struct {
	x, y, vx, vy, xSize, ySize float64
}

func (c mockCollisionBox) posX() float64 {
	return c.x
}
func (c mockCollisionBox) posY() float64 {
	return c.y
}
func (c mockCollisionBox) vX() float64 {
	return c.vx
}
func (c mockCollisionBox) vY() float64 {
	return c.vy
}
func (c mockCollisionBox) sizeX() float64 {
	return c.xSize
}
func (c mockCollisionBox) sizeY() float64 {
	return c.ySize
}

func TestSegmentCollision(t *testing.T) {
	v1a := vertex{0, 0}
	v1b := vertex{1, -1}
	v2a := vertex{2, 1}
	v2b := vertex{0, -1}
	v3a := vertex{3, 0}
	v3b := vertex{2, -1}

	assert.False(t, segmentCollision(v1a, v1b, v1a, v1b))
	assert.True(t, segmentCollision(v1a, v1b, v2a, v2b))
	assert.False(t, segmentCollision(v1a, v1b, v3a, v3b))

	assert.True(t, segmentCollision(v2a, v2b, v1a, v1b))
	assert.False(t, segmentCollision(v2a, v2b, v2a, v2b))
	assert.False(t, segmentCollision(v2a, v2b, v3a, v3b))

	assert.False(t, segmentCollision(v3a, v3b, v1a, v1b))
	assert.False(t, segmentCollision(v3a, v3b, v2a, v2b))
	assert.False(t, segmentCollision(v3a, v3b, v3a, v3b))
}

func TestVertexInside(t *testing.T) {
	v1a := vertex{0, 0}
	v1b := vertex{1, 0}
	v1c := vertex{1, 2}
	v1d := vertex{0, 1}

	v2a := vertex{1, 1}
	v2b := vertex{2, 2}
	v2c := vertex{1, 3}
	v2d := vertex{0, 2}

	assert.False(t, vertexInside(v1a, v2a, v2b, v2c, v2d))
	assert.False(t, vertexInside(v1b, v2a, v2b, v2c, v2d))
	assert.True(t, vertexInside(v1c, v2a, v2b, v2c, v2d))
	assert.False(t, vertexInside(v1d, v2a, v2b, v2c, v2d))

	assert.False(t, vertexInside(v2a, v1a, v1b, v1c, v1d))
	assert.False(t, vertexInside(v2b, v1a, v1b, v1c, v1d))
	assert.False(t, vertexInside(v2c, v1a, v1b, v1c, v1d))
	assert.False(t, vertexInside(v2d, v1a, v1b, v1c, v1d))

}

func TestQuadCollision(t *testing.T) {
	v1a := vertex{0, 0}
	v1b := vertex{1, 0}
	v1c := vertex{1, 2}
	v1d := vertex{0, 1}

	v2a := vertex{2, 1}
	v2b := vertex{3, 1}
	v2c := vertex{3, 2}
	v2d := vertex{2, 2}

	v3a := vertex{1, 1}
	v3b := vertex{2, 2}
	v3c := vertex{1, 3}
	v3d := vertex{0, 2}

	assert.False(t, quadrangleCollision(v1a, v1b, v1c, v1d, v2a, v2b, v2c, v2d))
	assert.True(t, quadrangleCollision(v1a, v1b, v1c, v1d, v3a, v3b, v3c, v3d))
	assert.False(t, quadrangleCollision(v2a, v2b, v2c, v2d, v1a, v1b, v1c, v1d))
	assert.True(t, quadrangleCollision(v2a, v2b, v2c, v2d, v3a, v3b, v3c, v3d))
	assert.True(t, quadrangleCollision(v3a, v3b, v3c, v3d, v1a, v1b, v1c, v1d))
	assert.True(t, quadrangleCollision(v3a, v3b, v3c, v3d, v2a, v2b, v2c, v2d))
}

func TestStaticCollision(t *testing.T) {
	c1 := mockCollisionBox{0, 0, 0, 0, 10, 10}
	c2 := mockCollisionBox{-5, -5, 0, 0, 6, 7}
	c3 := mockCollisionBox{8, -5, 0, 0, 4, 6}
	c4 := mockCollisionBox{8, 8, 0, 0, 3, 5}
	c5 := mockCollisionBox{-3, 7, 0, 0, 4, 9}

	assert.True(t, staticCollision(c1, c1))
	assert.True(t, staticCollision(c1, c2))
	assert.True(t, staticCollision(c1, c3))
	assert.True(t, staticCollision(c1, c4))
	assert.True(t, staticCollision(c1, c5))

	assert.True(t, staticCollision(c2, c1))
	assert.True(t, staticCollision(c2, c2))
	assert.False(t, staticCollision(c2, c3))
	assert.False(t, staticCollision(c2, c4))
	assert.False(t, staticCollision(c2, c5))

	assert.True(t, staticCollision(c3, c1))
	assert.False(t, staticCollision(c3, c2))
	assert.True(t, staticCollision(c3, c3))
	assert.False(t, staticCollision(c3, c4))
	assert.False(t, staticCollision(c3, c5))

	assert.True(t, staticCollision(c4, c1))
	assert.False(t, staticCollision(c4, c2))
	assert.False(t, staticCollision(c4, c3))
	assert.True(t, staticCollision(c4, c4))
	assert.False(t, staticCollision(c4, c5))

	assert.True(t, staticCollision(c5, c1))
	assert.False(t, staticCollision(c5, c2))
	assert.False(t, staticCollision(c5, c3))
	assert.False(t, staticCollision(c5, c4))
	assert.True(t, staticCollision(c5, c5))
}

func TestDynamicCollision(t *testing.T) {
	c1 := mockCollisionBox{1, 4, 2, -3, 1, 1}
	c2 := mockCollisionBox{0, 2, 6, 0, 1, 1}
	c3 := mockCollisionBox{5, 4, 0, -3, 1, 1}

	assert.True(t, dynamicCollision(c1, c2))
	assert.False(t, dynamicCollision(c1, c3))

	assert.True(t, dynamicCollision(c2, c1))
	assert.True(t, dynamicCollision(c2, c3))

	assert.False(t, dynamicCollision(c3, c1))
	assert.True(t, dynamicCollision(c3, c2))

}
