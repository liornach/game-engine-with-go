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
	rotation                  float64
}

func (p *Player) setHead(pos WorldPos) {
	p.head.X = float64(pos.X)
	p.head.Y = float64(pos.Y)
}

func (p *Player) Head() WorldPos {
	return p.head.toWorldPos()
}

func (p *Player) TurnRightKey() Key {
	return p.turnRightKey
}

func (p *Player) TurnLeftKey() Key {
	return p.turnLeftKey
}

func (p *Player) Velocity() Velocity {
	return p.velocity
}

func NewPlayer(col Color, left, right ebiten.Key, rot float64) *Player {
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
		rotation:     rot,
	}
}

func (p *Player) EstimatePhysics(t time.Duration) playerPos {
	return p.head.estimatePhysics(t, p.velocity)
}

func (p *Player) ApplyPhysics(t time.Duration) {
	p.head = p.EstimatePhysics(t)
}

func (p *Player) rotateRight() {
	p.velocity.rotate(p.rotation)
}

func (p *Player) rotateLeft() {
	p.velocity.rotate(-p.rotation)
}

func (p *Player) Rotation() float64 {
	return p.rotation
}

type playerPos struct {
	X, Y float64
}

func (p playerPos) toWorldPos() WorldPos {
	return WorldPos{
		X: int(p.X),
		Y: int(p.Y),
	}
}

func (p playerPos) estimatePhysics(t time.Duration, v Velocity) playerPos {
	p.X += v.X * t.Seconds()
	p.Y += v.Y * t.Seconds()
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

func (p *Player) Uid() uid {
	return p.uniqueId
}

func (p *Player) color() color.RGBA {
	return p.col
}

func (p *Player) IsCollided(other ObjectInWorld, pos WorldPos) bool {
	// the suspected collider is this very player
	if other.Uid() == p.Uid() {
		// player head can not collide with itself
		if p.head.toWorldPos() == pos {
			return false
		}
	}

	return true
}
