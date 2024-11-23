package main

import (
	"arkanoid/internal/ball"
	"arkanoid/internal/brick"
	"arkanoid/internal/game"
	"arkanoid/internal/global"
	"arkanoid/internal/paused"
	"arkanoid/internal/platform"
	"arkanoid/internal/score"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &game.Game{
		Ball:       *ball.NewBall(2, 2),
		Platform:   *platform.NewPlatform(),
		Brick:      brick.CreateMatrixBricks(40, 80),
		Background: *game.NewBackground(),
		Paused:     *paused.NewPaused(),
		Score:      *score.NewScore(),
	}

	ebiten.SetWindowSize(global.SCREEN_WIDTH, global.SCREEN_HEIGHT)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
