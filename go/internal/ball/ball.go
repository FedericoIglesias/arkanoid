package ball

import (
	"arkanoid/internal/global"
	img "arkanoid/internal/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	X      float64
	Y      float64
	DirY   int
	DirX   int
	Width  int
	Sprite *ebiten.Image
}

func NewBall(dirX int, dirY int) *Ball {
	return &Ball{
		X:      global.SCREEN_WIDTH / 2,
		Y:      global.SCREEN_HEIGHT / 2,
		DirX:   1,
		DirY:   1,
		Width:  2,
		Sprite: ebiten.NewImageFromImage(img.GetSprite("internal/sprite/img/sprite.png", 5, 5, 46, 94)), //ebiten.NewImageFromImage(sprite.SubImage(r)),
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	halfX := b.Sprite.Bounds().Dx() / 2
	halfY := b.Sprite.Bounds().Dy() / 2
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(b.X-float64(halfX), b.Y-float64(halfY))
	screen.DrawImage(b.Sprite, op)
}

func (b *Ball) Update() {
	// b.X += float64(b.DirX)
	// b.Y += float64(b.DirY)
}
