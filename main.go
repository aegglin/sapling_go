package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 960
	screenHeight = 540
)

var (
	leftSprite1  *ebiten.Image
	rightSprite1 *ebiten.Image
	upSprite1    *ebiten.Image
	downSprite1  *ebiten.Image
	// leftSprite2     *ebiten.Image
	// rightSprite2    *ebiten.Image
	// upSprite2       *ebiten.Image
	// downSprite2     *ebiten.Image
	// backgroundImage *ebiten.Image
)

func init() {
	var err error
	leftSprite1, _, err = ebitenutil.NewImageFromFile("assets/BeetleLeft1.png")
	if err != nil {
		log.Fatal(err)
	}
	rightSprite1, _, err = ebitenutil.NewImageFromFile("assets/BeetleRight1.png")
	if err != nil {
		log.Fatal(err)
	}
	upSprite1, _, err = ebitenutil.NewImageFromFile("assets/BeetleUp1.png")
	if err != nil {
		log.Fatal(err)
	}
	downSprite1, _, err = ebitenutil.NewImageFromFile("assets/BeetleDown1.png")
	if err != nil {
		log.Fatal(err)
	}

	// leftSprite2, _, err = ebitenutil.NewImageFromFile("assets/BeetleLeft2.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rightSprite2, _, err = ebitenutil.NewImageFromFile("assets/BeetleRight2.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// upSprite2, _, err = ebitenutil.NewImageFromFile("assets/BeetleUp2.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// downSprite2, _, err = ebitenutil.NewImageFromFile("assets/BeetleDown2.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// the character has x, y, vx, and vy
type character struct {
	// x             int
	// y             int
	// vx            int
	// vy            int
	direction     Direction
	currentSprite *ebiten.Image
}

// thte game has the main character beetle
type Game struct {
	beetle *character
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.beetle.direction = Left
		g.beetle.currentSprite = leftSprite1
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.beetle.direction = Right
		g.beetle.currentSprite = rightSprite1
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.beetle.direction = Up
		g.beetle.currentSprite = upSprite1
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.beetle.direction = Down
		g.beetle.currentSprite = downSprite1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(g.beetle.currentSprite, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return outsideWidth, outsideHeight
	// return 320, 240
	return 320, 240
}

func main() {
	// g := &Game{}

	// err := ebiten.RunGame(g)
	// if err != nil {
	// 	panic(err)
	// }
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Render an image")
	g := Game{}
	beetle := character{}
	g.beetle = &beetle
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
