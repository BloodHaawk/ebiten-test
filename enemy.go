package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// Enemy struct
type enemy struct {
	skin          sprite
	hitBox        hitBox
	bullets       []bullet
	bulletSkin    *ebiten.Image
	bulletSprite  sprite
	lastShotFrame int

	skinSize int
	mvtSpeed float64

	bulletSize        int
	bulletSpeed       float64
	bulletFreq        int
	bulletSpread      float64
	bulletStreams     int
	bulletSpawnOffset float64

	color       color.Color
	bulletColor color.Color
}

// Implement collisionBox interface

func (e *enemy) posX() float64 {
	return e.hitBox.x
}
func (e *enemy) posY() float64 {
	return e.hitBox.y
}
func (e *enemy) sizeX() float64 {
	return float64(e.hitBox.xSize)
}
func (e *enemy) sizeY() float64 {
	return float64(e.hitBox.ySize)
}

func (e *enemy) update(screen *ebiten.Image) {
	e.move(e.mvtSpeed)

	// Draw the square and update the position from keyboard input
	drawSprite(screen, e.skin)
}

func (e *enemy) updateBullets(screen *ebiten.Image, p *player) {
	e.shootBullet(e.bulletFreq, e.bulletStreams, e.bulletSpread)

	for i := range e.bullets {
		if e.bullets[i].isOnScreen {
			e.bulletSprite.opts.GeoM.Reset()
			e.bulletSprite.opts.GeoM.Translate(e.bullets[i].x, e.bullets[i].y)
			drawSprite(screen, e.bulletSprite)
			e.bullets[i].move(e.bulletSpeed, float64(e.bulletSize))
			if collision(p, &e.bullets[i]) {
				e.bullets[i].isOnScreen = false
				screen.Fill(color.RGBA{255, 0, 0, 255})
			}
		}
	}
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (e *enemy) move(speed float64) {
	var tx, ty float64

	if (frameCounter+60)%240 < 120 {
		tx = 1
	} else {
		tx = -1
	}
	if (frameCounter+30)%120 < 60 {
		ty = 1
	} else {
		ty = -1
	}

	if r := math.Sqrt(tx*tx + ty*ty); r != 0 {
		tx = tx / r * speed
		ty = ty / r * speed
	}

	e.skin.opts.GeoM.Translate(tx, ty)
	e.hitBox.x += tx
	e.hitBox.y += ty

	return
}

func (e *enemy) shootBullet(freq int, n int, spreadDeg float64) {

	if frameCounter-e.lastShotFrame >= ebiten.MaxTPS()/e.bulletFreq {
		indices := findNFirsts(e.bullets, n, func(b bullet) bool { return !b.isOnScreen })

		if len(indices) == n {
			for i := 0; i < n; i++ {
				angleDeg := -spreadDeg/2 + float64(i)*spreadDeg/float64(n-1)
				e.bullets[indices[i]].x = e.hitBox.x + e.bulletSpawnOffset*math.Sin(angleDeg*math.Pi/180) + (e.hitBox.xSize-float64(e.bulletSize))/2
				e.bullets[indices[i]].y = e.hitBox.y + e.bulletSpawnOffset*math.Cos(angleDeg*math.Pi/180) + (e.hitBox.ySize-float64(e.bulletSize))/2
				e.bullets[indices[i]].vx = math.Sin(angleDeg * math.Pi / 180)
				e.bullets[indices[i]].vy = math.Cos(angleDeg * math.Pi / 180)
				e.bullets[indices[i]].isOnScreen = true
			}
		}
		e.lastShotFrame = frameCounter
	}

}

func initEnemy(x, y float64) enemy {
	var e enemy

	e.hitBox.xSize = 60
	e.hitBox.ySize = 60
	e.skinSize = 60
	e.mvtSpeed = 2
	e.bulletSize = 20
	e.bulletSpeed = 4
	e.bulletFreq = 4
	e.bulletSpread = 270 //degrees
	e.bulletStreams = 25
	e.bulletSpawnOffset = 80
	e.color = color.RGBA{0, 255, 0, 255}
	e.bulletColor = color.RGBA{210, 30, 210, 255}

	var errH, errS, errB error
	e.skin.image, errS = ebiten.NewImage(e.skinSize, e.skinSize, ebiten.FilterNearest)
	logError(errH)
	logError(errS)

	e.skin.image.Fill(e.color)
	e.skin.opts.GeoM.Translate((e.hitBox.xSize-float64(e.skinSize))/2, (e.hitBox.ySize-float64(e.skinSize))/2)

	e.hitBox.x = x
	e.hitBox.y = y
	e.skin.opts.GeoM.Translate(x, y)

	e.bulletSkin, errB = ebiten.NewImage(e.bulletSize, e.bulletSize, ebiten.FilterNearest)
	logError(errB)
	e.bulletSkin.Fill(e.bulletColor)
	e.bulletSprite = sprite{e.bulletSkin, ebiten.DrawImageOptions{}}

	e.bullets = make([]bullet, maxBullets)

	return e
}
