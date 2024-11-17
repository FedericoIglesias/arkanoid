package platform

import "github.com/hajimehoshi/ebiten/v2"

type Platform struct {
	X      float64
	Y      float64
	DirX   int
	Sprite ebiten.Image
}
