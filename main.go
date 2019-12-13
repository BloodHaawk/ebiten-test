package main

import (
	"fmt"
	"image/color"

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
		switch {
		case spr.opts.GeoM.Element(0, 2) < windowWidth-sprSize-speed:
			matrix.tx = speed
		case spr.opts.GeoM.Element(0, 2) < windowWidth-sprSize:
			matrix.tx = windowWidth - spr.opts.GeoM.Element(0, 2) - sprSize
		default:
			matrix.tx = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		switch {
		case spr.opts.GeoM.Element(0, 2) < speed:
			matrix.tx = 0 - spr.opts.GeoM.Element(0, 2)
		case spr.opts.GeoM.Element(0, 2) > 0:
			matrix.tx = 0 - speed
		default:
			matrix.tx = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		switch {
		case spr.opts.GeoM.Element(1, 2) < windowHeight-sprSize-speed:
			matrix.ty = speed
		case spr.opts.GeoM.Element(1, 2) < windowHeight-sprSize:
			matrix.ty = windowHeight - spr.opts.GeoM.Element(1, 2) - sprSize
		default:
			matrix.ty = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		switch {
		case spr.opts.GeoM.Element(1, 2) < speed:
			matrix.ty = 0 - spr.opts.GeoM.Element(1, 2)
		case spr.opts.GeoM.Element(1, 2) > 0:
			matrix.ty = 0 - speed
		default:
			matrix.tx = 0
		}
	}
	return matrix
}

var square = sprite{}
var hitBox = sprite{}

// Display the square
func update(screen *ebiten.Image) error {

	if square.image == nil {
		square.image, _ = ebiten.NewImage(squareSize, squareSize, ebiten.FilterNearest)
	}
	if hitBox.image == nil {
		hitBox.image, _ = ebiten.NewImage(hitBoxSize, hitBoxSize, ebiten.FilterNearest)
		hitBox.opts.GeoM.Translate((squareSize-hitBoxSize)/2, (squareSize-hitBoxSize)/2)
	}

	// Draw the square and update the position from keyboard input
	square.image.Fill(color.White)
	trans := square.moveSprite(squareSize, mvtSpeed)
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
