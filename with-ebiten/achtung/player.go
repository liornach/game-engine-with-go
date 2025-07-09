package achtung

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Color string

const (
	Blue  Color = "Blue"
	Green Color = "Green"
	Red   Color = "Red"
)

type Player struct {
	col                       color.RGBA
	turnLeftKey, turnRightKey ebiten.Key
	head                      playerPos
	velocity                  Velocity
	uniqueId                  uid
}

func NewPlayer(col Color, left, right ebiten.Key) *Player {
	var rgba color.RGBA

	switch col {
	case Blue:
		rgba = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	case Red:
		rgba = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case Green:
		rgba = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	default:
		panic(fmt.Errorf("unsupported color %s", col))
	}

	return &Player{
		col:          rgba,
		turnLeftKey:  left,
		turnRightKey: right,
		head:         playerPos{},
		velocity:     Velocity{},
		uniqueId:     uid(col),
	}
}

func (p *Player) estimatePhysics(t time.Duration) playerPos {
	return p.head.estimatePhysics(t, p.velocity)
}

func (p *Player) rotate(rad float64) {
	p.velocity.rotate(rad)
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

func (p playerPos) estimatePhysics(t time.Duration, v Velocity) playerPos {
	p.x += v.X * t.Seconds()
	p.y += v.Y * t.Seconds()
	return p
}

type Velocity struct {
	X, Y float64
}

func (v *Velocity) rotate(rad float64) {
	cos := math.Cos(rad)
	sin := math.Sin(rad)
	oldx, oldy := v.X, v.Y
	v.X = oldx*cos - oldy*sin
	v.Y = oldx*sin + oldy*cos
}

func (p *Player) uid() uid {
	return p.uniqueId
}

func (p *Player) color() color.RGBA {
	return p.col
}

func (p *Player) isCollided(other objectInWorld, pos worldPos) bool {
	// the suspected collider is this very player
	if other.uid() == p.uid() {
		// player head can not collide with itself
		if p.head.toWorldPos() == pos {
			return false
		}
	}

	return true
}
