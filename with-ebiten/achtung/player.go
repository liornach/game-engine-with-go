package achtung

type Color string

const (
	Blue Color = "Blue"
	Green Color = "Green"
	Red Color = "Red"
	Orange Color = "Orange"
)

type Player struct {
	x, y, vx, vy float64
	color Color
}
