package main

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
