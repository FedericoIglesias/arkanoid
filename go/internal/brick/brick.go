package brick

import (
	img "arkanoid/internal/sprite"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Brick struct {
	X      float64
	Y      float64
	Score  int
	HalfX  float64
	HalfY  float64
	Appear bool
	Sprite *ebiten.Image
	Scale  float64
}

func NewBrick(x, y float64, ID int) *Brick {
	sprite := img.GetSprite("sprite.png", 16, 8, 30+(16*ID), 83)

	return &Brick{
		X:      x,
		Y:      y,
		Score:  ID,
		Appear: true,
		HalfX:  float64(sprite.Bounds().Dx() / 2),
		HalfY:  float64(sprite.Bounds().Dy() / 2),
		Sprite: ebiten.NewImageFromImage(sprite),
		Scale:  3,
	}
}

func (b *Brick) Update() {

}

func (b *Brick) Draw(screen *ebiten.Image) {
	if b.Appear {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(b.Scale, b.Scale)
		op.GeoM.Translate(b.X, b.Y)
		screen.DrawImage(b.Sprite, op)
	}
}

func CreateMatrixBricks(X, Y int) [][]*Brick {

	matrixBricks := make([][]*Brick, 7)
	for column := range 7 {
		matrixBricks[column] = make([]*Brick, 15)
		for row := range 15 {
			ID := (rand.Intn(7) + 1)
			X := X + 16*(row)*3
			Y := Y + 8*(column)*3
			matrixBricks[column][row] = NewBrick(float64(X), float64(Y), ID)
		}
	}
	return matrixBricks
}
