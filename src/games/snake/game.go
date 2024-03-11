package snake

import (
	"fmt"
	"time"

	"wasm/internal/games"
	"wasm/internal/games/math"
)

const (
	EventBounderyCollision = iota
	EventSnakeCollision
	EventApple
)

type GameObject struct {
	boardRows int
	boardCols int
	StartTime time.Time

	snake *Snake
	apple Apple

	Score  int
	Paused bool
}

func InitializeGameObjects(rows, cols int) *GameObject {
	return &GameObject{
		boardRows: rows,
		boardCols: cols,
		StartTime: time.Now(),
		snake:     NewSnake(1, 1),
		apple: Apple{
			C:      games.NewCoordinate(rows/2, cols/2),
			Symbol: rune('A'),
		},
	}
}

func (g *GameObject) Snake() *Snake {
	return g.snake
}

func (g *GameObject) SetSnake(s *Snake) {
	g.snake = s
}

func (g *GameObject) Head() *games.Node {
	return g.snake.Head()
}

func (g *GameObject) Rows() int {
	return g.boardRows
}

func (g *GameObject) Cols() int {
	return g.boardCols
}

func HandleUserInput(g *GameObject, key string) {
	snake := g.Snake()
	x := snake.Head().X()
	y := snake.Head().Y()

	switch key {
	case "w", "ArrowUp":
		if y-1 < 0 {
			return
		}
		snake.SetVelocity(-1, 0)
	case "s", "ArrowDown":
		if y+1 > g.Cols() {
			return
		}
		// snake.UpdateSnake(NewNode(games.NewCoordinate(x, y+1)))
		snake.SetVelocity(1, 0)
	case "a", "ArrowLeft":
		if x-1 < 0 {
			return
		}
		// snake.SetHead(NewNode(games.NewCoordinate(x-1, y)))
		snake.SetVelocity(0, -1)
	case "d", "ArrowRight":
		if x+1 > g.Rows() {
			return
		}
		// snake.SetHead(NewNode(games.NewCoordinate(x+1, y)))
		snake.SetVelocity(0, 1)
	case " ":
		fmt.Println("space pause")
	case "r":
		fmt.Println("r reset")
	}
}

func (g *GameObject) Apple() *Apple {
	return &g.apple
}

func UpdateApple(g *GameObject) {
	var coord games.Coordinate
	for {
		rr := math.RandomInRange(0, g.Rows())
		rc := math.RandomInRange(0, g.Cols())

		coord = games.NewCoordinate(rr, rc)

		if validCoords(g, coord) {
			break
		}
	}
	g.apple = Apple{
		C:      coord,
		Symbol: rune('A'),
	}
}

func validCoords(g *GameObject, c games.Coordinate) bool {
	curr := g.Snake().Head()
	for curr != nil {
		if curr.Coord().Equal(c) {
			return false
		}
		curr = curr.Next
	}

	return true
}

// function to check if snake ate an apple and return a boolean result
func isAppleInsideSnake(s *Snake, a *Apple) bool {
	if s.Head().Coord().Equal(a.Coord()) {
		return true
	}
	return false
}

func isHeadInsideSnake(s *Snake) bool {
	if s.Head() == nil {
		panic("snake head nil")
	}
	headcoord := s.Head().Coord()

	if s.Head().Next != nil {
		curr := s.Head().Next
		for curr != nil {
			if curr.Coord().Equal(headcoord) {
				return true
			}
			curr = curr.Next
		}
	}

	return false
}

func CollisionCheck(s *Snake, a *Apple) int {
	// Check Apple
	if isAppleInsideSnake(s, a) {
		return EventApple
	}

	// Check Collision with snake
	if isHeadInsideSnake(s) {
		return EventSnakeCollision
	}

	return -1
}

func isPaused(g *GameObject) bool {
	return g.Paused
}
