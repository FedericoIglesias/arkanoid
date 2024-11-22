package paused

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

type Paused struct {
	Background *ebiten.Image
	Show       int8
	Scale      float64
	Timer      *timer.Timer
}

var mplusFaceSource *text.GoTextFaceSource
var rec = image.Rectangle{
	Min: image.Point{0, 0},
	Max: image.Point{X: global.SCREEN_WIDTH, Y: global.SCREEN_HEIGHT},
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func NewPaused() *Paused {
	return &Paused{
		Scale:      3,
		Show:       1,
		Background: ebiten.NewImageFromImage(rec),
		Timer:      timer.NewTimer(1000),
	}
}

func (p *Paused) Update() {
	p.Timer.Update()
	if (ebiten.IsKeyPressed(ebiten.KeyP) ||
		ebiten.IsKeyPressed(ebiten.KeyEscape)) && p.Timer.IsTimerDone() {
		p.Timer.Reset()
		if p.Show == 1 {
			p.Show = 0
		} else {
			p.Show = 1
		}
	}
}

func (p *Paused) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 125})
	screen.DrawImage(p.Background, op)

	opt := &text.DrawOptions{}
	opt.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	opt.GeoM.Scale(p.Scale, p.Scale)
	opt.GeoM.Translate(200, float64(100))
	text.Draw(screen, "PAUSED", &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   10 * p.Scale,
	}, opt)
}
