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
	halfX  float64
	halfY  float64
}

func NewBall(dirX int, dirY int) *Ball {
	sprite := ebiten.NewImageFromImage(img.GetSprite("internal/sprite/img/sprite.png", 5, 5, 46, 94))
	return &Ball{
		X:      global.SCREEN_WIDTH / 2,
		Y:      global.SCREEN_HEIGHT / 2,
		DirX:   -10,
		DirY:   -10,
		Width:  2,
		Sprite: sprite, //ebiten.NewImageFromImage(sprite.SubImage(r)),
		halfX:  float64(sprite.Bounds().Dx() / 2),
		halfY:  float64(sprite.Bounds().Dy() / 2),
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(b.X-b.halfX, b.Y-b.halfY)
	screen.DrawImage(b.Sprite, op)
}

func (b *Ball) Update() {
	if b.X-b.halfX <= 0 {
		b.DirX = -b.DirX
	}
	if b.X+b.halfX >= global.SCREEN_WIDTH {
		b.DirX = -b.DirX
	}
	if b.Y-b.halfY <= 0 {
		b.DirY = -b.DirY
	}
	if b.Y+b.halfY >= global.SCREEN_HEIGHT {
		b.DirY = -b.DirY
	}

	b.X += float64(b.DirX)
	b.Y += float64(b.DirY)
}
