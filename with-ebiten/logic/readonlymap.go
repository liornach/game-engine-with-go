package logic

type ReadOnlyMap struct {
	data map[Pos]Owner
}

func NewReadOnlyMap(m map[Pos]Owner) ReadOnlyMap {
	return ReadOnlyMap{data: m}
}

func (r ReadOnlyMap) Range(fn func(Pos, Owner)) {
	for k, v := range r.data {
		fn(k, v)
	}
}
