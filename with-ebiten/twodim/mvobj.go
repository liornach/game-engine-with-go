package twodim

import "time"

type Vec2 struct {
	X float64
	Y float64
}

type (
	Position Vec2
	Velocity Vec2
)

type MvObj2 struct {
	Pos Position
	Vel Velocity
}

func (o MvObj2) setX(x float64) {
	o.Pos.X = x
}

func (o MvObj2) setY(y float64) {
	o.Pos.Y = y
}

func (o MvObj2) vX() float64 {
	return o.Vel.X
}

func (o MvObj2) vY() float64 {
	return o.Vel.Y
}

func (o MvObj2) posX() float64 {
	return o.Pos.X
}

func (o MvObj2) posY() float64 {
	return o.Pos.Y
}

func (o MvObj2) moveX(t time.Duration) {
	o.setX(t.Seconds()*o.vX() + o.posX())
}

func (o MvObj2) moveY(t time.Duration) {
	o.setY(t.Seconds()*o.vY() + o.posY())
}

func (o MvObj2) Move(t time.Duration) {
	o.moveX(t)
	o.moveY(t)
}
