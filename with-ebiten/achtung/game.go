package achtung

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type uid = string

type objectInWorld interface {
	isCollided(other objectInWorld, pos worldPos) bool
	uid() uid
	color() color.RGBA
}

type worldPos struct {
	x, y int
}

func newWorldPos(x, y int) worldPos {
	return worldPos{
		x: x,
		y: y,
	}
}

type collision struct {
	objectsInvolved []objectInWorld
	pos             worldPos
}

type state interface {
	update(g *game) (bool, state)
}

type game struct {
	backgroundColor   color.RGBA
	borderColor       color.RGBA
	players           map[color.RGBA]*Player
	world             map[worldPos]objectInWorld
	rotateSensitivity float64
	lastUpdate        time.Time
	xratio, yratio    float64
	logger            *gameLogger
	warmupsCount      int
	velocity          Velocity
	collisions        []collision
	state             state
	randomPos         randomPos
}

func NewGame(rotation float64, xratio, yratio float64, v Velocity, bg, border color.RGBA) (*game, error) {
	if rotation <= 0 {
		return nil, fmt.Errorf("rotation must be greater than zero")
	}
	if xratio <= 0 {
		return nil, fmt.Errorf("xratio must be greater than zero")
	}
	if yratio <= 0 {
		return nil, fmt.Errorf("yratio must be greater than zero")
	}
	// todo - set players head, check that they are not overlapping
	// todo - set borders, add logic to borders (you have an interface for that)

	logger, err := newLogger("logs")
	if err != nil {
		return nil, err
	}

	return &game{
		backgroundColor:   bg,
		players:           make(map[color.RGBA]*Player),
		world:             make(map[worldPos]objectInWorld),
		rotateSensitivity: rotation,
		lastUpdate:        time.Time{},
		xratio:            xratio,
		yratio:            yratio,
		logger:            logger,
		warmupsCount:      0,
		borderColor:       border,
		velocity:          v,
		collisions:        []collision{},
		state:             &initialState{},
		randomPos:         newRandomPos([]worldPos{newWorldPos(10, 10), newWorldPos(100, 150)}),
	}, nil
}

func (g *game) RegisterPlayer(newP Player) error {
	if len(g.players) == 2 {
		return fmt.Errorf("currently only 2 max players are allowed")
	}

	if _, ok := g.players[newP.color()]; ok {
		return fmt.Errorf("player with uid %s already exist", newP.uid())
	}

	newP.velocity = g.velocity
	g.players[newP.color()] = &newP
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.log("enteting draw loop")
	screen.Fill(g.backgroundColor)

	w := screen.Bounds().Dx()
	h := screen.Bounds().Dy()

	for pos, objInWorld := range g.world {
		xpix := int(float64(pos.x) * g.xratio)
		ypix := int(float64(pos.y) * g.yratio)

		if xpix < 0 || xpix >= w || ypix < 0 || ypix >= h {
			panic(fmt.Sprintf("invalid draw position: (%d, %d)", xpix, ypix))
		}

		screen.Set(xpix, ypix, objInWorld.color())
	}

	g.log("leaving draw loop")
}

func (g *game) Update() error {
	g.log("entering update loop")

	if g.warmupsCount < 1 {
		g.warmupsCount++
		return nil
	}

	if stateChanged, newState := g.state.update(g); stateChanged {
		g.state = newState
	}

	g.log("leaving update loop")
	return nil
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *game) Close() {
	g.logger.close()
}

func (g *game) log(msg string, args ...any) {
	g.logger.write(msg, args...)
}

func (g *game) logCollision(c collision) {
	objectsCount := len(c.objectsInvolved)
	first := c.objectsInvolved[0]

	if objectsCount == 1 {
		g.log("Collision at %v between %s and itself", c.pos, first)
	} else if objectsCount == 2 {
		second := c.objectsInvolved[1]
		g.log("Collision at %v between %s and %s", c.pos, first.uid(), second.uid())
	} else {
		panic("unknown collision case had occured")
	}
}

func (g *game) touchTimer() time.Duration {
	now := time.Now()

	if g.lastUpdate.IsZero() {
		g.lastUpdate = now
	}

	elapsed := now.Sub(g.lastUpdate)
	g.lastUpdate = now
	return elapsed
}

func (g *game) resetTimer() {
	g.lastUpdate = time.Time{}
}
