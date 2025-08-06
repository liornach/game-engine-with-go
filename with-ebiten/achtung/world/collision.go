package world

import (
	"fmt"

	"github.com/liornach/game-engine-ebiten/achtung/ds"
)

type Collision struct {
	objectsInvolved ds.Array[Uid]
	pos             WorldPos
}

func NewCollision(wp WorldPos, uids []Uid) Collision {
	objectsInvolved := ds.NewArray[Uid]()

	for _, u := range uids {
		if objectsInvolved.Contains(u) {
			panic(fmt.Errorf("object with uid %s was given more then once", u))
		}

		objectsInvolved.Add(u)
	}

	return Collision{
		objectsInvolved: objectsInvolved,
		pos:             wp,
	}
}

func (c *Collision) AddObjectInvolved(u Uid) {
	if c.IsInvolved(u) {
		panic(fmt.Errorf("object with uid is %s already involved", u))
	}

	c.objectsInvolved.Add(u)
}

func (c Collision) ObjectsInvolved() ds.Array[Uid] {
	return c.objectsInvolved
}

func (c Collision) IsInvolved(u Uid) bool {
	return c.objectsInvolved.Contains(u)
}
