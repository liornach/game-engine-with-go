package ebiten

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/logic"
)

type Clbk func(ebiten.Key)
type DrawClbk func(ebiten.Image)

type Game struct {
	clbks                     map[ebiten.Key]Clbk
	world                     *logic.World
	colors                    map[logic.Owner]color.Color
	defaultColor              color.Color
	xratio, yratio            float64
	screenWidth, screenHeight int
}

func NewGame(w *logic.World, x, y float64, dc color.Color, screenWidth, screenHeight int) *Game {
	return &Game{
		clbks:        map[ebiten.Key]Clbk{},
		world:        w,
		xratio:       x,
		yratio:       y,
		colors:       map[logic.Owner]color.Color{},
		defaultColor: dc,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
	}
}

func (g *Game) SetColor(o logic.Owner, c color.Color) bool {
	_, ok := g.colors[o]
	if !ok {
		return false
	}

	g.colors[o] = c
	return true
}

func (g *Game) Update() error {
	for k, f := range g.clbks {
		if isPressed(k) {
			f(k)
		}
	}

	return nil
}

func isPressed(k ebiten.Key) bool {
	return ebiten.IsKeyPressed(k)
}

func (g *Game) Draw(screen *ebiten.Image) {
	m := g.world.Map()
	img := image.NewRGBA(image.Rect(0, 0, g.screenWidth, g.screenHeight))
	m.Range(func(p logic.Pos, o logic.Owner) {
		x := float64(p.X) * g.xratio
		y := float64(p.Y) * g.yratio
		color, ok := g.colors[o]
		if !ok {
			panic(fmt.Errorf("could not find owner :%v", o))
		}

		img.Set(x, y, color) // how to floats
	})
}

// TODO : still need to understand that
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) SetOnKeyPress(k ebiten.Key, c Clbk) error {
	var err error
	if g.IsRegistered(k) {
		err = AlreadyRegisteredError{}
	} else {
		g.clbks[k] = c
	}

	return err
}

func (g *Game) IsRegistered(k ebiten.Key) bool {
	_, ok := g.clbks[k]
	return ok
}

func (g *Game) RmKeyPress(k ebiten.Key) {
	delete(g.clbks, k)
}

type AlreadyRegisteredError struct{}

func (e AlreadyRegisteredError) Error() string {
	return "already registerd"
}

func (g *Game) SetOnDraw(c DrawClbk) error {
	var e error
	if g.onDraw != nil {
		e = AlreadyRegisteredError{}
	} else {
		g.onDraw = c
	}

	return e
}
