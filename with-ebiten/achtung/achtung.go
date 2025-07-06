package achtung

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type worldPos struct {
	x, y int
}

type playerPos struct {
	x, y float64
}

func (p playerPos) toWorldPos() worldPos {
	return worldPos{
		x: int(p.x),
		y: int(p.y),
	}
}

func (p playerPos) estimatePhysics(t time.Duration, v velocity) playerPos {
	p.x += v.x * t.Seconds()
	p.y += v.y * t.Seconds()
	return p
}

type velocity struct {
	x, y float64
}

func (v *velocity) rotate(rad float64) {
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	oldx, oldy := v.x, v.y
	v.x = oldx*cos - oldy*sin
	v.y = oldx*sin - oldy*cos
}

type player struct {
	color                     color.RGBA
	uid                       string
	turnLeftKey, turnRightKey ebiten.Key
	head                      playerPos
	velocity                  velocity
}

func NewPlayer(color color.RGBA, uid string, left, right ebiten.Key) *player {
	return &player{
		color:        color,
		uid:          uid,
		turnLeftKey:  left,
		turnRightKey: right,
		head:         playerPos{},
		velocity:     velocity{},
	}
}

func (p *player) estimatePhysics(t time.Duration) playerPos {
	return p.head.estimatePhysics(t, p.velocity)
}

func (p *player) rotate(rad float64) {
	p.velocity.rotate(rad)
}

type collision struct {
	original *player
	collider *player
	pos      worldPos
}

type game struct {
	backgroundColor     color.RGBA
	players             []*player
	world               map[worldPos]*player
	rotateSensitivity   float64
	lastUpdate          time.Time
	xratio, yratio      float64
	isHandlingCollision bool
}

func NewGame(players []*player, rotation float64, background color.RGBA, xratio, yratio float64) (*game, error) {
	for i, pi := range players {
		for _, pj := range players[i+1:] {
			if pj.uid == pi.uid {
				return nil, fmt.Errorf("duplication of player with a uid %s", pi.uid)
			}
		}

		players[i].velocity = velocity{}
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
	return &game{
		backgroundColor:     background,
		players:             players,
		world:               map[worldPos]*player{},
		rotateSensitivity:   rotation,
		lastUpdate:          time.Now(),
		xratio:              xratio,
		yratio:              yratio,
		isHandlingCollision: false,
	}, nil
}

func (g *game) Draw(screen ebiten.Image) {
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
	if g.isHandlingCollision {
		// todo add collision handling
		g.isHandlingCollision = false
		return nil
	}

	now := time.Now()
	elapsed := now.Sub(g.lastUpdate)
	g.lastUpdate = now

	var colls []collision

	for i := range g.players {
		curP := g.players[i]
		newHead := curP.estimatePhysics(elapsed)
		nextWorldPos := newHead.toWorldPos()

		if existing, ok := g.world[nextWorldPos]; ok {
			colls = append(colls, collision{
				original: existing,
				collider: curP,
				pos:      nextWorldPos,
			})
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

	if len(colls) != 0 {
		for _, c := range colls {
			fmt.Printf("Collision at %v between %s and %s\n", c.pos, c.original.uid, c.collider.uid)
		}
		g.isHandlingCollision = true
	}

	return nil
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
