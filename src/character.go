package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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

	upSprite1    *ebiten.Image
	downSprite1  *ebiten.Image
	rightSprite1 *ebiten.Image
	leftSprite1  *ebiten.Image
	upSprite2    *ebiten.Image
	downSprite2  *ebiten.Image
	rightSprite2 *ebiten.Image
	leftSprite2  *ebiten.Image
}
