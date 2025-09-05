package achtung

type CollisionType int

const (
	BothBlame CollisionType = iota
	Player1Blame
	Player2Blame
)

struct collision {
	player1, player2 *Player
	collisionType CollisionType
}

	
