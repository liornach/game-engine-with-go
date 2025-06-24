package logic

import (
	"github.com/google/uuid"
)

type Owner uuid.UUID

type Pos struct {
	X, Y int
}

type RectWorld struct {
	min     Pos
	max     Pos
	wrldMap map[Pos]Owner
}

type NotInWorldError struct{}
type NotFreeError struct{}

func (e NotInWorldError) Error() string {
	return "not in world"
}

func (e NotFreeError) Error() string {
	return "not free"
}

func (r *RectWorld) Set(o Owner, p Pos) (Owner, error) {
	var err error
	var owner Owner

	if !r.IsInWorld(p) {
		err = NotInWorldError{}
	} else if owner, ok := r.Owner(p); ok {
		if owner != o {
			err = NotFreeError{}
		}
	} else {
		r.wrldMap[p] = o
		owner = o
	}

	return owner, err
}

func (r *RectWorld) Owner(p Pos) (Owner, bool) {
	o, ok := r.wrldMap[p]
	return o, ok
}

func (r *RectWorld) IsFree(p Pos) bool {
	_, ok := r.wrldMap[p]
	return !ok
}

func (r *RectWorld) IsInWorld(p Pos) bool {
	return p.X <= r.max.X &&
		p.X >= r.min.X &&
		p.Y <= r.min.Y &&
		p.Y >= r.min.Y
}

func (r *RectWorld) Map() ReadOnlyMap {
	return NewReadOnlyMap(r.wrldMap)
}
