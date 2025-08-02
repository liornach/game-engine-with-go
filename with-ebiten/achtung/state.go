package achtung

import (
	"fmt"
)

type gameState int

const (
	gameIsRunning gameState = iota
	collisionOccured
)

type state interface {
	update(g *game) (bool, gameState)
}

type gameIsRunningState struct {
}

func (r *gameIsRunningState) update(g *game) (bool, gameState) {
	g.log("entering update loop")
	elapsed := g.touchTimer()
	colls := 0

	for _, curPlayer := range g.players {
		newHead := curPlayer.estimatePhysics(elapsed)
		nextWorldPos := newHead.toWorldPos()

		if existObjInWorld, ok := g.world[nextWorldPos]; ok {
			if existObjInWorld.isCollided(curPlayer, nextWorldPos) {
				colls++
				g.logCollision(existObjInWorld, curPlayer, nextWorldPos)
				continue
			}
		} else { // this condition will meet only if player is not already own that position in world
			g.world[nextWorldPos] = curPlayer
			g.log("player %s was set in %v", curPlayer.uid, nextWorldPos)
		}

		curPlayer.head = newHead

		prevVel := curPlayer.velocity
		if curPlayer.rotateIfKeysPressed(g.rotateSensitivity) {
			g.log("velocity of %s changed from %v to %v", curPlayer.uid, prevVel, curPlayer.velocity)
		}
	}

	if colls != 0 {
		fmt.Scanln()
		return true, collisionOccured
	}

	g.log("leaving update loop")
	return false, gameIsRunning
}
