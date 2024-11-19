package brick

import (
	img "arkanoid/internal/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Brick struct {
	X      float64
	Y      float64
	Score  int
	Sprite ebiten.Image
	halfX  float64
	halfY  float64
}

func NewBrick(x, y float64, ID int) *Brick {
	sprite := img.GetSprite("sprite.png", 128, 8, 30, 83)

	return &Brick{
		X:      x,
		Y:      y,
		Score:  ID,
		Sprite: *ebiten.NewImageFromImage(sprite),
		halfX:  float64(sprite.Bounds().Dx() / 2),
		halfY:  float64(sprite.Bounds().Dy() / 2),
	}
}

func (b *Brick) Update() {

}

func (b *Brick) Draw(screen *ebiten.Image) {

}
