package achtung

import "github.com/liornach/game-engine-ebiten/logic"

type Pos logic.Pos

type Collision struct {
	Owner    Player
	Collider Player
	Pos      Pos
}

func NewCollision(o, c Player, p Pos) *Collision {
	return &Collision{
		Owner:    o,
		Collider: c,
		Pos:      p,
	}
}
