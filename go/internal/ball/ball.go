package ball

import (
	"arkanoid/internal/global"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var sprite = global.LoadImg("img/sprite.png")

type Ball struct {
	X      float64
	Y      float64
	DirY   int
	DirX   int
	Width  int
	Sprite *ebiten.Image
}

func NewBall(dirX int, dirY int) *Ball {
	r := image.Rectangle{
		Min: image.Point{46, 94},
		Max: image.Point{51, 99},
	}
	return &Ball{
		X:      global.SCREEN_WIDTH / 2,
		Y:      global.SCREEN_HEIGHT / 2,
		DirX:   0,
		DirY:   0,
		Width:  2,
		Sprite: ebiten.NewImageFromImage(sprite.SubImage(r)),
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.X, b.Y)
	op.GeoM.Scale(3, 3)
	screen.DrawImage(b.Sprite, op)
}

func (b *Ball) Update() {
	b.X += float64(b.DirX)
	b.Y += float64(b.DirY)
}
