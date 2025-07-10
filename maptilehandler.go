package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type MapTileHandler struct {
	mapTileNumbers [numWorldRows][numWorldColumns]int
	mapTiles       [11]MapTile
}

func (mapTileHandler MapTileHandler) LoadMap() {
	contents, err := os.ReadFile("assets/maps/map1.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	map_text := string(contents)
	lines := strings.Split(map_text, "\n")

	for r, line := range lines {
		numbers := strings.Split(line, " ")

		for c, number := range numbers {
			currentNumber := strings.Replace(number, "\n", "", 1)
			currentNumber = strings.Replace(number, "\r", "", 1)
			mapTileHandler.mapTileNumbers[r][c], err = strconv.Atoi(currentNumber)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (mapTileHandler MapTileHandler) LoadTileImages() {
	var err error
	grass, _, err = ebitenutil.NewImageFromFile("assets/tiles/Grass.png")
	if err != nil {
		log.Fatal(err)
	}
	grassTile := MapTile{Image: grass, IsSolid: false}
	mapTileHandler.mapTiles[0] = grassTile

	tree1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1.png")
	if err != nil {
		log.Fatal(err)
	}
	tree1Tile := MapTile{Image: tree1, IsSolid: true}
	mapTileHandler.mapTiles[1] = tree1Tile

	tree2, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree2.png")
	if err != nil {
		log.Fatal(err)
	}
	tree2Tile := MapTile{Image: tree2, IsSolid: true}
	mapTileHandler.mapTiles[2] = tree2Tile

	tree3, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree3.png")
	if err != nil {
		log.Fatal(err)
	}
	tree3Tile := MapTile{Image: tree3, IsSolid: true}
	mapTileHandler.mapTiles[3] = tree3Tile

	shrub, _, err = ebitenutil.NewImageFromFile("assets/tiles/Shrub.png")
	if err != nil {
		log.Fatal(err)
	}
	shrubTile := MapTile{Image: shrub, IsSolid: true}
	mapTileHandler.mapTiles[4] = shrubTile

	underbrush, _, err = ebitenutil.NewImageFromFile("assets/tiles/Underbrush.png")
	if err != nil {
		log.Fatal(err)
	}
	underbrushTile := MapTile{Image: underbrush, IsSolid: false}
	mapTileHandler.mapTiles[5] = underbrushTile

	shrubUnderbrush, _, err = ebitenutil.NewImageFromFile("assets/tiles/Shrub_Underbrush.png")
	if err != nil {
		log.Fatal(err)
	}
	shrubUnderbrushTile := MapTile{Image: shrubUnderbrush, IsSolid: false}
	mapTileHandler.mapTiles[6] = shrubUnderbrushTile

	orangeFlower, _, err = ebitenutil.NewImageFromFile("assets/tiles/OrangeFlower.png")
	if err != nil {
		log.Fatal(err)
	}
	orangeFlowerTile := MapTile{Image: orangeFlower, IsSolid: false}
	mapTileHandler.mapTiles[7] = orangeFlowerTile

	treeFlies1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1_Flies1.png")
	if err != nil {
		log.Fatal(err)
	}
	treeFliesTile := MapTile{Image: treeFlies1, IsSolid: true}
	mapTileHandler.mapTiles[8] = treeFliesTile

	treeBeehive1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1_Beehive1.png")
	if err != nil {
		log.Fatal(err)
	}
	treeBeehiveTile := MapTile{Image: treeBeehive1, IsSolid: true}
	mapTileHandler.mapTiles[9] = treeBeehiveTile

	treeWoodpecker1, _, err = ebitenutil.NewImageFromFile("assets/tiles/Tree1_Woodpecker1.png")
	if err != nil {
		log.Fatal(err)
	}
	treeWoodpeckerTile := MapTile{Image: treeWoodpecker1, IsSolid: true}
	mapTileHandler.mapTiles[10] = treeWoodpeckerTile
}

func (mapTileHandler MapTileHandler) DrawAll(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScale, imageScale)

	for row := range numWorldRows {
		for col := range numWorldColumns {
			op.GeoM.Translate(float64(row), float64(col))
			mapTileNumber := mapTileHandler.mapTileNumbers[row][col]
			image := mapTileHandler.mapTiles[mapTileNumber]
			screen.DrawImage(image.Image, op)
		}
	}
}
