package states

import "github.com/liornach/game-engine-ebiten/achtung"

type player = achtung.Player

type elimanted struct {
	p *player
}

func newEliminated(p *player) elimanted {
	return elimanted{
		p: p,
	}
}

func (s elimanted) Update(g *game) (bool, state) {
	panic("not implemented") // handle this state, not implemented at all
	// one player should be out, and if no players left, than one player won
	// currently there are only two players, so its important to check who's won
	// and to decide how much scores to give them for their win (for example, does a player gets
	// score only for eliminating another player, or only if he's the only one stayed in the game)
}
