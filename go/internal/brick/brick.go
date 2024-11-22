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
	SizeH  float64
	SizeW  float64
}

func NewBrick(x, y float64, ID int) *Brick {
	sprite := img.GetSprite("../img/sprite.png", 16, 8, 30+(16*ID), 83)
	Scale := 3.0
	return &Brick{
		X:      x,
		Y:      y,
		Score:  ID,
		Appear: true,
		HalfX:  float64(sprite.Bounds().Dx()/2) * Scale,
		HalfY:  float64(sprite.Bounds().Dy()/2) * Scale,
		Sprite: ebiten.NewImageFromImage(sprite),
		Scale:  Scale,
		SizeH:  8 * Scale,
		SizeW:  16 * Scale,
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
			X := X + 16*(row)*3    // Be carefull because the size isn't dynamic
			Y := Y + 8*(column)*3  // Be carefull because the size isn't dynamic
			matrixBricks[column][row] = NewBrick(float64(X), float64(Y), ID)
		}
	}
	return matrixBricks
}
