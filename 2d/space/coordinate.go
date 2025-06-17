package space

type Coordinate struct {
	X, Y, Z int
}

func NewCoordinate(x, y, z int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
		Z: z,
	}
}

func