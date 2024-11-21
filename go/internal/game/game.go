package game

import (
	"arkanoid/internal/ball"
	"arkanoid/internal/brick"
	"arkanoid/internal/global"
	"arkanoid/internal/platform"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Ball       ball.Ball
	Brick      [][]*brick.Brick
	Platform   platform.Platform
	Background Background
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBkg(screen)
	g.Ball.Draw(screen)
	g.Platform.Draw(screen)
	g.CollisionPlatformBall()
	for i := range g.Brick {
		for j := range g.Brick[i] {
			g.Brick[i][j].Draw(screen)
		}
	}
}
func (g *Game) Update() error {
	g.Ball.Update()
	g.Platform.Update()
	g.CollisionBrickBall()
	return nil
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) CollisionPlatformBall() {
	if g.Ball.X <= g.Platform.X+g.Platform.HalfX*2 &&
		g.Ball.X+g.Ball.HalfX*2 >= g.Platform.X-g.Platform.HalfX &&
		g.Ball.Y+g.Ball.HalfY*2 == g.Platform.Y { //g.Ball.Y+g.Ball.HalfY*2 == g.Platform.Y
		g.Ball.DirY = -g.Ball.DirY
	}
}

func (g *Game) CollisionBrickBall() {
	for i := range g.Brick {
		for j := range g.Brick[i] {
			widthBrick := g.Brick[i][j].HalfX * 2
			heightBrick := g.Brick[i][j].HalfY * 2
			// top side of block
			touchTopBlock := g.Brick[i][j].Y == g.Ball.Y-g.Ball.HalfY*2

			// down side of block
			touchBotBlock := g.Brick[i][j].Y+heightBrick == g.Ball.Y

			//left side of block
			touchLeftBlock := g.Brick[i][j].X == g.Ball.X+g.Ball.HalfY*2

			//right side of block
			touchRightBlock := g.Brick[i][j].X+widthBrick == g.Ball.X

			betweenWidthBlock := g.Brick[i][j].X <= g.Ball.X && g.Brick[i][j].X+widthBrick >= g.Ball.X

			betweenHeightBlock := g.Brick[i][j].Y <= g.Ball.Y && g.Brick[i][j].Y+heightBrick >= g.Ball.Y

			if touchBotBlock && betweenWidthBlock {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirY = -g.Ball.DirY
			}

			if touchTopBlock && betweenWidthBlock {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirY = -g.Ball.DirY
			}

			if touchLeftBlock && betweenHeightBlock {
				g.Brick[i][j].X = -16 * 3
				g.Brick[i][j].Y = -16 * 3
				g.Brick[i][j].Appear = false
				g.Ball.DirX = -g.Ball.DirX
			}

			if touchRightBlock && betweenHeightBlock {
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
