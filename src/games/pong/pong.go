package pong

import "wasm/internal/games"

type Pong struct {
	Player1 Player
	Player2 Player
	Ball    Ball

	games.GameGrid

	player1Score int
	player2Score int
}
type Player struct {
	coord      games.Coordinate
	matchscore int
}
type Ball struct {
	coord    games.Coordinate
	velocity float32
}

func NewPlayer(x, y int) Player {
	return Player{
		coord: games.NewCoordinate(x, y),
	}
}

func NewBall() Ball {
	return Ball{
		coord:    games.NewCoordinate(0, 0),
		velocity: 0,
	}
}

func NewPong() Pong {
	p := Pong{
		Player1: NewPlayer(0, 0),
		Player2: NewPlayer(0, 0),
		Ball:    NewBall(),
	}
	p.Rows = 10
	p.Cols = 10
	return p
}

func Score(p *Pong, player int) {
	if player == 1 {
		p.player1Score += 1
	} else if player == 2 {
		p.player2Score += 1
	}

	// If match is over
	if p.Player1.matchscore == 3 {
		// end match
		p.Player1.matchscore += 1
	} else if p.Player2.matchscore == 3 {
		// end match
		p.Player2.matchscore += 1
	}
	// Reset Ball
}
