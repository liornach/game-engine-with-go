package junk

type junk struct {
	i int
}

func NewJunk() *junk {
	return &junk{}
}

func PublicFunc(j junk) {
}

