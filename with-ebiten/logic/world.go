package logic

import (
	"fmt"

	"github.com/google/uuid"
)

type Owner uuid.UUID

type Pos struct {
	X, Y int
}

func (p Pos) String() string {
	return fmt.Sprintf("x : %d, y : %d", p.X, p.Y)
}

type World struct {
	min     Pos
	max     Pos
	wrldMap map[Pos]Owner
}

func NewRectWorld(min, max Pos) *World {
	return &World{
		min:     min,
		max:     max,
		wrldMap: map[Pos]Owner{},
	}
}

func (r *World) Set(o Owner, p Pos) (Owner, error) {
	var err error
	var owner Owner

	if !r.IsInWorld(p) {
		err = NotInWorldError{}
	} else if owner, ok := r.Owner(p); ok {
		if owner != o {
			err = NotFreeError{
				Pos:   p,
				Owner: o,
			}
		}
	} else {
		r.wrldMap[p] = o
		owner = o
	}

	return owner, err
}

func (r *World) Owner(p Pos) (Owner, bool) {
	o, ok := r.wrldMap[p]
	return o, ok
}

func (r *World) IsFree(p Pos) bool {
	_, ok := r.wrldMap[p]
	return !ok
}

func (r *World) IsInWorld(p Pos) bool {
	return p.X <= r.max.X &&
		p.X >= r.min.X &&
		p.Y <= r.min.Y &&
		p.Y >= r.min.Y
}

func (r *World) Map() ReadOnlyMap {
	return NewReadOnlyMap(r.wrldMap)
}
