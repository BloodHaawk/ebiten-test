package main

import (
	"fmt"
	"image/color"
	"math"
	"os"

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
	playerBulletSize  = 3
	playerBulletSpeed = 5
	playerBulletFreq  = 60
	baseBulletSpread  = 100 //degrees
	focusBulletSpread = 90  //degrees
	bulletStreams     = 9
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
	skin          sprite
	hitBox        sprite
	bullets       []bullet
	lastShotFrame int
	isFocus       bool
}

// Bullet type
type bullet struct {
	x, y, vx, vy float64
	isOnScreen   bool
}

var bulletSkin *ebiten.Image
var bulletSprite sprite

var frameCounter int

// Display the square
func update(screen *ebiten.Image, p *player) error {
	p.move(mvtSpeed)

	// Draw the square and update the position from keyboard input
	drawSprite(screen, p.skin)

	// Show the hitBox in red when pressing Shift
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		drawSprite(screen, p.hitBox)
		p.isFocus = true
	} else {
		p.isFocus = false
	}

	// Shoot a bullet with Z key
	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		if p.isFocus {
			p.shootBullet(playerBulletFreq, bulletStreams, focusBulletSpread)
		} else {
			p.shootBullet(playerBulletFreq, bulletStreams, baseBulletSpread)
		}
	}

	for i := range p.bullets {
		if p.bullets[i].isOnScreen {
			bulletSprite.opts.GeoM.Reset()
			bulletSprite.opts.GeoM.Translate(p.bullets[i].x, p.bullets[i].y)
			drawSprite(screen, bulletSprite)
			p.bullets[i].move(playerBulletSpeed, playerBulletSize)
		}
	}

	msg := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	frameCounter++

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (p *player) move(speed float64) {
	// Use Shift to slow down
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		speed /= 2
	}
	var tx, ty float64

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		tx = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		tx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ty = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ty = -1
	}

	if r := math.Sqrt(tx*tx + ty*ty); r != 0 {
		tx = tx / r * speed
		ty = ty / r * speed

		tx = math.Max(0, p.hitBox.x()+tx) - p.hitBox.x()
		tx = math.Min(windowWidth-hitBoxSize, p.hitBox.x()+tx) - p.hitBox.x()

		ty = math.Max(0, p.hitBox.y()+ty) - p.hitBox.y()
		ty = math.Min(windowHeight-hitBoxSize, p.hitBox.y()+ty) - p.hitBox.y()
	}

	p.skin.opts.GeoM.Translate(tx, ty)
	p.hitBox.opts.GeoM.Translate(tx, ty)

	return
}

func (p *player) shootBullet(freq int, n int, spreadDeg float64) {

	if frameCounter-p.lastShotFrame >= ebiten.MaxTPS()/playerBulletFreq {
		indices := findNFirsts(p.bullets, n, func(b bullet) bool { return !b.isOnScreen })

		if len(indices) == n {
			for i := 0; i < n; i++ {
				angleDeg := -spreadDeg/2 + float64(i)*spreadDeg/float64(n-1)
				p.bullets[indices[i]].x = p.hitBox.x() + 15*math.Sin(angleDeg*math.Pi/180) + (hitBoxSize-playerBulletSize)/2
				p.bullets[indices[i]].y = p.hitBox.y() - 15*math.Cos(angleDeg*math.Pi/180) + (hitBoxSize-playerBulletSize)/2
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

func (b *bullet) move(speed, size float64) {
	b.x = b.x + speed*b.vx
	b.y = b.y + speed*b.vy
	if b.x+size < -30 || b.y+size < -30 || b.x > windowWidth+30 || b.y > windowHeight+30 {
		b.isOnScreen = false
	}
}

func drawSprite(screen *ebiten.Image, spr sprite) {
	screen.DrawImage(spr.image, &spr.opts)
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

func initPlayer() player {
	var p player
	var errH, errS, errB error
	p.hitBox.image, errH = ebiten.NewImage(hitBoxSize, hitBoxSize, ebiten.FilterNearest)
	p.skin.image, errS = ebiten.NewImage(squareSize, squareSize, ebiten.FilterNearest)
	logError(errH)
	logError(errS)

	p.hitBox.image.Fill(color.RGBA{255, 0, 0, 255})
	p.skin.image.Fill(color.White)
	p.skin.opts.GeoM.Translate((hitBoxSize-squareSize)/2, (hitBoxSize-squareSize)/2)

	// Start at middle of screen
	p.hitBox.opts.GeoM.Translate((windowWidth-hitBoxSize)/2, (windowHeight-hitBoxSize)/2)
	p.skin.opts.GeoM.Translate((windowWidth-hitBoxSize)/2, (windowHeight-hitBoxSize)/2)

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
