package brick

import "github.com/hajimehoshi/ebiten/v2"

type Brick struct {
	X      float64
	Y      float64
	Score  int
	Sprite ebiten.Image
}
