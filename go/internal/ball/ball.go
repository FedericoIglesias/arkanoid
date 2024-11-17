package ball

import "github.com/hajimehoshi/ebiten/v2"

type Ball struct {
	X      float64
	Y      float64
	DirY   int
	DirX   int
	Width  int
	Sprite ebiten.Image
}
