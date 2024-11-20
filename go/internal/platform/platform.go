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
	halfX  float64
	halfY  float64
}

func NewPlatform() *Platform {
	sprite := img.GetSprite("sprite.png", 51, 12, 29, 174)

	return &Platform{
		X:      global.SCREEN_WIDTH / 2,
		Y:      global.SCREEN_HEIGHT - 40,
		DirX:   10,
		Sprite: ebiten.NewImageFromImage(sprite),
		halfX:  float64(sprite.Bounds().Dx() / 2),
		halfY:  float64(sprite.Bounds().Dy() / 2),
	}
}

func (p *Platform) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && p.X+p.halfX*3 <= global.SCREEN_WIDTH {
		p.X += p.DirX
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && p.X-p.halfX >= 0 {
		p.X -= p.DirX
	}
}

func (p *Platform) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(2.5, 2.5)
	op.GeoM.Translate(p.X-p.halfX, p.Y-p.halfY)
	screen.DrawImage(p.Sprite, op)
}
