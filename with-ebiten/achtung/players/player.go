package players

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/liornach/game-engine-ebiten/achtung"
)

type Color string
type PlayerPos = achtung.PlayerPos

const (
	Blue  Color = "Blue"
	Green Color = "Green"
	Red   Color = "Red"
)

type Uid = string
type Key = ebiten.Key

type Player struct {
	col                       color.RGBA
	turnLeftKey, turnRightKey ebiten.Key
	//head                      PlayerPos
	//head physics.StateVector
	//velocity Velocity
	uid Uid
	//rotation float64
}

// func (p *Player) setHead(pos WorldPos) {
// 	p.head.X = float64(pos.X)
// 	p.head.Y = float64(pos.Y)
// }

// func (p *Player) SetHead(pos PlayerPos) {
// 	p.head.X = float64(pos.X)
// 	p.head.Y = float64(pos.Y)
// }

// func (p *Player) Head() WorldPos {
// 	return p.head.toWorldPos()
// }

// func (p *Player) Head() PlayerPos {
// 	return p.head
// }

func (p *Player) TurnRightKey() Key {
	return p.turnRightKey
}

func (p *Player) TurnLeftKey() Key {
	return p.turnLeftKey
}

// func (p *Player) Velocity() Velocity {
// 	return p.velocity
// }

// func NewPlayer(col Color, left, right ebiten.Key, rot float64, v Velocity) *Player {
// 	var rgba color.RGBA

// 	switch col {
// 	case Blue:
// 		rgba = color.RGBA{R: 0, G: 0, B: 255, A: 255}
// 	case Red:
// 		rgba = color.RGBA{R: 255, G: 0, B: 0, A: 255}
// 	case Green:
// 		rgba = color.RGBA{R: 0, G: 255, B: 0, A: 255}
// 	default:
// 		panic(fmt.Errorf("unsupported color %s", col))
// 	}

// 	return &Player{
// 		col:          rgba,
// 		turnLeftKey:  left,
// 		turnRightKey: right,
// 		head:         PlayerPos{},
// 		velocity:     Velocity{},
// 		uid:          Uid(col),
// 		rotation:     rot,
// 	}
// }

func (p *Player) EstimateHeadFutureLocation(t time.Duration) PlayerPos {
	pos := p.head
	v := p.velocity

	pos.X += v.X * t.Seconds()
	pos.Y += v.Y * t.Seconds()

	return pos
}

func (p *Player) ApplyPhysics(t time.Duration) PlayerPos {
	p.head = p.EstimateHeadFutureLocation(t)
	return p.head
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

// func (p playerPos) toWorldPos() WorldPos {
// 	return WorldPos{
// 		X: int(p.X),
// 		Y: int(p.Y),
// 	}
// }

// func (pos PlayerPos) estimatePhysics(t time.Duration, v Velocity) PlayerPos {
// 	pos.X += v.X * t.Seconds()
// 	pos.Y += v.Y * t.Seconds()
// 	return pos
// }

// func (v *Velocity) rotate(rad float64) {
// 	cos := math.Cos(rad)
// 	sin := math.Sin(rad)
// 	oldx, oldy := v.X, v.Y
// 	v.X = oldx*cos - oldy*sin
// 	v.Y = oldx*sin + oldy*cos
// }

func (p *Player) Uid() Uid {
	return p.uid
}

func (p *Player) Color() color.RGBA {
	return p.col
}

// func (p *Player) IsCollided(other ObjectInWorld, pos WorldPos) bool {
// 	// the suspected collider is this very player
// 	if other.Uid() == p.Uid() {
// 		// player head can not collide with itself
// 		if p.head.toWorldPos() == pos {
// 			return false
// 		}
// 	}

// 	return true
// }
