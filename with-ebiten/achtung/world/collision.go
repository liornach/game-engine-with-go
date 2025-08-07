package world

import (
	"fmt"

	"github.com/liornach/game-engine-ebiten/achtung/ds"
)

type Collision struct {
	objects ds.Array[Uid]
	pos     WorldPos
}

func NewCollision(wp WorldPos, uids []Uid) Collision {
	objects := ds.NewArray[Uid]()

	for _, u := range uids {
		if objects.Contains(u) {
			panic(fmt.Errorf("object with uid %s was given more then once", u))
		}

		objects.Add(u)
	}

	return Collision{
		objects: objects,
		pos:     wp,
	}
}

func (c *Collision) AddObjectInvolved(u Uid) {
	if c.IsInvolved(u) {
		panic(fmt.Errorf("object with uid is %s already involved", u))
	}

	c.objects.Add(u)
}

func (c Collision) Objects() ds.Array[Uid] {
	return c.objects
}

func (c Collision) IsInvolved(u Uid) bool {
	return c.objects.Contains(u)
}
