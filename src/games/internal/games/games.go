package games

type GameGrid struct {
	Rows int
	Cols int
}

func NewGameGrid(r, c int) GameGrid {
	return GameGrid{r, c}
}

type Coordinate struct {
	x, y int
}

func (g *GameGrid) IN(x, y int) bool {
	if x > g.Cols || x < 0 || y < 0 || y > g.Rows {
		return false
	}
	return true
}

func (c1 Coordinate) Equal(c2 Coordinate) bool {
	return c1.x == c2.x && c1.y == c2.y
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{x, y}
}

func (c *Coordinate) X() int {
	return c.x
}

func (c *Coordinate) Y() int {
	return c.y
}
