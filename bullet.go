package main

// Bullet type
type bullet struct {
	x, y, vx, vy float64
	isOnScreen   bool
	xSize, ySize int
}

// Implement collisionBox

func (b *bullet) PosX() float64 {
	return b.x
}
func (b *bullet) PosY() float64 {
	return b.y
}
func (b *bullet) VX() float64 {
	return b.vx
}
func (b *bullet) VY() float64 {
	return b.vy
}
func (b *bullet) SizeX() float64 {
	return float64(b.xSize)
}
func (b *bullet) SizeY() float64 {
	return float64(b.ySize)
}

func (b *bullet) centreX() float64 {
	return b.x + float64(b.xSize)/2
}
func (b *bullet) centreY() float64 {
	return b.y + float64(b.ySize)/2
}

func (b *bullet) move(speed, size float64) {
	b.x += speed * b.vx
	b.y += speed * b.vy
	if b.x+size < -30 || b.y+size < -30 || b.x > windowWidth+30 || b.y > windowHeight+30 {
		b.isOnScreen = false
	}
}

func findNFirsts(bullets []bullet, n int, f func(bullet) bool) []int {
	indices := make([]int, 0, n)
	for i := range bullets {
		if f(bullets[i]) {
			indices = append(indices, i)
		}
		if len(indices) == n {
			break
		}
	}
	return indices
}
