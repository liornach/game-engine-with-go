package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/logic"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	g := logic.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
