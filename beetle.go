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
			beetle.currentSprite = upSprite1
			beetle.y -= beetle.speed
			if beetle.y < 0 {
				beetle.y = 0
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			beetle.direction = Down
			beetle.currentSprite = downSprite1
			beetle.y += beetle.speed
			if beetle.y+tileSize > gameHeight {
				beetle.y = gameHeight - tileSize
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			beetle.direction = Right
			beetle.currentSprite = rightSprite1
			beetle.x += beetle.speed
			if beetle.x+tileSize > gameWidth {
				beetle.x = gameWidth - tileSize
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			beetle.direction = Left
			beetle.currentSprite = leftSprite1
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

func (beetle *Beetle) Draw(screen *ebiten.Image) {
	switch beetle.direction {
	case Up:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = upSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = upSprite2
		}
	case Down:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = downSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = downSprite2
		}
	case Right:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = rightSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = rightSprite2
		}
	case Left:
		if beetle.currentSpriteNumber == 1 {
			beetle.currentSprite = leftSprite1
		} else if beetle.currentSpriteNumber == 2 {
			beetle.currentSprite = leftSprite2
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(imageScale, imageScale)
	op.GeoM.Translate(float64(beetle.x), float64(beetle.y))
	screen.DrawImage(beetle.currentSprite, op)
}
