package display

import "github.com/liornach/game-engine-ebiten/logic"

type Display struct {
	xRatio, yRatio float64
}

func NewDisplay(xRatio, yRatio float64) *Display {
	return &Display{
		xRatio: xRatio,
		yRatio: yRatio,
	}
}

func Display(w *logic.RectWorld, ebichan) {

}
