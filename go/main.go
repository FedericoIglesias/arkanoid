package main

import (
	"arkanoid/internal/game"
	"arkanoid/internal/global"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(global.SCREEN_WIDTH, global.SCREEN_HEIGHT)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
