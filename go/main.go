package main

import (
	"arkanoid/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := &game.Game{}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
