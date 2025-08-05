package achtung

import "github.com/hajimehoshi/ebiten/v2"

type Key = ebiten.Key

type inputHandler struct {
}

func NewInputHandler() inputHandler {
	return inputHandler{}
}

func (inputHandler) IsKeyPressed(k Key) bool {
	return ebiten.IsKeyPressed(k)
}
