package achtung

import (
	"math"
	"time"

	"github.com/liornach/game-engine-ebiten/logic"
	"github.com/liornach/game-engine-ebiten/twodim"
)

type Dot struct {
	dot twodim.MvDot
}

func (d *Dot) ApplyTime(t time.Duration) logic.Pos {
	p := d.dot.ApplyTime(t)
	lp := toLogicPos(p)
	return lp
}

func (d *Dot) Rot(r Rad) {
	d.dot.Rot(twodim.Radians(r))
}

func toLogicPos(p twodim.Pos) logic.Pos {
	return logic.Pos{
		X: round(p.X),
		Y: round(p.Y),
	}
}

func round(f float64) int {
	return int(math.Round(f))
}
