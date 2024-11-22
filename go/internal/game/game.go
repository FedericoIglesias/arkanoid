package game

import (
	"arkanoid/internal/ball"
	"arkanoid/internal/brick"
	"arkanoid/internal/global"
	"arkanoid/internal/paused"
	"arkanoid/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Ball       ball.Ball
	Brick      [][]*brick.Brick
	Platform   platform.Platform
	Background Background
	Paused     paused.Paused
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBkg(screen)
	g.Platform.Draw(screen)
	g.CollisionPlatformBall()
	for i := range g.Brick {
		for j := range g.Brick[i] {
			g.Brick[i][j].Draw(screen)
		}
	}
	g.Ball.Draw(screen)
	if g.Paused.Show == 1 {
		g.Paused.Draw(screen)
	}
}
func (g *Game) Update() error {
	g.Paused.Update()
	if g.Paused.Show == 1 {
		return nil
	}
	g.Platform.Update()
	g.Ball.Update()
	g.CollisionBrickBall()
	// if g.Ball.FlagOut == 1 {
	// panic("endGam")
	// }
	return nil
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) CollisionPlatformBall() {
	betweenHeightPlatform := g.Ball.Y+g.Ball.Size >= g.Platform.Y && g.Ball.Y+g.Ball.Size <= g.Platform.Y+g.Platform.HalfY
	betweenWidthPlatform := g.Ball.X <= g.Platform.X+g.Platform.SizeW && g.Ball.X+g.Ball.Size >= g.Platform.X-g.Platform.Scale
	if betweenWidthPlatform && betweenHeightPlatform && g.Ball.FlagPlatform == 0 {
		g.Ball.FlagPlatform = 1
		g.Ball.DirY = -g.Ball.DirY
	}

}

func (g *Game) CollisionBrickBall() {
	for i := range g.Brick {
		for j := range g.Brick[i] {
			widthBrick := g.Brick[i][j].SizeW
			heightBrick := g.Brick[i][j].SizeH
			fringe := 1 * g.Brick[i][j].Scale
			betweenWidthBlockRight := g.Brick[i][j].X+widthBrick-fringe >= g.Ball.X && g.Brick[i][j].X+widthBrick <= g.Ball.X
			betweenWidthBlockLeft := g.Brick[i][j].X >= g.Ball.X && g.Brick[i][j].X+fringe <= g.Ball.X
			betweenHeightBlockTop := g.Brick[i][j].Y <= g.Ball.Y && g.Brick[i][j].Y+fringe >= g.Ball.Y
			betweenHeightBlockBot := g.Brick[i][j].Y+heightBrick >= g.Ball.Y && g.Brick[i][j].Y+heightBrick-fringe <= g.Ball.Y
			betweenHeight := g.Brick[i][j].Y >= g.Ball.Y && g.Brick[i][j].Y+heightBrick <= g.Ball.Y
			betweenWidth := g.Brick[i][j].X <= g.Ball.X && g.Brick[i][j].X+widthBrick >= g.Ball.X

			if betweenHeightBlockTop && betweenWidth {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirY = -g.Ball.DirY
			}

			if betweenHeightBlockBot && betweenWidth {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirY = -g.Ball.DirY
			}
			if betweenWidthBlockLeft && betweenHeight {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirX = -g.Ball.DirX
			}
			if betweenWidthBlockRight && betweenHeight {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirX = -g.Ball.DirX
			}

		}
	}
}

func (g *Game) DrawBkg(screen *ebiten.Image) {
	timesX := global.SCREEN_WIDTH / g.Background.SizeW
	timesY := global.SCREEN_HEIGHT / g.Background.SizeHeight
	for vX := range int(timesX + 1) {
		for vY := range int(timesY + 1) {
			g.Background.PositionX = float64(vX) * g.Background.SizeW
			g.Background.PositionY = float64(vY) * g.Background.SizeHeight
			g.Background.Draw(screen)
		}
	}

}
