package snake

import "wasm/internal/games"

type Apple struct {
	C      games.Coordinate
	Symbol rune
}

func (a *Apple) Coord() games.Coordinate {
	return a.C
}

func (a *Apple) SetCoord(c games.Coordinate) {
	a.C = c
}
