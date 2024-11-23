package ball

import (
	"arkanoid/internal/global"
	img "arkanoid/internal/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	X            float64
	Y            float64
	DirY         int
	DirX         int
	HalfX        float64
	HalfY        float64
	Sprite       *ebiten.Image
	Scale        float64
	Size         float64
	FlagOut      bool
	FlagPlatform int8
}

func NewBall(dirX int, dirY int) *Ball {
	sprite := ebiten.NewImageFromImage(img.GetSprite("../img/sprite.png", 5, 5, 46, 94))
	Scale := 3.0
	return &Ball{
		X:            global.SCREEN_WIDTH / 2,
		Y:            global.SCREEN_HEIGHT / 2,
		DirX:         -3,
		DirY:         -3,
		HalfX:        float64(sprite.Bounds().Dx()/2) * Scale,
		HalfY:        float64(sprite.Bounds().Dy()/2) * Scale,
		Sprite:       sprite, //ebiten.NewImageFromImage(sprite.SubImage(r)),
		Scale:        Scale,
		Size:         5 * Scale,
		FlagOut:      false,
		FlagPlatform: 0,
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(b.Scale, b.Scale)
	op.GeoM.Translate(b.X-b.HalfX, b.Y-b.HalfY)
	screen.DrawImage(b.Sprite, op)
}

func (b *Ball) Update() {
	if b.Y < 400 {
		b.FlagPlatform = 0
	}
	if b.X-b.HalfX <= 0 {
		b.DirX = -b.DirX
	}
	if b.X+b.HalfX >= global.SCREEN_WIDTH {
		b.DirX = -b.DirX
	}
	if b.Y-b.HalfY <= 0 {
		b.DirY = -b.DirY
	}
	if b.Y+b.HalfY >= global.SCREEN_HEIGHT {
		b.FlagOut = true
	}
	b.X += float64(b.DirX)
	b.Y += float64(b.DirY)
}
