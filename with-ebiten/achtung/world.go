package achtung

import (
	"errors"
	"time"

	"github.com/liornach/game-engine-ebiten/logic"
)

type World struct {
	World   *logic.RectWorld
	Players map[Player]*Dot
}

func NewWorld(positions map[Player]Pos, min, max Pos) (*World, error) {
	wrld := &World{
		World:   logic.NewRectWorld(logic.Pos(min), logic.Pos(max)),
		Players: map[Player]*Dot{},
	}

	for pl, pos := range positions {
		_, err := wrld.set(pl, pos)
		if err != nil {
			return wrld, err
		}
	}

	return wrld, nil
}

func (w *World) Rot(p Player, r Rad) error {
	dot, ok := w.Players[p]
	var e error
	if !ok {
		e = PlayerNotFoundError{
			p: p,
		}
	} else {
		dot.Rot(r)
	}

	return e
}

func (w *World) set(pl Player, pos Pos) (Player, error) {
	o, err := w.World.Set(logic.Owner(pl), logic.Pos(pos))
	return Player(o), err
}

func (w *World) ApplyTime(t time.Duration) ([]error, []Collision) {
	var errs []error
	var colls []Collision

	for pl, dot := range w.Players {
		pos := dot.ApplyTime(t)
		o, err := w.set(pl, pos)
		if err != nil {
			if errors.Is(err, logic.NotFreeError{}) {
				colls = append(colls, *NewCollision(o, pl, pos))
			} else {
				errs = append(errs, err)
			}
		}
	}

	return errs, colls
}

func (w *World) Map() logic.ReadOnlyMap {
	return w.World.Map()
}
