package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	rawPixelSize    = 16
	scale           = 3
	maxScreenCol    = 16
	maxScreenRow    = 12
	fps             = 60
	imageDimension  = 512
	tileSize        = rawPixelSize * scale
	gameWidth       = tileSize * maxScreenCol
	gameHeight      = tileSize * maxScreenRow
	imageScale      = float64(tileSize) / imageDimension
	numWorldColumns = 50
	numWorldRows    = 50
)

var (
	upSprite1    *ebiten.Image
	downSprite1  *ebiten.Image
	rightSprite1 *ebiten.Image
	leftSprite1  *ebiten.Image
	upSprite2    *ebiten.Image
	downSprite2  *ebiten.Image
	rightSprite2 *ebiten.Image
	leftSprite2  *ebiten.Image
)

func init() {}

var (
	grass           *ebiten.Image
	tree1           *ebiten.Image
	tree2           *ebiten.Image
	tree3           *ebiten.Image
	shrub           *ebiten.Image
	underbrush      *ebiten.Image
	shrubUnderbrush *ebiten.Image
	orangeFlower    *ebiten.Image
	treeFlies1      *ebiten.Image
	treeBeehive1    *ebiten.Image
	treeWoodpecker1 *ebiten.Image
)

// the game has the main character beetle and the tiles
type Game struct {
	beetle         *Beetle
	mapTileHandler *MapTileHandler
}

func (g *Game) Update() error {
	g.beetle.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	fmt.Println("Starting drawing")
	screen.Fill(color.White)
	g.mapTileHandler.DrawAll(screen)
	g.beetle.Draw(screen)
	fmt.Println("Done drawing")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameWidth, gameHeight
}

func main() {

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Sapling by Aiden Egglin")

	fmt.Println("creating map handler")
	mapTileHandler := &MapTileHandler{}
	mapTileHandler.LoadMap()
	mapTileHandler.LoadTileImages()
	fmt.Println("Done creating map handler")

	// fmt.Println(mapTileHandler.mapTileNumbers)
	// fmt.Println(mapTileHandler.mapTiles)

	beetle := Beetle{character{x: 50, y: 50, direction: Up, speed: 4, currentSpriteNumber: 1, currentSprite: leftSprite1, spriteFrameSwitchThreshold: 12}}
	beetle.LoadImages()
	g := Game{beetle: &beetle, mapTileHandler: mapTileHandler}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
