package snake

import (
	"fmt"
)

var (
	game *Game
)

type Game struct {
	boardRows int
	boardCols int

	snake  Snake
	score  int
	paused bool
}

func init() {
	game = initializeGameObjects(20, 20)
}

func initializeGameObjects(rows, cols int) *Game {
	return &Game{
		boardRows: rows,
		boardCols: cols,
		snake:     *NewSnake(point{1, 1}),
	}
}
func CurrGame() *Game {
	return game
}
func Head() *node {
	return game.snake.Head()
}
func Rows() int {
	return game.boardRows
}
func Cols() int {
	return game.boardCols
}

func updateGameState(g *Game) {
	if isGamePaused(g) {
		return
	}
	checkScore(g)
	updateSnake(g)
}

func updateSnake(g *Game) {

}

func checkScore(g *Game) {

}

func isGamePaused(g *Game) bool {
	return g.paused
}

func Paused() bool {
	game.paused = true
}

func HandleUserInput(key string) {
	x := Head().Point.X
	y := Head().Point.Y
	switch key {
	case "w":
		if y-1 < 0 {
			return
		}
		game.snake.body = &node{Point: point{x, y - 1}}
	case "s":
		if y+1 > Cols() {
			return
		}
		game.snake.body = &node{Point: point{x, y + 1}}
	case "a":
		if x-1 < Rows() {
			return
		}
		game.snake.body = &node{Point: point{x - 1, y}}
	case "d":
		if x+1 > Rows() {
			return
		}
		game.snake.body = &node{Point: point{x + 1, y}}
	case " ":
		fmt.Println("space pause")
	case "r":
		fmt.Println("r reset")

	}
}

func UpdateApple(g *Game) {
}

func isSnakeEatingItself() bool {
	return false
}

// function to check if snake ate an apple and return a boolean result
func isAppleInsideSnake() bool {
	return false

}

func updateScore() {
}

func GameState() string {
	return GameStr()

}

func GameStr() string {
	r := ""
	for y := range game.Height {
		r += "<div>"
		for x := range game.Width {
			// fmt.Println(Head(), y, x)
			if Head().Point.X == x && Head().Point.Y == y {
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
