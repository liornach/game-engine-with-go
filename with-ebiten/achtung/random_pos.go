package achtung

type randomPos struct {
	options []worldPos
	curIdx  int
}

func newRandomPos(options []worldPos) randomPos {
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