package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockCollisionBox struct {
	x, y, xSize, ySize float64
}

func (c mockCollisionBox) posX() float64 {
	return c.x
}
func (c mockCollisionBox) posY() float64 {
	return c.y
}
func (c mockCollisionBox) sizeX() float64 {
	return c.xSize
}
func (c mockCollisionBox) sizeY() float64 {
	return c.ySize
}

func TestCollision(t *testing.T) {
	c1 := mockCollisionBox{0, 0, 10, 10}
	c2 := mockCollisionBox{-5, -5, 6, 7}
	c3 := mockCollisionBox{8, -5, 4, 6}
	c4 := mockCollisionBox{8, 8, 3, 5}
	c5 := mockCollisionBox{-3, 7, 4, 9}

	assert.True(t, collision(c1, c1))
	assert.True(t, collision(c1, c2))
	assert.True(t, collision(c1, c3))
	assert.True(t, collision(c1, c4))
	assert.True(t, collision(c1, c5))

	assert.True(t, collision(c2, c1))
	assert.True(t, collision(c2, c2))
	assert.False(t, collision(c2, c3))
	assert.False(t, collision(c2, c4))
	assert.False(t, collision(c2, c5))

	assert.True(t, collision(c3, c1))
	assert.False(t, collision(c3, c2))
	assert.True(t, collision(c3, c3))
	assert.False(t, collision(c3, c4))
	assert.False(t, collision(c3, c5))

	assert.True(t, collision(c4, c1))
	assert.False(t, collision(c4, c2))
	assert.False(t, collision(c4, c3))
	assert.True(t, collision(c4, c4))
	assert.False(t, collision(c4, c5))

	assert.True(t, collision(c5, c1))
	assert.False(t, collision(c5, c2))
	assert.False(t, collision(c5, c3))
	assert.False(t, collision(c5, c4))
	assert.True(t, collision(c5, c5))

}
