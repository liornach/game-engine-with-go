package world

type World struct {
	world map[WorldPos]*WorldObject
}

func NewWorld() World {
	return World{
		world: map[WorldPos]*WorldObject{},
	}
}

func (w World) At(wp WorldPos) (WorldObject, bool) {
	ret, ok := w.world[wp]
	return *ret, ok
}

func (w World) 
