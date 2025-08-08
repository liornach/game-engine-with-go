package achtung

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "hello world")
}

func (g *Game) Layout(outWidth, outHeight int) (screenWidth, screenHeight int) {
	return outWidth, outHeight 
}
