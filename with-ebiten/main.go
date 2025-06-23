package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type obj struct {
	x, y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Game struct {
	xmin, xmax, ymin, ymax int
	obj                    obj
	d                      Direction
}

func NewGame(xmin, xmax, ymin, ymax int) *Game {
	return &Game{
		obj:  obj{},
		xmin: xmin,
		xmax: xmax,
		ymin: ymin,
		ymax: ymax,
	}
}

func (g *Game) Update() error {
	speed := 2
	switch g.d {
	case Up:
		if g.obj.y+speed == g.ymax {
			g.d = Right
		} else {
			g.obj.y += speed
		}
	case Right:
		if g.obj.x+speed == g.xmax {
			g.d = Down
		} else {
			g.obj.x += speed
		}
	case Down:
		if g.obj.y-speed == g.ymin {
			g.d = Left
		} else {
			g.obj.y -= speed
		}
	case Left:
		if g.obj.x-speed == g.xmin {
			g.d = Up
		} else {
			g.obj.x -= speed
		}
	default:
		return fmt.Errorf("unknown direction	 : %d", g.d)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := g.obj.x, g.obj.y
	screen.Set(x, y, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	game := NewGame(0, 100, 0, 100)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
