package pong

import "wasm/internal/games"

type Pong struct {
	Player1 Player
	Player2 Player
	Ball    Ball

	games.GameGrid
}
type Player struct {
	coord games.Coordinate
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
