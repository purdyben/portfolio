package main

import (
	"fmt"
	"syscall/js"

	"wasm/internal/jsutil"
)

var global = js.Global()

func main() {
	fmt.Println("Hello from pong")
	jsutil.Console("hello console from pong")
	select {}
}

func SetCellSnake(x, y int) {
	setCellCss("grid-cell-snake", x, y)
}

func SetCellApple(x, y int) {
	fmt.Println(x, y)
	setCellCss("grid-cell-apple", x, y)
}

func SetCellGrass(x, y int) {
	setCellCss("grid-cell", x, y)
}

func setCellCss(css string, x, y int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover()", err)
		}
	}()
	cell := js.Global().Get("document").Call("getElementById", fmt.Sprintf("cell%d-%d", x, y))

	if cell.IsUndefined() {
		fmt.Println("Log", "Error", "Setting", css, x, y)
	} else {
		cell.Set("className", css)
	}
}

func Render(x, y int) {
	// Update Player1
	// Update Player2
	// Update Ball
	// Update Game

	// for row := range snake.Game().Rows() {
	// 	for col := range snake.Game().Cols() {
	// 		SetCellGrass(row, col)
	// 	}
	// }
	// SetCellSnake(x, y)
	// snake.UpdateApple(snake.Game())
	// SetCellApple(snake.Game().Apple.Coord.X(), snake.Game().Apple.Coord.Y())
}
