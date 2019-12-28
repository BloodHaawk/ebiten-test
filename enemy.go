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
	pattern       pattern
	lastShotFrame int

	vx, vy float64

	skinSize int
	mvtSpeed float64

	color color.Color
}

// Implement collisionBox interface

func (e *enemy) posX() float64 {
	return e.hitBox.x
}
func (e *enemy) posY() float64 {
	return e.hitBox.y
}
func (e *enemy) vX() float64 {
	return e.vx
}
func (e *enemy) vY() float64 {
	return e.vy
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

	e.vx = tx
	e.vy = ty

	return
}

func (e *enemy) shootBullets() {
	if frameCounter-e.lastShotFrame >= ebiten.MaxTPS()/e.pattern.opts.bulletFreq {
		e.pattern.spawn(e.hitBox)
		e.lastShotFrame = frameCounter
	}
}

func initEnemy(
	x float64,
	y float64,
	hitBoxSizeX float64,
	hitBoxSizeY float64,
	mvtSpeed float64,
	skinSize int,
	isAimed bool,
	clr color.Color,
	pattOpts patternOpts) enemy {

	skinImage, errS := ebiten.NewImage(skinSize, skinSize, ebiten.FilterNearest)
	logError(errS)
	skinImage.Fill(clr)
	skinOpts := ebiten.DrawImageOptions{}
	skinOpts.GeoM.Translate((hitBoxSizeX-float64(skinSize))/2, (hitBoxSizeY-float64(skinSize))/2)
	skinOpts.GeoM.Translate(x, y)

	skin := sprite{
		skinImage,
		skinOpts,
	}

	hitBox := hitBox{x, y, hitBoxSizeX, hitBoxSizeY}

	pattern := initPattern(isAimed, pattOpts)

	return enemy{
		skin,
		hitBox,
		pattern,
		0, // lastShotFrame

		0, 0, // Initialise at 0 speed

		skinSize,
		mvtSpeed,

		clr,
	}
}
