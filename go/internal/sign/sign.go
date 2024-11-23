package sign

import (
	"arkanoid/internal/global"
	"arkanoid/internal/timer"
	"bytes"

	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Sign struct {
	Background *ebiten.Image
	Show       bool
	Lossed     bool
	Paused     bool
	Winner     bool
	Scale      float64
	Timer      *timer.Timer
}

var mplusFaceSource *text.GoTextFaceSource

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

var rec = image.Rectangle{
	Min: image.Point{0, 0},
	Max: image.Point{X: global.SCREEN_WIDTH, Y: global.SCREEN_HEIGHT},
}

func NewPaused() *Sign {
	return &Sign{
		Scale:      3,
		Show:       true,
		Paused:     true,
		Winner:     false,
		Lossed:     false,
		Background: ebiten.NewImageFromImage(rec),
		Timer:      timer.NewTimer(1000),
	}
}

func (s *Sign) Update() {
	s.Timer.Update()
	if (ebiten.IsKeyPressed(ebiten.KeyP) ||
		ebiten.IsKeyPressed(ebiten.KeyEscape)) && s.Timer.IsTimerDone() {
		s.Timer.Reset()
		if s.Show {
			s.Show = false
		} else {
			s.Show = true
		}
	}

}

func (s *Sign) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 125})
	screen.DrawImage(s.Background, op)

	opt := &text.DrawOptions{}
	opt.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	opt.GeoM.Scale(s.Scale, s.Scale)
	opt.GeoM.Translate(200, float64(100))
	text.Draw(screen, s.SelectWord(), &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   10 * s.Scale,
	}, opt)
}

func (s *Sign) SelectWord() string {

	if s.Lossed {
		return "You Lost"
	}
	if s.Paused {
		return "Paused"
	}
	if s.Winner {
		return "You Win"
	}
	return ""
}
