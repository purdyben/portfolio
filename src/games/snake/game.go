package snake

import (
	"fmt"
	"time"

	"wasm/internal/games"
)

var game *GameObject

var seed int64 // Seed for the LCG

func init() {
	// Initialize the seed (you can use any initial value)
	seed = 42
	game = initializeGameObjects(20, 20)
}

type GameObject struct {
	boardRows int
	boardCols int
	startTime time.Time

	snake  *Snake
	Apple  Apple
	score  int
	paused bool
}

type Apple struct {
	Coord  games.Coordinate
	Symbol rune
}

func initializeGameObjects(rows, cols int) *GameObject {
	return &GameObject{
		boardRows: rows,
		boardCols: cols,
		startTime: time.Now(),
		snake:     NewSnake(1, 1),
		Apple: Apple{
			Coord:  games.NewCoordinate(rows/2, cols/2),
			Symbol: rune('A'),
		},
	}
}

func Game() *GameObject {
	return game
}

func (g *GameObject) Snake() *Snake {
	return g.snake
}

func (g *GameObject) SetSnake(s *Snake) {
	game.snake = s
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

func Paused() bool {
	return game.paused
}

func HandleUserInput(key string) {
	x := game.snake.Head().X()
	y := game.snake.Head().Y()
	snake := game.snake
	switch key {
	case "w":
		if y-1 < 0 {
			return
		}
		// snake.UpdateSnake(NewNode(games.NewCoordinate(x, y-1)))
		snake.SetVelocity(-1, 0)
		// fmt.Println("setting", -1, 0)
		// a, b := snake.Velocity()
		// fmt.Println("got", a, b)
	case "s":
		if y+1 > game.Cols() {
			return
		}
		// snake.UpdateSnake(NewNode(games.NewCoordinate(x, y+1)))
		snake.SetVelocity(1, 0)
	case "a":
		if x-1 < 0 {
			return
		}
		// snake.SetHead(NewNode(games.NewCoordinate(x-1, y)))
		snake.SetVelocity(0, -1)
	case "d":
		if x+1 > game.Rows() {
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

func lcg() int {
	// LCG parameters (use appropriate values for your needs)
	a, c, m := int64(1664525), int64(1013904223), int64(1<<31)

	// Linear Congruential Generator formula
	seed = (a*seed + c) % m

	// Convert the result to an int
	return int(seed)
}

func randomInRange(max, min int) int {
	// Ensure the range is valid
	if min >= max {
		panic("Invalid range")
	}

	// Use lcg() to generate a random number within the range [min, max)
	return min + lcg()%(max-min)
}

func UpdateApple(g *GameObject) {
	rr := randomInRange(g.Rows(), 0)
	rc := randomInRange(g.Cols(), 0)
	g.Apple = Apple{
		Coord:  games.NewCoordinate(rr, rc),
		Symbol: rune('A'),
	}
}

func isSnakeEatingItself() bool {
	return false
}

// function to check if snake ate an apple and return a boolean result
func isAppleInsideSnake() bool {
	return false
}

func GameState() string {
	return GameStr()
}

func GameStr() string {
	r := ""
	for y := range game.Rows() {
		r += "<div>"
		for x := range game.Cols() {
			// fmt.Println(Head(), y, x)
			if game.Head().X() == x && game.Head().Y() == y {
				r += "[S]"
			} else {
				r += "[ ]"
			}
		}
		r += "</div>\n"
		// b = append(b, make([]int, game.width))
	}
	return r
}

func isPaused(g *GameObject) bool {
	return g.paused
}

func Pause() {
	game.paused = !game.paused
}
