package platform

import (
	"arkanoid/internal/global"
	img "arkanoid/internal/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Platform struct {
	X      float64
	Y      float64
	DirX   float64
	Sprite *ebiten.Image
	HalfX  float64
	HalfY  float64
	Scale  float64
	SizeH  float64
	SizeW  float64
}

func NewPlatform() *Platform {
	sprite := img.GetSprite("../img/sprite.png", 51, 12, 29, 174)
	Scale := 2.5
	return &Platform{
		X:      (global.SCREEN_WIDTH / 2) - float64(sprite.Bounds().Dx()/2)*Scale,
		Y:      global.SCREEN_HEIGHT - 40,
		DirX:   5,
		Sprite: ebiten.NewImageFromImage(sprite),
		HalfX:  float64(sprite.Bounds().Dx()/2) * Scale,
		HalfY:  float64(sprite.Bounds().Dy()/2) * Scale,
		Scale:  Scale,
		SizeH:  12 * Scale,
		SizeW:  51 * Scale,
	}
}

func (p *Platform) Update() {

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.X+p.SizeW <= global.SCREEN_WIDTH {
		p.X += p.DirX
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.X-p.Scale >= 0 {
		p.X -= p.DirX
	}
}

func (p *Platform) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Sprite, op)
}
