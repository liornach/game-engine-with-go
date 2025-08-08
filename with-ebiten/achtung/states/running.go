package states

type runningState struct {
	game *game
}

//type objectInWorld = achtung.ObjectInWorld

func (rs runningState) Update(g *game) (bool, state) {
	rs.game = g
	if g.CollisionsCount() > 0 {
		panic("game has an unhandled collisions")
	}

	timer := g.Timer()
	elapsed := timer.Touch()

	for _, p := range g.Players {
		nextWorldPos := g.EstimatedNextWorldPos(p, elapsed)

		if g.IsPosFree(nextWorldPos) {
			g.SetPosOwner(nextWorldPos, p)
		} else if g.IsPlayerHeadAt(p, nextWorldPos) {

		} else {
			owner, _ := g.PosOwner(nextWorldPos)
			col := collision{
				Objects: []objectInWorld{p, owner},
				Pos:     nextWorldPos,
			}
			g.AddCollision(col)
			continue
		}

		g.ApplyPhysicsToPlayer(p, elapsed)
		g.HandlePlayerKeys(p)
	}

	if g.CollisionsCount() > 0 {
		nextState := newHandlingCollisions(g.Collisions())
		return true, &nextState
	}

	return false, rs
}
