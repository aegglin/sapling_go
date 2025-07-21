package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Beetle struct {
	character
}

func (beetle *Beetle) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) ||
		ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) ||
		ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {

		if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			beetle.direction = Up
			beetle.currentSprite = beetle.upSprite1
			beetle.y -= beetle.speed
			if beetle.y < 0 {
				beetle.y = 0
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			beetle.direction = Down
			beetle.currentSprite = beetle.downSprite1
			beetle.y += beetle.speed
			if beetle.y+tileSize > gameHeight {
				beetle.y = gameHeight - tileSize
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			beetle.direction = Right
			beetle.currentSprite = beetle.rightSprite1
			beetle.x += beetle.speed
			if beetle.x+tileSize > gameWidth {
				beetle.x = gameWidth - tileSize
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			beetle.direction = Left
			beetle.currentSprite = beetle.leftSprite1
			beetle.x -= beetle.speed
			if beetle.x < 0 {
				beetle.x = 0
			}
		}
		beetle.spriteUpdateFrameCount++
		if beetle.spriteUpdateFrameCount > beetle.spriteFrameSwitchThreshold {
			if beetle.currentSpriteNumber == 1 {
				beetle.currentSpriteNumber = 2
			} else if beetle.currentSpriteNumber == 2 {
				beetle.currentSpriteNumber = 1
			}
			beetle.spriteUpdateFrameCount = 0
		}
	}
}

func (beetle *Beetle) LoadImages() {
	var err error

	beetle.upSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleUp1.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.downSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleDown1.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.rightSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleRight1.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.leftSprite1, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleLeft1.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.upSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleUp2.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.downSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleDown2.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.rightSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleRight2.png")
	if err != nil {
		log.Fatal(err)
	}
	beetle.leftSprite2, _, err = ebitenutil.NewImageFromFile("assets/beetle/BeetleLeft2.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (beetle *Beetle) Draw(screen *ebiten.Image) {
	switch beetle.direction {
	case Up:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = beetle.upSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = beetle.upSprite2
		}
	case Down:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = beetle.downSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = beetle.downSprite2
		}
	case Right:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = beetle.rightSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = beetle.rightSprite2
		}
	case Left:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = beetle.leftSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = beetle.leftSprite2
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScale, imageScale)
	op.GeoM.Translate(float64(beetle.x), float64(beetle.y))
	screen.DrawImage(beetle.currentSprite, op)
}
