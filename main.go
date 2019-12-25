package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 320
	windowHeight = 240
	scaleFactor  = 4

	maxBullets = 1000
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

type hitBox struct {
	x, y, xSize, ySize float64
}

var frameCounter int

// Display the square
func update(screen *ebiten.Image, p *player, e *enemy) error {

	e.update(screen)
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
	e := initEnemy()

	if err := ebiten.Run(func(screen *ebiten.Image) error { return update(screen, &p, &e) }, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}
