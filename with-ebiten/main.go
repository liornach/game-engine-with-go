package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/game"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
