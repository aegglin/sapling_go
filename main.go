package main

// import "github.com/hajimehoshi/ebiten/v2"
import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 960
	screenHeight = 540
)

var (
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	upSprite        *ebiten.Image
	downSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	// var err error

	img, _, err := image.Decode(bytes.NewReader("BeetleRight1_png"))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader("BeetleLeft1_png"))
	if err != nil {
		panic(err)
	}
	leftSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader("BeetleDown1_png"))
	if err != nil {
		panic(err)
	}
	downSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader("BeetleRight1_png"))
	if err != nil {
		panic(err)
	}
	rightSprite = ebiten.NewImageFromImage(img)
}

// the character has x, y, vx, and vy
type character struct {
	x  int
	y  int
	vx int
	vy int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// thte game has the main character beetle
type Game struct {
	beetle *character
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {

	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {

	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {

	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {

	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return outsideWidth, outsideHeight
	// return 320, 240
	return screenWidth, screenHeight
}

func main() {
	// g := &Game{}

	// err := ebiten.RunGame(g)
	// if err != nil {
	// 	panic(err)
	// }
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
