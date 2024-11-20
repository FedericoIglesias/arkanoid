package game

import (
	"arkanoid/internal/ball"
	"arkanoid/internal/brick"
	"arkanoid/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Ball     ball.Ball
	Brick    [][]*brick.Brick
	Platform platform.Platform
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ball.Draw(screen)
	g.Platform.Draw(screen)

	for i := range g.Brick {
		for j := range g.Brick[i] {
			g.Brick[i][j].Draw(screen)
		}
	}
}
func (g *Game) Update() error {
	g.Ball.Update()
	g.Platform.Update()

	for i := range g.Brick {
		for j := range g.Brick[i] {

			// top side of block
			touchTopBlock := g.Brick[i][j].Y == g.Ball.Y-g.Ball.HalfY*6

			// down side of block
			touchBotBlock := g.Brick[i][j].Y+g.Brick[i][j].HalfY*6 == g.Ball.Y

			//left side of block
			touchLeftBlock := g.Brick[i][j].X == g.Ball.X+g.Ball.HalfY*6

			//right side of block
			touchRightBlock := g.Brick[i][j].X+(16*3) == g.Ball.X

			betweenWidthBlock := g.Brick[i][j].X <= g.Ball.X && g.Brick[i][j].X+(16*3) >= g.Ball.X

			betweenHeightBlock := g.Brick[i][j].Y <= g.Ball.Y && g.Brick[i][j].Y+g.Brick[i][j].HalfY*6 >= g.Ball.Y

			if touchBotBlock && betweenWidthBlock {
				g.Brick[i][j].X = -16*3
				g.Brick[i][j].Y = -16*3
				g.Brick[i][j].Appear = false
				g.Ball.DirY = -g.Ball.DirY
			}

			if touchTopBlock && betweenWidthBlock {
				g.Brick[i][j].X = -16 *3
				g.Brick[i][j].Y = -16*3
				g.Brick[i][j].Appear = false
				g.Ball.DirY = -g.Ball.DirY
			}

			if touchLeftBlock && betweenHeightBlock {
				g.Brick[i][j].X = -16*3
				g.Brick[i][j].Y = -16*3
				g.Brick[i][j].Appear = false
				g.Ball.DirX = -g.Ball.DirX
			}

			if touchRightBlock && betweenHeightBlock {
				g.Brick[i][j].X = -16*3
				g.Brick[i][j].Y = -16*3
				g.Brick[i][j].Appear = false
				g.Ball.DirX = -g.Ball.DirX
			}
		}
	}
	return nil
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
