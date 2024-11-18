package main

import (
	"arkanoid/internal/ball"
	"arkanoid/internal/game"
	"arkanoid/internal/global"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := &game.Game{
		Ball: *ball.NewBall(2, 2),
	}
	
	ebiten.SetWindowSize(global.SCREEN_WIDTH, global.SCREEN_HEIGHT)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
