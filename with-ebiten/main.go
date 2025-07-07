package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/achtung"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	bluePlayer := achtung.NewPlayer(achtung.Blue, ebiten.KeyArrowLeft, ebiten.KeyArrowRight)
	//redPlayer := achtung.NewPlayer(achtung.Red, ebiten.KeyA, ebiten.KeyD)
	//players := []*achtung.Player{bluePlayer, redPlayer}
	players := []*achtung.Player{bluePlayer}
	rotation := 0.1

	game, err := achtung.NewGame(players, rotation, 1, 1)
	defer game.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
