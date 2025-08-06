package physics

type TwoDVector struct {
	X, Y float64
}

type Location TwoDVector
type Velocity TwoDVector

type StateVector struct {
	location Location
	velocity Velocity
}

func NewStateVector(l Location, v Velocity) StateVector {
	return StateVector{
		location: l,
		velocity: v,
	}
}

func (sv StateVector) Location() Location {
	return sv.location
}

func (sv StateVector) Velocity() Velocity {
	return sv.velocity
}
