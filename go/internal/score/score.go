package score

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Score struct {
	Score int
	Scale float64
}

func NewScore() *Score {
	return &Score{
		Score: 0,
		Scale: 2,
	}
}

var mplusFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func (s *Score) Draw(screen *ebiten.Image) {
	textEmbed := fmt.Sprintf("SCORE: %v", s.Score)

	opt := &text.DrawOptions{}
	opt.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 150})
	opt.GeoM.Scale(s.Scale, s.Scale)
	opt.GeoM.Translate(0, 0)
	text.Draw(screen, textEmbed, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   5 * s.Scale,
	}, opt)
}
func (s *Score) Update(score int) {
	s.Score += score
}

// func (s *Score) Draw(){

// }
