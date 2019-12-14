package main

import (
	"fmt"
	"image/color"
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

type translationVector struct {
	tx, ty float64
}

// Basic sprite with options
type sprite struct {
	image *ebiten.Image
	opts  ebiten.DrawImageOptions
}

// Move a sprite from keyboard inputs (use Shift to slow down)
func (spr sprite) moveSprite(sprSize, speed float64) translationVector {
	matrix := translationVector{}

	// Use Shift to slow down
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		speed = speed / 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		matrix.tx = math.Min(windowWidth-sprSize-spr.opts.GeoM.Element(0, 2), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		matrix.tx = -math.Min(spr.opts.GeoM.Element(0, 2), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		matrix.ty = math.Min(windowWidth-sprSize-spr.opts.GeoM.Element(1, 2), speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		matrix.ty = -math.Min(spr.opts.GeoM.Element(1, 2), speed)
	}
	return matrix
}

var square = sprite{}
var hitBox = sprite{}

// Display the square
func update(screen *ebiten.Image) error {

	if hitBox.image == nil {
		hitBox.image, _ = ebiten.NewImage(hitBoxSize, hitBoxSize, ebiten.FilterNearest)
	}
	if square.image == nil {
		square.image, _ = ebiten.NewImage(squareSize, squareSize, ebiten.FilterNearest)
		square.opts.GeoM.Translate((hitBoxSize-squareSize)/2, (hitBoxSize-squareSize)/2)
	}

	trans := hitBox.moveSprite(hitBoxSize, mvtSpeed)

	// Draw the square and update the position from keyboard input
	square.image.Fill(color.White)
	square.opts.GeoM.Translate(trans.tx, trans.ty)
	screen.DrawImage(square.image, &square.opts)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %f, Y: %f", square.opts.GeoM.Element(0, 2), square.opts.GeoM.Element(1, 2)), 40, 40)

	// Show the hitBox in red when pressing Shitf
	hitBox.image.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	hitBox.opts.GeoM.Translate(trans.tx, trans.ty)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %f, Y: %f", hitBox.opts.GeoM.Element(0, 2), hitBox.opts.GeoM.Element(1, 2)), 40, 80)

	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		screen.DrawImage(hitBox.image, &hitBox.opts)
	}

	return nil
}

// Initialise Ebiten, then loop the update function
func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}
