package states

import (
	"fmt"

	"github.com/liornach/game-engine-ebiten/achtung"
)

type collision = achtung.Collision
type game = achtung.Game
type state = achtung.State

type HandlingCollisions struct {
	col        collision
	eliminated []*player
}

func newHandlingCollisions(colls []collision) HandlingCollisions {
	assertCollisionsCount(colls)
	state := HandlingCollisions{
		col: colls[0],
	}

	return state
}

func (s HandlingCollisions) Update(g *game) (bool, state) {
	newState := s.handlingCollisons(g)
	_, isHandlingCollisionState := newState.(HandlingCollisions)
	if !isHandlingCollisionState {
		g.ClearCollisions()
	}

	return !isHandlingCollisionState, newState
}

func (s *HandlingCollisions) handlingCollisons(g *game) state {
	if s.IsSelfCollision() {
		return s.handleSingleCollision(g)
	}

	if s.IsTwoObjectsCollision() {
		return s.handleTwoObjectsCollision(g)
	}

	panic("unsupported amount of collided objects")
}

func (s *HandlingCollisions) handleSingleCollision(g *game) state {
	c := s.col

	if len(c.Objects) != 1 {
		panic("not a single-object collision")
	}

	objectInPosition, ok := g.PosOwner(c.Pos)
	if !ok {
		err := fmt.Errorf("no object seems to be in position %v", c.Pos)
		panic(err)
	}

	collidedObject := c.Objects[0]
	if objectInPosition.Uid() != collidedObject.Uid() {
		err := fmt.Errorf("object in position belongs to another object")
		panic(err)
	}

	return s.newEliminated(collidedObject)
}

func (s *HandlingCollisions) handleTwoObjectsCollision(g *game) state {
	c := s.col

	if len(c.Objects) != 2 {
		panic("not a two-objects collision")
	}

	first := c.Objects[0]
	second := c.Objects[1]
	objectInPosition, ok := g.PosOwner(c.Pos)

	if !ok {
		err := fmt.Errorf("no object seems to be in position %v", c.Pos)
		panic(err)
	}

	objInPositionUid := objectInPosition.Uid()

	if objInPositionUid == first.Uid() {
		return s.newEliminated(second)
	}

	if objInPositionUid == second.Uid() {
		return s.newEliminated(first)
	}

	panic("non of the involved objects holds the position of collision")
}

func (s *HandlingCollisions) newEliminated(o objectInWorld) elimanted {
	if p, ok := o.(*player); ok {
		return newEliminated(p)
	}

	panic("unknown object involved in collision with itself")
}

func (state *HandlingCollisions) IsSelfCollision() bool {
	return len(state.col.Objects) == 1
}

func (state *HandlingCollisions) IsTwoObjectsCollision() bool {
	return len(state.col.Objects) == 2
}

func assertCollisionsCount(c []collision) {
	if len(c) <= 0 {
		panic("wrong state, seems like there are no collisions")
	}

	if len(c) > 1 {
		panic("can't handle more than one collision at a time")
	}
}
