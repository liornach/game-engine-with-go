package twodim

type Pos struct {
	X, Y float64
}

func (p *Pos) IsLessThanOrEqualTo(o *Pos) bool {
	return p.X <= o.X && p.Y <= o.Y
}

func (p *Pos) IsGreaterThanOrEqualTo(o *Pos) bool {
	return p.X >= o.X && p.Y >= o.Y
}
