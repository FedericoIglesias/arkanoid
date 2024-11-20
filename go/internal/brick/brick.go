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
	Sprite ebiten.Image
	halfX  float64
	halfY  float64
	draw   bool
}

func NewBrick(x, y float64, ID int) *Brick {
	sprite := img.GetSprite("sprite.png", 16, 8, 30+(16*ID), 83)

	return &Brick{
		X:      x,
		Y:      y,
		Score:  ID,
		Sprite: *ebiten.NewImageFromImage(sprite),
		halfX:  float64(sprite.Bounds().Dx() / 2),
		halfY:  float64(sprite.Bounds().Dy() / 2),
		draw:   true,
	}
}

func (b *Brick) Update() {

}

func (b *Brick) Draw(screen *ebiten.Image) {
	if b.draw {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(3, 3)
		op.GeoM.Translate(b.X, b.Y)
		screen.DrawImage(&b.Sprite, op)
	}
}

func CreateMatrixBricks(X, Y int) [][]*Brick {

	matrixBRicks := make([][]*Brick, 7)
	for column := range 7 {
		matrixBRicks[column] = make([]*Brick, 15)
		for row := range 15 {
			ID := (rand.Intn(7) + 1)
			X := X + 16*(row)*3
			Y := Y + 8*(column)*3
			matrixBRicks[column][row] = NewBrick(float64(X), float64(Y), ID)
		}
	}
	return matrixBRicks
}
