package main

import "github.com/hajimehoshi/ebiten/v2"

type MapTile struct {
	isSolid bool
	name    string
	image   *ebiten.Image
}
