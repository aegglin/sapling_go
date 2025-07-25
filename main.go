package main

import (
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
	screen.Fill(color.White)
	g.mapTileHandler.DrawAll(screen)
	g.beetle.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameWidth, gameHeight
}

func main() {

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Sapling by Aiden Egglin")

	mapTileHandler := &MapTileHandler{}
	mapTileHandler.LoadMap()
	mapTileHandler.LoadTileImages()

	beetle := Beetle{character{x: 50, y: 50, direction: Up, speed: 4, currentSpriteNumber: 1, spriteFrameSwitchThreshold: 12}}
	beetle.LoadImages()
	beetle.currentSprite = beetle.downSprite1
	g := Game{beetle: &beetle, mapTileHandler: mapTileHandler}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
