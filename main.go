package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	windowWidth  = 320
	windowHeight = 240
	scaleFactor  = 4

	squareSize = 8
	hitBoxSize = 2
	mvtSpeed   = 3

	maxPlayerBullets  = 1000
	playerBulletSize  = 2
	playerBulletSpeed = 5
)

// Basic sprite with options
type sprite struct {
	image *ebiten.Image
	opts  ebiten.DrawImageOptions
}

func (spr sprite) x() float64 {
	return spr.opts.GeoM.Element(0, 2)
}
func (spr sprite) y() float64 {
	return spr.opts.GeoM.Element(1, 2)
}

// Player struct
type player struct {
	skin    sprite
	hitBox  sprite
	bullets []bullet
}

// Bullet type
type bullet struct {
	x, y       float64
	isOnScreen bool
}

var bulletSkin *ebiten.Image
var bulletSprite sprite

// Display the square
func update(screen *ebiten.Image, p *player) error {

	p.move(mvtSpeed)

	// Draw the square and update the position from keyboard input
	drawSprite(screen, p.skin)

	// Show the hitBox in red when pressing Shift
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		drawSprite(screen, p.hitBox)
	}

	// Shoot a bullet with SpaceBar
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootBullet(screen)
	}

	for i := range p.bullets {
		if p.bullets[i].isOnScreen {
			p.bullets[i].move(playerBulletSpeed)
			bulletSprite.opts.GeoM.Reset()
			bulletSprite.opts.GeoM.Translate(p.bullets[i].x, p.bullets[i].y)
			drawSprite(screen, bulletSprite)
		}
	}

	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f
`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	return nil
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (p *player) move(speed float64) {
	// Use Shift to slow down
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		speed = speed / 2
	}
	var tx, ty float64

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		tx = math.Min(windowWidth-hitBoxSize-p.hitBox.x(), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		tx = -math.Min(p.hitBox.x(), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ty = math.Min(windowHeight-hitBoxSize-p.hitBox.y(), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ty = -math.Min(p.hitBox.y(), speed)
	}

	p.skin.opts.GeoM.Translate(tx, ty)
	p.hitBox.opts.GeoM.Translate(tx, ty)

	return
}

func (p *player) shootBullet(screen *ebiten.Image) {
	for i := range p.bullets {
		if !p.bullets[i].isOnScreen {
			p.bullets[i].x = p.hitBox.x()
			p.bullets[i].y = p.hitBox.y() - 10
			p.bullets[i].isOnScreen = true
			break
		}
	}
}

func (b *bullet) move(speed float64) {
	b.y = b.y - speed
	if b.y < 0 {
		b.isOnScreen = false
	}
}

func drawSprite(screen *ebiten.Image, spr sprite) {
	screen.DrawImage(spr.image, &spr.opts)
}

func initPlayer() player {
	var p player
	var errH, errS, errB error
	p.hitBox.image, errH = ebiten.NewImage(hitBoxSize, hitBoxSize, ebiten.FilterNearest)
	p.skin.image, errS = ebiten.NewImage(squareSize, squareSize, ebiten.FilterNearest)
	logError(errH)
	logError(errS)
	p.hitBox.image.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	p.skin.image.Fill(color.White)
	p.skin.opts.GeoM.Translate((hitBoxSize-squareSize)/2, (hitBoxSize-squareSize)/2)

	bulletSkin, errB = ebiten.NewImage(playerBulletSize, playerBulletSize, ebiten.FilterNearest)
	logError(errB)
	bulletSkin.Fill(color.White)
	bulletSprite = sprite{bulletSkin, ebiten.DrawImageOptions{}}

	p.bullets = make([]bullet, maxPlayerBullets)

	return p
}

// Initialise Ebiten, then loop the update function
func main() {

	p := initPlayer()

	if err := ebiten.Run(func(screen *ebiten.Image) error { return update(screen, &p) }, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
