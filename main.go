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
)

// Basic sprite with options
type sprite struct {
	image *ebiten.Image
	opts  ebiten.DrawImageOptions
}

// Player struct
type player struct {
	skin   sprite
	hitBox sprite
}

// Display the square
func update(screen *ebiten.Image, p *player) error {

	tx, ty := p.hitBox.moveSprite(hitBoxSize, mvtSpeed)
	p.skin.opts.GeoM.Translate(tx, ty)
	p.hitBox.opts.GeoM.Translate(tx, ty)

	// Draw the square and update the position from keyboard input
	drawSprite(screen, p.skin, color.White)

	// Show the hitBox in red when pressing Shift
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		drawSprite(screen, p.hitBox, color.RGBA{0xff, 0x00, 0x00, 0xff})
	}

	debugPlayer(screen, p)

	return nil
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (spr sprite) moveSprite(sprSize, speed float64) (tx, ty float64) {
	// Use Shift to slow down
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		speed = speed / 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		tx = math.Min(windowWidth-sprSize-spr.opts.GeoM.Element(0, 2), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		tx = -math.Min(spr.opts.GeoM.Element(0, 2), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ty = math.Min(windowWidth-sprSize-spr.opts.GeoM.Element(1, 2), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ty = -math.Min(spr.opts.GeoM.Element(1, 2), speed)
	}
	return
}

func drawSprite(screen *ebiten.Image, spr sprite, clr color.Color) {
	spr.image.Fill(clr)
	screen.DrawImage(spr.image, &spr.opts)
}

func debugPlayer(screen *ebiten.Image, p *player) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %f, Y: %f", p.skin.opts.GeoM.Element(0, 2), p.skin.opts.GeoM.Element(1, 2)), 40, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %f, Y: %f", p.hitBox.opts.GeoM.Element(0, 2), p.hitBox.opts.GeoM.Element(1, 2)), 40, 80)
}

func initPlayer() player {
	var p player
	var errH, errS error
	p.hitBox.image, errH = ebiten.NewImage(hitBoxSize, hitBoxSize, ebiten.FilterNearest)
	p.skin.image, errS = ebiten.NewImage(squareSize, squareSize, ebiten.FilterNearest)
	logError(errH)
	logError(errS)
	p.skin.opts.GeoM.Translate((hitBoxSize-squareSize)/2, (hitBoxSize-squareSize)/2)

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
