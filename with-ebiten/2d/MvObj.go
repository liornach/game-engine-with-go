package twodim

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

func (o MvObj2) moveX(t float64) {
	o.setX(t*o.vX() + o.posX())
}

func (o MvObj2) moveY(t float64) {
	o.setY(t*o.vY() + o.posY())
}

func (o MvObj2) Move(t float64) {
	o.moveX(t)
	o.moveY(t)
}
