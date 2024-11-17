package game

import (
	"arkanoid/internal/ball"
	"arkanoid/internal/brick"
	"arkanoid/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Ball     ball.Ball
	Brick    brick.Brick
	Platform platform.Platform
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ball.Draw(screen)
	
}
func (g *Game) Update() error {
	g.Ball.Update()
	return nil
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
