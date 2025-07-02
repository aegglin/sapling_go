package main

import "github.com/hajimehoshi/ebiten/v2"

type MapTile struct {
	Image   *ebiten.Image
	IsSolid bool
}
