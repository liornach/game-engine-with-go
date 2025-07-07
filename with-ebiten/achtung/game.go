package achtung

import (
	"errors"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type worldPos struct {
	x, y int
}

type game struct {
	backgroundColor   color.RGBA
	players           []*Player
	world             map[worldPos]*Player
	rotateSensitivity float64
	lastUpdate        time.Time
	xratio, yratio    float64
	logger            *gameLogger
}

func NewGame(players []*Player, rotation float64, xratio, yratio float64) (*game, error) {

	pos := playerPos{
		x: 10,
		y: 20,
	}

	vel := velocity{
		x: 3,
		y: 3,
	}

	world := map[worldPos]*Player{}

	for i, pi := range players {
		for _, pj := range players[i+1:] {
			if pj.uid == pi.uid {
				return nil, fmt.Errorf("duplication of player with a uid %s", pi.uid)
			}
		}

		players[i].velocity = vel
		players[i].head = pos
		world[players[i].head.toWorldPos()] = players[i]
		pos.x += 10
		pos.y += 20
	}

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

	background := color.RGBA{0, 0, 0, 1}

	return &game{
		backgroundColor:   background,
		players:           players,
		world:             world,
		rotateSensitivity: rotation,
		lastUpdate:        time.Now(),
		xratio:            xratio,
		yratio:            yratio,
		logger:            logger,
	}, nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)

	w := screen.Bounds().Dx()
	h := screen.Bounds().Dy()

	for pos, player := range g.world {
		xpix := int(float64(pos.x) * g.xratio)
		ypix := int(float64(pos.y) * g.yratio)

		if xpix < 0 || xpix >= w || ypix < 0 || ypix >= h {
			panic(fmt.Sprintf("invalid draw position: (%d, %d)", xpix, ypix))
		}

		screen.Set(xpix, ypix, player.color)
	}
}

func (g *game) Update() error {
	now := time.Now()
	elapsed := now.Sub(g.lastUpdate)
	g.lastUpdate = now
	colls := 0

	for i := range g.players {
		curP := g.players[i]
		newHead := curP.estimatePhysics(elapsed)
		nextWorldPos := newHead.toWorldPos()
		// now, need to debug
		existing, ok := g.world[nextWorldPos]

		if ok && (existing.uid != curP.uid || nextWorldPos != curP.head.toWorldPos()) {
			colls++
			g.logger.write("Collision at %v between %s and %s\n", nextWorldPos, existing.uid, curP.uid)
		} else {
			curP.head = newHead
			g.world[nextWorldPos] = curP

			if ebiten.IsKeyPressed(curP.turnLeftKey) {
				curP.rotate(-g.rotateSensitivity)
			}

			if ebiten.IsKeyPressed(curP.turnRightKey) {
				curP.rotate(g.rotateSensitivity)
			}
		}
	}

	if colls != 0 {
		fmt.Scanln()
		return errors.New("collision occured")
	}

	return nil
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *game) Close() {
	g.logger.close()
}
