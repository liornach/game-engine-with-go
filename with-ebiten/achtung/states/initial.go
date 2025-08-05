package states

import "github.com/liornach/game-engine-ebiten/achtung"

type worldPos = achtung.WorldPos

type randomPos struct {
	options []worldPos
	curIdx  int
}

func NewRandomPos(options []worldPos) randomPos {
	if len(options) <= 0 {
		panic("options length must be at least 1")
	}

	return randomPos{
		options: options,
		curIdx:  -1,
	}
}

func (rp *randomPos) next() worldPos {
	rp.curIdx++
	if rp.curIdx == len(rp.options) {
		panic("already used all options")
	}

	return rp.options[rp.curIdx]
}

type intial struct {
	randomPos randomPos
}

func NewInitialState(rp randomPos) intial {
	return intial{
		randomPos: rp,
	}
}

func (i intial) Update(g *game) (bool, state) {
	for _, p := range g.Players {
		randPos := i.randomPos.next()
		g.SetPlayerHead(p, randPos)
	}

	return true, runningState{}
}
