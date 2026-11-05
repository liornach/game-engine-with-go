package achtung

type Color string

type Coord struct {
	X, Y int
}

type floatVector struct {
	x, y float64
}
type player struct {
	Color Color
	direction floatVector 
	head floatVector 
	unitPerTick floatVector 
}

type Achtung struct {
	board map[Coord]Color	
	players []player
}

func (a Achtung) Board() [Coord]Color {
	return a.board
}

func (a *Achtung) ApplyTicks(t uint) {
	for i range a.players {
		x := t * a.players[i].unitPerTick.x
		y := t * a.players[i].unitPerTock.y
	}
