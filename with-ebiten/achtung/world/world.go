package world

type Uid = string

type World struct {
	world map[WorldPos]Uid
}

func NewWorld() World {
	return World{
		world: map[WorldPos]Uid{},
	}
}

func (w World) At(wp WorldPos) (Uid, bool) {
	ret, ok := w.world[wp]
	return ret, ok
}

func (w *World) Set(o Uid, wp WorldPos) bool {
	if _, ok := w.At(wp); ok {
		return false
	}

	w.world[wp] = o
	return true
}
