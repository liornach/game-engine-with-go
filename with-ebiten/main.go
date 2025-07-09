package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/achtung"
)

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	greenPl := achtung.NewPlayer(achtung.Green, ebiten.KeyArrowLeft, ebiten.KeyArrowRight)
	redPlayer := achtung.NewPlayer(achtung.Red, ebiten.KeyA, ebiten.KeyD)
	players := []*achtung.Player{greenPl, redPlayer}
	rotation := 0.08
	vel := achtung.Velocity{
		X: 40,
		Y: 40,
	}

	bg := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}

	game, err := achtung.NewGame(players, rotation, 1, 1, vel, bg)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer game.Close()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
