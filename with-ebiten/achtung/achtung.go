package achtung

import (
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/liornach/game-engine-ebiten/twodim"
)

var MaxRad = toRad(15)
var MinRad = toRad(0)

func toRad(deg float64) Rad {
	res := deg * math.Pi / 180
	return Rad(res)
}

type Rad twodim.Radians

type Player uuid.UUID

type Achtung struct {
	world       *World
	sensitivity Rad
}

func NewAchtung(r Rad) Achtung {
	if r < MinRad || r >= MaxRad {
		panic(InvalidRadError{
			r: r,
		})
	}

	return Achtung{
		sensitivity: r,
	}
}

func (a *Achtung) ApplyTime(t time.Duration) ([]error, []Collision) {
	errs, colls := a.world.ApplyTime(t)
	return errs, colls
}

func (a *Achtung) RotLeft(p Player) error {
	err := a.world.Rot(p, -a.sensitivity)
	return err
}

func (a *Achtung) RotRight(p Player) error {
	err := a.world.Rot(p, -a.sensitivity)
	return err
}
