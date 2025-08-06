package achtung

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/achtung/world"
)

type State interface {
	Update(g *Game) (bool, State)
}

type PlayerPos struct {
	X, Y float64
}

func (p PlayerPos) toWorldPos() WorldPos {
	return WorldPos{
		X: int(p.X),
		Y: int(p.Y),
	}
}

type Player interface {
	Head() PlayerPos
	SetHead(PlayerPos)
	EstimateHeadFutureLocation(time.Duration) PlayerPos
	ApplyPhysics(time.Duration) PlayerPos
	Uid() Uid
	Color() color.RGBA
}

type Uid = string

// type ObjectInWorld interface {
// 	//IsCollided(other ObjectInWorld, pos WorldPos) bool
// 	Uid() Uid
// 	Color() color.RGBA
// }

func (wp WorldPos) toPlayerPos() PlayerPos {
	return PlayerPos{
		X: float64(wp.X),
		Y: float64(wp.Y),
	}
}

// func (c *Collision) AddObject(o ObjectInWorld) {
// 	if ()
// 	c.Objects[o.Uid()] = o
// }

type Score struct {
	Player *Player
	Score  int
}

type Game struct {
	backgroundColor color.RGBA
	borderColor     color.RGBA
	Players         []Player
	world           world.World
	lastUpdate      time.Time
	xratio, yratio  float64
	logger          *gameLogger
	warmupsCount    int
	//velocity          Velocity
	//collisions        []Collision
	state             State
	inputHandler      inputHandler
	scores            []Score
	eliminatedPlayers []Uid
}

func (g *Game) World() world.World {
	return g.world
}

// func (g *Game) PosOwner(wp WorldPos) (ObjectInWorld, bool) {
// 	o, ok := g.world[wp]
// 	return o, ok
// }

func (g *Game) EstimatedNextWorldPos(p Player, elapsed time.Duration) WorldPos {
	return p.EstimateHeadFutureLocation(elapsed).toWorldPos()
}

func (g *Game) PlayerHead(p Player) WorldPos {
	return p.Head().toWorldPos()
}

func (g *Game) SetPlayerHead(p Player, wp WorldPos) {
	playerPos := wp.toPlayerPos()
	p.SetHead(playerPos)
}

func (g *Game) EstimateCollisions(t time.Duration) []Collision {
	collisions := make(map[WorldPos]*Collision)

	for _, p := range g.Players {
		if g.IsEliminated(p.Uid()) {
			continue
		}

		futureHead := p.EstimateHeadFutureLocation(t).toWorldPos()
		if futureHead == p.Head().toWorldPos() {
			continue
		}

		if g.IsPosFree(futureHead) {
			if col, ok := collisions[futureHead]; ok {
				col.Objects = append(col.Objects, p)
			}
		}

	}
}

func (g *Game) ApplyPhysicsToPlayer(p Player, t time.Duration) {
	p.ApplyPhysics(t)
}

func (g *Game) IsPlayerHeadAt(p *Player, wp WorldPos) bool {
	return g.PlayerHead(p) == wp
}

func (g *Game) IsPosOwnedBy(wp WorldPos, o ObjectInWorld) bool {
	exist, ok := g.PosOwner(wp)
	if !ok {
		return false
	}

	return exist.Uid() == o.Uid()
}

func (g *Game) IsPosFree(wp WorldPos) bool {
	_, ok := g.PosOwner(wp)
	return !ok
}

func (g *Game) SetPosOwner(wp WorldPos, o ObjectInWorld) {
	g.world[wp] = o
}

func NewGame(xratio, yratio float64, bg, border color.RGBA, initialState State) (*Game, error) {
	if xratio <= 0 {
		return nil, fmt.Errorf("xratio must be greater than zero")
	}
	if yratio <= 0 {
		return nil, fmt.Errorf("yratio must be greater than zero")
	}

	logger, err := newLogger("logs")
	if err != nil {
		return nil, err
	}

	return &Game{
		backgroundColor:   bg,
		Players:           []Player{},
		world:             make(map[WorldPos]ObjectInWorld),
		lastUpdate:        time.Time{},
		xratio:            xratio,
		yratio:            yratio,
		logger:            logger,
		warmupsCount:      0,
		borderColor:       border,
		collisions:        []Collision{},
		state:             initialState,
		inputHandler:      inputHandler{},
		scores:            []Score{},
		eliminatedPlayers: []Uid{},
	}, nil
}

func (g *Game) IsRegistered(puid Uid) bool {
	for _, registered := range g.Players {
		if registered.Uid() == puid {
			return true
		}
	}

	return false
}

func (g *Game) IsEliminated(p Uid) bool {
	if !g.IsRegistered(p) {
		panic("player is not registered")
	}

	for _, eliminated := range g.eliminatedPlayers {
		if eliminated == p {
			return true
		}
	}

	return false
}

func (g *Game) RegisterPlayer(newP Player) error {
	if len(g.Players) == 2 {
		return fmt.Errorf("currently only 2 max players are allowed")
	}

	if g.IsRegistered(newP.Uid()) {
		err := fmt.Errorf("player with uid %s already exist", newP.Uid())
		panic(err)
	}

	for _, existP := range g.Players {
		if existP.color() == newP.color() {
			return fmt.Errorf("player with color %v already exist", existP.color())
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.log("enteting draw loop")
	screen.Fill(g.backgroundColor)

	w := screen.Bounds().Dx()
	h := screen.Bounds().Dy()

	for pos, objInWorld := range g.world {
		xpix := int(float64(pos.X) * g.xratio)
		ypix := int(float64(pos.Y) * g.yratio)

		if xpix < 0 || xpix >= w || ypix < 0 || ypix >= h {
			panic(fmt.Sprintf("invalid draw position: (%d, %d)", xpix, ypix))
		}

		screen.Set(xpix, ypix, objInWorld.color())
	}

	g.log("leaving draw loop")
}

func (g *Game) Update() error {
	g.log("entering update loop")

	if g.warmupsCount < 1 {
		g.warmupsCount++
		return nil
	}

	if stateChanged, newState := g.state.Update(g); stateChanged {
		g.state = newState
	}

	g.log("leaving update loop")
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Close() {
	g.logger.close()
}

func (g *Game) log(msg string, args ...any) {
	g.logger.write(msg, args...)
}

func (g *Game) logCollision(c Collision) {
	objectsCount := len(c.Objects)
	first := c.Objects[0]

	if objectsCount == 1 {
		g.log("Collision at %v between %s and itself", c.Pos, first)
	} else if objectsCount == 2 {
		second := c.Objects[1]
		g.log("Collision at %v between %s and %s", c.Pos, first.Uid(), second.Uid())
	} else {
		panic("unknown collision case had occured")
	}
}

func (g *Game) TouchTimer() time.Duration {
	now := time.Now()

	if g.lastUpdate.IsZero() {
		g.lastUpdate = now
	}

	elapsed := now.Sub(g.lastUpdate)
	g.lastUpdate = now
	return elapsed
}

func (g *Game) ResetTimer() {
	g.lastUpdate = time.Time{}
}

func (g *Game) AddCollision(c Collision) {
	g.collisions = append(g.collisions, c)
	g.logCollision(c)
}

func (g *Game) CollisionsCount() int {
	return len(g.collisions)
}

func (g *Game) Collisions() []Collision {
	return g.collisions
}

func (g *Game) ClearCollisions() {
	g.collisions = []Collision{}
}

func (g *Game) InputHandler() inputHandler {
	return g.inputHandler
}

func (g *Game) HandlePlayerKeys(p *Player) {
	ih := g.InputHandler()
	if ih.IsKeyPressed(p.TurnRightKey()) {
		p.rotateRight()
		g.log("player rotate right")
	}

	if ih.IsKeyPressed(p.TurnLeftKey()) {
		p.rotateLeft()
		g.log("player rotate left")
	}
}
