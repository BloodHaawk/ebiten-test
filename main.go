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

	squareSize = 16
	mvtSpeed   = 7
)

type translationVector struct {
	tx, ty float64
}

var square *ebiten.Image
var opts = &ebiten.DrawImageOptions{}

// Display the square
func update(screen *ebiten.Image) error {

	if square == nil {
		square, _ = ebiten.NewImage(squareSize, squareSize, ebiten.FilterNearest)
	}

	square.Fill(color.White)

	opts.GeoM.Translate(move().tx, move().ty)

	screen.DrawImage(square, opts)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %f, Y: %f", opts.GeoM.Element(0, 2), opts.GeoM.Element(1, 2)), 20, 20)

	return nil
}

// Move the square
func move() translationVector {
	matrix := translationVector{}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		switch {
		case opts.GeoM.Element(0, 2) < windowWidth-squareSize-mvtSpeed:
			matrix.tx = matrix.tx + mvtSpeed
		case opts.GeoM.Element(0, 2) < windowWidth-squareSize:
			matrix.tx = windowWidth - opts.GeoM.Element(0, 2) - squareSize
		default:
			matrix.tx = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		switch {
		case opts.GeoM.Element(0, 2) < mvtSpeed:
			matrix.tx = 0 - opts.GeoM.Element(0, 2)
		case opts.GeoM.Element(0, 2) > 0:
			matrix.tx = matrix.tx - mvtSpeed
		default:
			matrix.tx = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		switch {
		case opts.GeoM.Element(1, 2) < windowHeight-squareSize-mvtSpeed:
			matrix.ty = matrix.ty + mvtSpeed
		case opts.GeoM.Element(1, 2) < windowHeight-squareSize:
			matrix.ty = windowHeight - opts.GeoM.Element(1, 2) - squareSize
		default:
			matrix.ty = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		switch {
		case opts.GeoM.Element(1, 2) < mvtSpeed:
			matrix.ty = 0 - opts.GeoM.Element(1, 2)
		case opts.GeoM.Element(1, 2) > 0:
			matrix.ty = matrix.ty - mvtSpeed
		default:
			matrix.tx = 0
		}
	}
	return matrix
}

// Initialise Ebiten, then loop the update function
func main() {
	if err := ebiten.Run(update, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}
