package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/achtung"
	"github.com/liornach/game-engine-ebiten/achtung/states"
)

type worldPos = achtung.WorldPos

func main() {
	ebiten.SetWindowSize(800, 800)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetTPS(ebiten.SyncWithFPS)

	rotation := 0.08
	vel := achtung.Velocity{
		X: 40,
		Y: 40,
	}
	greenPl := achtung.NewPlayer(achtung.Green, ebiten.KeyArrowLeft, ebiten.KeyArrowRight, rotation, vel)
	redPlayer := achtung.NewPlayer(achtung.Red, ebiten.KeyA, ebiten.KeyD, rotation, vel)

	players := []*achtung.Player{greenPl, redPlayer}

	bg := color.RGBA{
		R: 0,
		G: 0,
		B: 0,
		A: 0,
	}

	border := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	posOptions := []worldPos{worldPos{
		X: 65,
		Y: 30,
	}, worldPos{
		X: 32,
		Y: 78,
	}}

	randomPos := states.NewRandomPos(posOptions)
	initialState := states.NewInitialState(randomPos)
	game, err := achtung.NewGame(1, 1, bg, border, initialState)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, p := range players {
		if err := game.RegisterPlayer(*p); err != nil {
			panic(err)
		}
	}

	game.RegisterPlayer(*greenPl)
	game.RegisterPlayer(*redPlayer)

	defer game.Close()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
