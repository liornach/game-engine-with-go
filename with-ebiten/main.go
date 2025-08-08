package main

import (
	"fmt"
	"github.com/liornach/game-engine-with-go/achtung"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fmt.Println("hello world")
	game := achtung.Game{}
	ebiten.RunGame(&game)
}
