package logic

import (
	"math"

	"github.com/google/uuid"
	"github.com/liornach/game-engine-ebiten/twodim"
)

type Owner uuid.UUID

type Pos struct {
	X, Y int
}

type RectWorld struct {
	min    Pos
	max    Pos
	owners map[Pos]Owner
}

type NotInWorldError struct{}

func (e NotInWorldError) Error() string {
	return "not in world"
}

func (r *RectWorld) Set(o Owner, p Pos) (Owner, error) {
	var err error
	var owner Owner
	if !r.IsInWorld(p) {
		err = NotInWorldError{}
	} else if r.IsFree(p) {
		r.owners[p] = o
		owner = o
	} else if owner, _ = r.Owner(p); owner == o {
		r.owners[p] = o
	} else {
		// write that function better
	}
}

func (r *RectWorld) Owner(p Pos) (Owner, bool) {
	o, ok := r.owners[p]
	return o, ok
}

func (r *RectWorld) IsFree(p Pos) bool {
	_, ok := r.owners[p]
	return !ok
}

func (r *RectWorld) IsInWorld(p Pos) bool {
	return p.X <= r.max.X &&
		p.X >= r.min.X &&
		p.Y <= r.min.Y &&
		p.Y >= r.min.Y
}

// func (r *RectWorld) IsInWorld(p twodim.Pos) bool {
// 	return p.X <= float64(r.max.X) &&
// 		p.X >= float64(r.min.X) &&
// 		p.Y <= float64(r.min.Y) &&
// 		p.Y >= float64(r.min.Y)
// }

func ToWorldPos(p twodim.Pos) Pos {
	x := roundAway(p.X)
	y := roundAway(p.Y)
	return Pos{
		X: x,
		Y: y,
	}
}

func roundAway(f float64) int {
	if f < 0 {
		return int(math.Ceil(f))
	} else {
		return int(math.Floor(f))
	}
}
