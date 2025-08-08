package achtung

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	"slices"
)


type coord struct {
	x, y int
}

type Game struct {
	players []*player
	board map[coord]Color
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

func (g *Game) AddPlayer(c Color, left, right ebiten.Key) bool {
	if slices.ContainsFunc(g.players ,func(p *Player) bool { return p.color == c}) {
		return false
	}
	
	g.players = append(g.players, player{...})
}
	



type Color = int
const (
	Red Color = 0XFF0000
	Green Color = 0X00FF00
	Blue Color = 0X0000FF
)

type player struct {
	color Color
	left, right ebiten.Key
	x, y float64
	vx, vy float64
}
