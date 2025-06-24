package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Clbk func(ebiten.Key)
type DrawClbk func(ebiten.Image)

type Game struct {
	clbks  map[ebiten.Key]Clbk
	onDraw DrawClbk
}

func NewGame() *Game {
	return &Game{}
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

}

// TODO : still need to understand that
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) OnKeyPress(k ebiten.Key, c Clbk) error {
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

func (g *Game) OnDraw(c DrawClbk) error {
	var e error
	if g.onDraw != nil {
		e = AlreadyRegisteredError{}
	} else {
		g.onDraw = c
	}

	return e
}
