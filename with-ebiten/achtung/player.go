
type Color = int
const (
	Red Color = 0XFF0000
	Green Color = 0X00FF00
	Blue Color = 0X0000FF
)

type Player struct {
	Color Color
	Left, Right ebiten.Key
	x, y float64
	vx, vy float64
}

func (p *Player) UpdateLocation(ticks int) {
	p.x += vx * ticks
	p.y += vy * ticks
}
