package main

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"

	"github.com/bloodhaawk/shmup-1/controlmap"
	"github.com/bloodhaawk/shmup-1/utils"
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

var gamepadID int
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

func (hb *hitBox) centreX() float64 {
	return hb.x + hb.xSize/2
}
func (hb *hitBox) centreY() float64 {
	return hb.y + hb.ySize/2
}

var frameCounter int

// Display the square
func update(screen *ebiten.Image, p *player, e []enemy) error {
	p.updateBullets(screen, e)
	for i := range e {
		e[i].update(screen)
	}
	p.update(screen)
	for i := range e {
		e[i].shootBullets()
		e[i].pattern.updateBullets(screen, e[i].hitBox, p)
	}

	utils.PrintFPS(screen)

	frameCounter++

	if ebiten.IsKeyPressed(controlmap.KeyMap[keyConfig["quit"]]) || ebiten.IsGamepadButtonPressed(gamepadID, controlmap.ButtonMap[buttonConfig["quit"]]) {
		os.Exit(0)
	}

	return nil
}

func drawSprite(screen *ebiten.Image, spr sprite) {
	screen.DrawImage(spr.image, &spr.opts)
}

// Initialise Ebiten, then loop the update function
func main() {
	config = utils.MakeConfig()
	gamepadID = utils.ReadGamepadID(config)
	deadZone = utils.ReadDeadZone(config)

	keyConfig = controlmap.MakeKeyConfig()
	buttonConfig = controlmap.MakeButtonConfig()

	p := initPlayer()

	var (
		bulletSize        int         = 20
		bulletSpeed       float64     = 6
		bulletFreq        int         = 10
		bulletSpread      float64     = 270 //degrees
		bulletStreams     int         = 25
		bulletSpawnOffset float64     = 80
		bulletColor       color.Color = color.RGBA{210, 30, 210, 255}
	)

	bulletSkin, errB := ebiten.NewImage(bulletSize, bulletSize, ebiten.FilterNearest)
	utils.LogError(errB)
	bulletSkin.Fill(bulletColor)
	bulletSprite := sprite{bulletSkin, ebiten.DrawImageOptions{}}

	pattOps1 := patternOpts{
		bulletSkin,
		bulletSprite,

		bulletSize,
		bulletSpeed,
		bulletFreq,
		bulletSpread,
		bulletStreams,
		bulletSpawnOffset,

		bulletColor,
	}

	e1 := initEnemy(
		windowWidth/2,              // Spawn point x-coord
		150,                        // Spawn point y-coord
		60,                         // Hitbox x-size
		60,                         // Hitbox y-size
		2,                          // Movement speed
		60,                         // Skin size
		true,                       // isAimed
		color.RGBA{0, 255, 0, 255}, // Skin color
		pattOps1,                   // Patterns options
	)
	e2 := initEnemy(
		2*windowWidth/3,            // Spawn point x-coord
		150,                        // Spawn point y-coord
		60,                         // Hitbox x-size
		60,                         // Hitbox y-size
		2,                          // Movement speed
		60,                         // Skin size
		true,                       // isAimed
		color.RGBA{0, 255, 0, 255}, // Skin color
		pattOps1,                   // Patterns options
	)
	e := []enemy{e1, e2}

	if err := ebiten.Run(func(screen *ebiten.Image) error { return update(screen, &p, e) }, windowWidth, windowHeight, scaleFactor, "Hello, world!"); err != nil {
		panic(err)
	}
}
