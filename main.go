package main

import (
	"image/color"
	_ "image/png"
	"log"
	"strconv"

	"fmt"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	leftSprite1  *ebiten.Image
	rightSprite1 *ebiten.Image
	upSprite1    *ebiten.Image
	downSprite1  *ebiten.Image
	leftSprite2  *ebiten.Image
	rightSprite2 *ebiten.Image
	upSprite2    *ebiten.Image
	downSprite2  *ebiten.Image
)

func init() {
	loadTileImages()
	loadBeetleImages()
}

func loadBeetleImages() {
	var err error

	upSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleUp1.png")
	if err != nil {
		log.Fatal(err)
	}
	downSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleDown1.png")
	if err != nil {
		log.Fatal(err)
	}
	rightSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleRight1.png")
	if err != nil {
		log.Fatal(err)
	}
	leftSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleLeft1.png")
	if err != nil {
		log.Fatal(err)
	}
	upSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleUp2.png")
	if err != nil {
		log.Fatal(err)
	}
	downSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleDown2.png")
	if err != nil {
		log.Fatal(err)
	}
	rightSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleRight2.png")
	if err != nil {
		log.Fatal(err)
	}
	leftSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleLeft2.png")
	if err != nil {
		log.Fatal(err)
	}
}

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

func loadTileImages() {
	var err error
	grass, _, err = ebitenutil.NewImageFromFile("assets/tiles/Grass.png")
	if err != nil {
		log.Fatal(err)
	}
	tree1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1.png")
	if err != nil {
		log.Fatal(err)
	}
	tree2, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree2.png")
	if err != nil {
		log.Fatal(err)
	}
	tree3, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree3.png")
	if err != nil {
		log.Fatal(err)
	}
	shrub, _, err = ebitenutil.NewImageFromFile("assets/tiles/Shrub.png")
	if err != nil {
		log.Fatal(err)
	}
	underbrush, _, err = ebitenutil.NewImageFromFile("assets/tiles/Underbrush.png")
	if err != nil {
		log.Fatal(err)
	}
	orangeFlower, _, err = ebitenutil.NewImageFromFile("assets/tiles/OrangeFlower.png")
	if err != nil {
		log.Fatal(err)
	}
	treeFlies1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1_Flies1.png")
	if err != nil {
		log.Fatal(err)
	}
	treeBeehive1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1_Beehive1.png")
	if err != nil {
		log.Fatal(err)
	}
	treeWoodpecker1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1_Woodpecker1.png")
	if err != nil {
		log.Fatal(err)
	}
}

func loadMap(g *Game) {
	contents, err := os.ReadFile("assets/maps/map1.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	map_text := string(contents)
	lines := strings.Split(map_text, "\r")

	for r, line := range lines {
		numbers := strings.Split(line, " ")
		for c, number := range numbers {
			fmt.Printf("At %d, %d, the number is: %s\n", r, c, number)
			g.mapTileNumbers[r][c], err = strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
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
	x                          int
	y                          int
	speed                      int
	direction                  Direction
	currentSprite              *ebiten.Image
	currentSpriteNumber        int
	spriteUpdateFrameCount     int
	spriteFrameSwitchThreshold int
}

// the game has the main character beetle and the tiles
type Game struct {
	beetle         *character
	mapTileNumbers [numWorldColumns][numWorldRows]int
	// mapTiles       []MapTile
}

func (g *Game) Update() error {

	if g.beetle == nil {
		g.beetle = &character{x: 50, y: 50, direction: Down, speed: 4, currentSprite: downSprite1, currentSpriteNumber: 1, spriteUpdateFrameCount: 0, spriteFrameSwitchThreshold: 12}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.beetle.direction = Up
		g.beetle.currentSprite = upSprite1
		g.beetle.y -= g.beetle.speed
		if g.beetle.y < 0 {
			g.beetle.y = 0
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.beetle.direction = Down
		g.beetle.currentSprite = downSprite1
		g.beetle.y += g.beetle.speed
		if g.beetle.y+tileSize > gameHeight {
			g.beetle.y = gameHeight - tileSize
		}

	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.beetle.direction = Right
		g.beetle.currentSprite = rightSprite1
		g.beetle.x += g.beetle.speed
		if g.beetle.x+tileSize > gameWidth {
			g.beetle.x = gameWidth - tileSize
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.beetle.direction = Left
		g.beetle.currentSprite = leftSprite1
		g.beetle.x -= g.beetle.speed
		if g.beetle.x < 0 {
			g.beetle.x = 0
		}
	}

	g.beetle.spriteUpdateFrameCount++
	if g.beetle.spriteUpdateFrameCount > g.beetle.spriteFrameSwitchThreshold {
		if g.beetle.currentSpriteNumber == 1 {
			g.beetle.currentSpriteNumber = 2
		} else if g.beetle.currentSpriteNumber == 2 {
			g.beetle.currentSpriteNumber = 1
		}
		g.beetle.spriteUpdateFrameCount = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	switch g.beetle.direction {
	case Up:
		if g.beetle.currentSpriteNumber == 1 {
			g.beetle.currentSprite = upSprite1
		} else if g.beetle.currentSpriteNumber == 2 {
			g.beetle.currentSprite = upSprite2
		}
	case Down:
		if g.beetle.currentSpriteNumber == 1 {
			g.beetle.currentSprite = downSprite1
		} else if g.beetle.currentSpriteNumber == 2 {
			g.beetle.currentSprite = downSprite2
		}
	case Right:
		if g.beetle.currentSpriteNumber == 1 {
			g.beetle.currentSprite = rightSprite1
		} else if g.beetle.currentSpriteNumber == 2 {
			g.beetle.currentSprite = rightSprite2
		}
	case Left:
		if g.beetle.currentSpriteNumber == 1 {
			g.beetle.currentSprite = leftSprite1
		} else if g.beetle.currentSpriteNumber == 2 {
			g.beetle.currentSprite = leftSprite2
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScale, imageScale)
	op.GeoM.Translate(float64(g.beetle.x), float64(g.beetle.y))
	screen.Fill(color.White)
	screen.DrawImage(g.beetle.currentSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameWidth, gameHeight
}

func main() {

	ebiten.SetWindowSize(gameWidth, gameHeight)
	ebiten.SetWindowTitle("Sapling by Aiden Egglin")

	beetle := character{x: 50, y: 50, direction: Up, speed: 4, currentSprite: leftSprite1}
	g := Game{beetle: &beetle}
	loadMap(&g)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
