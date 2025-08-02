package achtung

type gameIsRunningState struct {
}

func (r *gameIsRunningState) update(g *game) (bool, state) {
	g.log("updating running state")

	if len(g.collisions) > 0 {
		panic("game has an unhandled collisions")
	}

	elapsed := g.touchTimer()

	for _, curPlayer := range g.players {
		newHead := curPlayer.estimatePhysics(elapsed)
		nextWorldPos := newHead.toWorldPos()

		if existObjInWorld, ok := g.world[nextWorldPos]; ok {
			if existObjInWorld.isCollided(curPlayer, nextWorldPos) {
				col := collision{
					objectsInvolved: []objectInWorld{curPlayer, existObjInWorld},
					pos:             nextWorldPos,
				}
				g.collisions = append(g.collisions, col)
				g.logCollision(col)
				continue
			}
		} else { // this condition will meet only if player is not already own that position in world
			g.world[nextWorldPos] = curPlayer
			g.log("player %s was set in %v", curPlayer.uid(), nextWorldPos)
		}

		curPlayer.head = newHead

		prevVel := curPlayer.velocity
		if curPlayer.rotateIfKeysPressed(g.rotateSensitivity) {
			g.log("velocity of %s changed from %v to %v", curPlayer.uid, prevVel, curPlayer.velocity)
		}
	}

	if len(g.collisions) > 0 {
		return true, &collisionOccuredState{}
	}

	return false, r
}

type collisionOccuredState struct {
}

func (curState *collisionOccuredState) update(g *game) (bool, state) {
	panic("collision occured")
}

type initialState struct {
}

func (curState *initialState) update(g *game) (bool, state) {
	for _, p := range g.players {
		randPos := g.randomPos.next()
		p.head.x = float64(randPos.x)
		p.head.y = float64(randPos.y)
	}

	return true, &gameIsRunningState{}
}
