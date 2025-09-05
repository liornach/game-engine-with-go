package achtung

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2"
	"slices"
	"sync"
)


type coord struct {
	x, y int
}

type game struct {
	// todo : find a proper ds to store both []*Player (players) and board, in term of iterating the board to find collision
	players []*Player
	board map[coord]Color
	maxCoord, minCoord coord
	initOnce sync.Once
	isRunning bool
	initMutex sync.Mutex
}

func NewGame() *game {
	return &game{
		players: []*Player,
		board : make(map[coord]Color)
	}
}

func (g *game) Update() error {
	g.initOnce.Do(func(){
		if isRunning {
			return errors.New("game has already been initialized, why do it twice?")
		}
		initMutex.Lock()
		defer initMutex.Unlock()
		g.initPlayersLocationAndVelocity()
		isRunning = true
	})

	for i := range g.players {
			
	}
	return nil
}


func (g *game) updatePlayersLocation(ticks int) []collision {
	var colls []collision
	for i := range g.players {

	}

}

func (g *game) initPlayersLocationAndVelocity() {
	if len(g.players) != 2 {
		panic("currently supports only 2 players")
	}

	c := g.minCoord
	for i := range g.players {
		c = coord{
			x : c.x + 33,
			y : c.y + 66,
		}
		g.players[i].x = float64(c.x)
		g.players[i].y = float64(c.y)
		g.players[i].vx = 20
		g.players[i].vy = 20
	}
}
		

func (g *game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "hello world")
}

func (g *game) Layout(outWidth, outHeight int) (screenWidth, screenHeight int) {
	return outWidth, outHeight 
}

func (g *game) AddPlayer(c Color, Left, Right ebiten.Key) bool {
	if slices.ContainsFunc(g.players ,func(p *Player) bool { return p.Color == c}) {
		return false
	}
	
	g.players = append(g.players, &Player{Color : c, Left : Left, Right : Right})
	return true
}
