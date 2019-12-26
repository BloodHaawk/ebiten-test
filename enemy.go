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

func (e *enemy) update(screen *ebiten.Image) {
	e.move(e.mvtSpeed)

	// Draw the square and update the position from keyboard input
	drawSprite(screen, e.skin)

	// Shoot a bullett
	if ebiten.IsKeyPressed(keyMap[keyConfig["bossShoot"]]) {
		e.shootBullet(e.bulletFreq, e.bulletStreams, e.bulletSpread)
	}

	for i := range e.bullets {
		if e.bullets[i].isOnScreen {
			e.bulletSprite.opts.GeoM.Reset()
			e.bulletSprite.opts.GeoM.Translate(e.bullets[i].x, e.bullets[i].y)
			drawSprite(screen, e.bulletSprite)
			e.bullets[i].move(e.bulletSpeed, float64(e.bulletSize))
		}
	}
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (e *enemy) move(speed float64) {
	var tx, ty float64

	if ebiten.IsKeyPressed(keyMap[keyConfig["bossRight"]]) {
		tx = 1
	}
	if ebiten.IsKeyPressed(keyMap[keyConfig["bossLeft"]]) {
		tx = -1
	}
	if ebiten.IsKeyPressed(keyMap[keyConfig["bossDown"]]) {
		ty = 1
	}
	if ebiten.IsKeyPressed(keyMap[keyConfig["bossUp"]]) {
		ty = -1
	}

	if r := math.Sqrt(tx*tx + ty*ty); r != 0 {
		tx = tx / r * speed
		ty = ty / r * speed

		tx = math.Max(0, e.hitBox.x+tx) - e.hitBox.x
		tx = math.Min(windowWidth-e.hitBox.xSize, e.hitBox.x+tx) - e.hitBox.x

		ty = math.Max(0, e.hitBox.y+ty) - e.hitBox.y
		ty = math.Min(windowHeight-e.hitBox.ySize, e.hitBox.y+ty) - e.hitBox.y
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

func initEnemy() enemy {
	var e enemy

	e.hitBox.xSize = 60
	e.hitBox.ySize = 60
	e.skinSize = 60
	e.mvtSpeed = 4
	e.bulletSize = 20
	e.bulletSpeed = 8
	e.bulletFreq = 5
	e.bulletSpread = 270 //degrees
	e.bulletStreams = 30
	e.bulletSpawnOffset = 80
	e.color = color.RGBA{0, 255, 0, 255}
	e.bulletColor = color.RGBA{210, 30, 210, 255}

	var errH, errS, errB error
	e.skin.image, errS = ebiten.NewImage(e.skinSize, e.skinSize, ebiten.FilterNearest)
	logError(errH)
	logError(errS)

	e.skin.image.Fill(e.color)
	e.skin.opts.GeoM.Translate((e.hitBox.xSize-float64(e.skinSize))/2, (e.hitBox.ySize-float64(e.skinSize))/2)

	// Start at middle of screen
	e.hitBox.x = (windowWidth - e.hitBox.xSize) / 2
	e.skin.opts.GeoM.Translate((windowWidth-e.hitBox.xSize)/2, 0)

	e.bulletSkin, errB = ebiten.NewImage(e.bulletSize, e.bulletSize, ebiten.FilterNearest)
	logError(errB)
	e.bulletSkin.Fill(e.bulletColor)
	e.bulletSprite = sprite{e.bulletSkin, ebiten.DrawImageOptions{}}

	e.bullets = make([]bullet, maxBullets)

	return e
}
