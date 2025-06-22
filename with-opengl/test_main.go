package main

import (
	"example.com/go-engine/ggl"
)

func init() {
}

func main() {
	g := ggl.NewOpengl(800, 500, "Lior Nachmias 2")
	defer g.Term()

	g.ViewPort(0, 0, 800, 500) // currently same as window size

}

func FbSizeClbk()
