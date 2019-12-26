package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// Player struct
type player struct {
	skin          sprite
	hitBox        sprite
	bullets       []bullet
	bulletSkin    *ebiten.Image
	bulletSprite  sprite
	lastShotFrame int
	isFocus       bool

	skinSize   int
	hitBoxSize int
	mvtSpeed   float64

	bulletSize        int
	bulletSpeed       float64
	bulletFreq        int
	baseBulletSpread  float64 //degrees
	focusBulletSpread float64 //degrees
	bulletStreams     int
	bulletSpawnOffset float64
}

func (p *player) update(screen *ebiten.Image) {
	p.move(p.mvtSpeed)

	// Draw the square and update the position from keyboard input
	drawSprite(screen, p.skin)

	// Show the hitBox in red when pressing focus
	if ebiten.IsKeyPressed(keyMap[keyConfig["focus"]]) {
		drawSprite(screen, p.hitBox)
		p.isFocus = true
	} else {
		p.isFocus = false
	}

	// Shoot a bullet
	if ebiten.IsKeyPressed(keyMap[keyConfig["shoot"]]) {
		if p.isFocus {
			p.shootBullet(p.bulletFreq, p.bulletStreams, p.focusBulletSpread)
		} else {
			p.shootBullet(p.bulletFreq, p.bulletStreams, p.baseBulletSpread)
		}
	}

	for i := range p.bullets {
		if p.bullets[i].isOnScreen {
			p.bulletSprite.opts.GeoM.Reset()
			p.bulletSprite.opts.GeoM.Translate(p.bullets[i].x, p.bullets[i].y)
			drawSprite(screen, p.bulletSprite)
			p.bullets[i].move(p.bulletSpeed, float64(p.bulletSize))
		}
	}
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (p *player) move(speed float64) {
	// Use Shift to slow down
	if ebiten.IsKeyPressed(keyMap[keyConfig["focus"]]) {
		speed /= 2
	}
	var tx, ty float64

	if ebiten.IsKeyPressed(keyMap[keyConfig["right"]]) {
		tx = 1
	}
	if ebiten.IsKeyPressed(keyMap[keyConfig["left"]]) {
		tx = -1
	}
	if ebiten.IsKeyPressed(keyMap[keyConfig["down"]]) {
		ty = 1
	}
	if ebiten.IsKeyPressed(keyMap[keyConfig["up"]]) {
		ty = -1
	}

	if r := math.Sqrt(tx*tx + ty*ty); r != 0 {
		tx = tx / r * speed
		ty = ty / r * speed

		tx = math.Max(0, p.hitBox.x()+tx) - p.hitBox.x()
		tx = math.Min(windowWidth-float64(p.hitBoxSize), p.hitBox.x()+tx) - p.hitBox.x()

		ty = math.Max(0, p.hitBox.y()+ty) - p.hitBox.y()
		ty = math.Min(windowHeight-float64(p.hitBoxSize), p.hitBox.y()+ty) - p.hitBox.y()
	}

	p.skin.opts.GeoM.Translate(tx, ty)
	p.hitBox.opts.GeoM.Translate(tx, ty)

	return
}

func (p *player) shootBullet(freq int, n int, spreadDeg float64) {

	if frameCounter-p.lastShotFrame >= ebiten.MaxTPS()/p.bulletFreq {
		indices := findNFirsts(p.bullets, n, func(b bullet) bool { return !b.isOnScreen })

		if len(indices) == n {
			for i := 0; i < n; i++ {
				angleDeg := -spreadDeg/2 + float64(i)*spreadDeg/float64(n-1)
				p.bullets[indices[i]].x = p.hitBox.x() + p.bulletSpawnOffset*math.Sin(angleDeg*math.Pi/180) + float64(p.hitBoxSize-p.bulletSize)/2
				p.bullets[indices[i]].y = p.hitBox.y() - p.bulletSpawnOffset*math.Cos(angleDeg*math.Pi/180) + float64(p.hitBoxSize-p.bulletSize)/2
				if p.isFocus {
					p.bullets[indices[i]].vx = 0
					p.bullets[indices[i]].vy = -1
				} else {
					p.bullets[indices[i]].vx = math.Sin(angleDeg * math.Pi / 180)
					p.bullets[indices[i]].vy = -math.Cos(angleDeg * math.Pi / 180)
				}
				p.bullets[indices[i]].isOnScreen = true
			}
		}
		p.lastShotFrame = frameCounter
	}

}

func initPlayer() player {
	var p player

	p.skinSize = 32
	p.hitBoxSize = 8
	p.mvtSpeed = 12
	p.bulletSize = 12
	p.bulletSpeed = 20
	p.bulletFreq = 60
	p.baseBulletSpread = 100 //degrees
	p.focusBulletSpread = 90 //degrees
	p.bulletStreams = 9
	p.bulletSpawnOffset = 60

	var errH, errS, errB error
	p.hitBox.image, errH = ebiten.NewImage(p.hitBoxSize, p.hitBoxSize, ebiten.FilterNearest)
	p.skin.image, errS = ebiten.NewImage(p.skinSize, p.skinSize, ebiten.FilterNearest)
	logError(errH)
	logError(errS)

	p.hitBox.image.Fill(color.RGBA{255, 0, 0, 255})
	p.skin.image.Fill(color.White)
	p.skin.opts.GeoM.Translate(float64(p.hitBoxSize-p.skinSize)/2, float64(p.hitBoxSize-p.skinSize)/2)

	// Start at middle of screen
	p.hitBox.opts.GeoM.Translate((windowWidth-float64(p.hitBoxSize))/2, (windowHeight-float64(p.hitBoxSize))/2)
	p.skin.opts.GeoM.Translate((windowWidth-float64(p.hitBoxSize))/2, (windowHeight-float64(p.hitBoxSize))/2)

	p.bulletSkin, errB = ebiten.NewImage(p.bulletSize, p.bulletSize, ebiten.FilterNearest)
	logError(errB)
	p.bulletSkin.Fill(color.White)
	p.bulletSprite = sprite{p.bulletSkin, ebiten.DrawImageOptions{}}

	p.bullets = make([]bullet, maxBullets)

	return p
}
