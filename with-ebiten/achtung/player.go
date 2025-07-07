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
	color                     color.RGBA
	uid                       string
	turnLeftKey, turnRightKey ebiten.Key
	head                      playerPos
	velocity                  velocity
}

func NewPlayer(col Color, left, right ebiten.Key) *Player {
	var rgba color.RGBA

	switch col {
	case Blue:
		rgba = color.RGBA{R: 0, G: 0, B: 225, A: 255}
	case Red:
		rgba = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case Green:
		rgba = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	default:
		panic(fmt.Errorf("unsupported color %s", col))
	}

	return &Player{
		color:        rgba,
		uid:          string(col),
		turnLeftKey:  left,
		turnRightKey: right,
		head:         playerPos{},
		velocity:     velocity{},
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
