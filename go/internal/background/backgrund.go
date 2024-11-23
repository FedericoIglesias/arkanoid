package background

import (
	img "arkanoid/internal/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct {
	Sprite     *ebiten.Image
	Scale      float64
	SizeW      float64
	SizeHeight float64
	PositionX  float64
	PositionY  float64
}

func NewBackground() *Background {
	sprite := ebiten.NewImageFromImage(img.GetSprite("../img/bkg.png", 24, 32, 0, 0))
	Scale := 3.0
	return &Background{
		Sprite:     sprite,
		Scale:      Scale,
		SizeW:      24 * Scale,
		SizeHeight: 32 * Scale,
		PositionX:  0.0,
		PositionY:  0.0,
	}
}

func (b *Background) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.Scale, b.Scale)
	op.GeoM.Translate(b.PositionX, b.PositionY)
	screen.DrawImage(b.Sprite, op)
}
