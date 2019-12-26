package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
)

const (
	windowWidth  = 1280
	windowHeight = 960
	scaleFactor  = 1

	maxBullets = 1000
)

// User-defined configs
var config map[string]string
var keyConfig map[string]string
var buttonConfig map[string]string

var deadZone float64

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

	p.updateBullets(screen)
	e.update(screen)
	p.update(screen)
	e.updateBullets(screen)

	printFPS(screen)

	frameCounter++

	if ebiten.IsKeyPressed(keyMap[keyConfig["quit"]]) || ebiten.IsGamepadButtonPressed(0, buttonMap[buttonConfig["quit"]]) {
		os.Exit(0)
	}

	return nil
}

func drawSprite(screen *ebiten.Image, spr sprite) {
	screen.DrawImage(spr.image, &spr.opts)
}

// Initialise Ebiten, then loop the update function
func main() {
	config = makeConfig()
	setDeadZone(config)

	keyConfig = makeKeyConfig()
	buttonConfig = makeButtonConfig()

	p := initPlayer()
	e := initEnemy()

	if err := ebiten.Run(func(screen *ebiten.Image) error { return update(screen, &p, &e) }, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}
