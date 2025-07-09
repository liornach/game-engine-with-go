package achtung

import (
	"errors"
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

// type objectInWorld struct {
// 	isPlayer, isBorder bool
// 	color              color.RGBA
// }

// func newObjectInWorld(isPlayer, isBorder bool, color color.RGBA) *objectInWorld {
// 	return &objectInWorld{
// 		isPlayer: isPlayer,
// 		isBorder: isBorder,
// 		color:    color,
// 	}
// }

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
	}, nil
}

func (g *game) RegisterPlayer(newP Player) error {
	if len(g.players) == 2 {
		return fmt.Errorf("currently only 2 max players are allowed")
	}

	if _, ok := g.players[newP.color]; ok {
		return fmt.Errorf("player with color %s already exist", newP.color)
	}

	newP.velocity = g.velocity
	g.players[newP.color] = &newP
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

		screen.Set(xpix, ypix, objInWorld.color)
	}

	g.log("leaving draw loop")
}

func (g *game) Update() error {
	g.log("entering update loop")

	if g.warmupsCount < 1 {
		g.warmupsCount++
		return nil
	}

	now := time.Now()

	if g.lastUpdate.IsZero() {
		g.lastUpdate = now
		g.log("last update time was being set to now")
	}

	elapsed := now.Sub(g.lastUpdate)
	g.lastUpdate = now
	colls := 0

	for i, curPlayer := range g.players {
		newHead := curPlayer.estimatePhysics(elapsed)
		nextWorldPos := newHead.toWorldPos()

		if existObjInWorld, ok := g.world[nextWorldPos]; ok {
			if existObjInWorld.isCollided(curPlayer, nextWorldPos) {
				colls++
				g.logCollision(existObjInWorld, curPlayer, nextWorldPos)
			}
		} else {
			curPlayer.head = newHead // some logic shit here
			g.world[nextWorldPos] = curPlayer
			g.log("player %s was set in %v", curPlayer.uid, nextWorldPos)

			isVelChanged := false
			prevVel := curPlayer.velocity

			if ebiten.IsKeyPressed(curPlayer.turnLeftKey) {
				curPlayer.rotate(-g.rotateSensitivity)
				isVelChanged = true
			}

			if ebiten.IsKeyPressed(curPlayer.turnRightKey) {
				curPlayer.rotate(g.rotateSensitivity)
				isVelChanged = true
			}

			if isVelChanged {
				g.log("velocity of %s changed from %v to %v", curPlayer.uid, prevVel, curPlayer.velocity)
			}
		}
	}

	if colls != 0 {
		fmt.Scanln()
		return errors.New("collision occured")
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

func (g *game) logCollision(exist objectInWorld, collider *Player, pos worldPos) {
	g.log("Collision at %v between %s (existing) and %s (collider)", pos, exist.uid(), collider.uid())
	g.log("Current player head : %v", collider.head.toWorldPos())
	g.log("World Position : %v", pos)
}

func (g *game) movePlayerHead(p *Player, newHead playerPos, wpos worldPos) {
	p.head = newHead
	g.world[wpos] = p
	g.log("%s player head was set in %v", p.uid, wpos)
}
