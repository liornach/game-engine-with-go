package achtung

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/liornach/game-engine-ebiten/logic"
	"github.com/liornach/game-engine-ebiten/twodim"
)

var MaxRad = toRad(15)
var MinRad = toRad(0)

type Rad twodim.Radians

type InvalidRadError struct {
	r Rad
}

func (e InvalidRadError) Error() string {
	var err string
	if e.r < MinRad {
		err = fmt.Sprintf("radian must be greater than or equal to %f", MinRad)
	} else if e.r >= MaxRad {
		err = fmt.Sprintf("radian must be less than %f", MaxRad)
	} else {
		panic("no error found, consider debugging")
	}

	return err
}

type PlayerNotFoundError struct {
	p Player
}

func (e PlayerNotFoundError) Error() string {
	err := fmt.Sprintf("player with uuid %v was not found", e.p)
	return err
}

type Player uuid.UUID

type Achtung struct {
	w              logic.RectWorld
	players        map[Player]*Dot
	radSensitivity Rad
}

func NewAchtung(r Rad) Achtung {
	if r < MinRad || r >= MaxRad {
		panic(InvalidRadError{
			r: r,
		})
	}

	return Achtung{
		radSensitivity: r,
	}
}

func (a *Achtung) ApplyTime(t time.Duration) []error {
	var errs []error
	for pl, dot := range a.players {
		pos := dot.ApplyTime(t)
		_, err := a.w.Set(logic.Owner(pl), pos)
		if err != nil {
			errs = append(errs, err)
		} // maybe insted of errors, decide who wins?
	}

	return errs
}

func (a *Achtung) RotLeft(p Player) error {
	err := a.rot(p, -a.radSensitivity)
	return err
}

func (a *Achtung) RotRight(p Player) error {
	err := a.rot(p, a.radSensitivity)
	return err
}

func (a *Achtung) rot(p Player, r Rad) error {
	dot, ok := a.players[p]
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

func toRad(deg float64) Rad {
	res := deg * math.Pi / 180
	return Rad(res)
}
