package twodim

import (
	"math"
	"time"
)

type Radians float64

type Speed float64

type MvDot struct {
	Pos       Pos
	Speed     Speed
	Direction Radians
}

func (m *MvDot) Move(dt time.Duration) Pos {
	seconds := dt.Seconds()

	distance := float64(m.Speed) * seconds

	dx := distance * math.Cos(float64(m.Direction))
	dy := distance * math.Sin(float64(m.Direction))

	m.Pos.X += dx
	m.Pos.Y += dy

	return m.Pos
}

func (m *MvDot) Rotate(r Radians) {
	m.Direction += r
}
