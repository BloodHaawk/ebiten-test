package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 320
	windowHeight = 240
	scaleFactor  = 4
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

var frameCounter int

// Display the square
func update(screen *ebiten.Image, p *player) error {

	p.update(screen)

	printFPS(screen)

	frameCounter++

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	return nil
}

func drawSprite(screen *ebiten.Image, spr sprite) {
	screen.DrawImage(spr.image, &spr.opts)
}

// Initialise Ebiten, then loop the update function
func main() {

	p := initPlayer()

	if err := ebiten.Run(func(screen *ebiten.Image) error { return update(screen, &p) }, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}
