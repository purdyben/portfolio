package util

import (
	"fmt"
	"syscall/js"

	"wasm/snake"
)

func Render(g *snake.GameObject) {
	for row := range g.Rows() {
		for col := range g.Cols() {
			SetCellGrass(row, col)
		}
	}

	curr := g.Snake().Head()
	for curr != nil {
		SetCellSnake(curr.X(), curr.Y())
		curr = curr.Next
	}
	ac := g.Apple().Coord()
	SetCellApple(ac.X(), ac.Y())
}

func SetCellSnake(x, y int) {
	setCellCss("grid-cell-snake", x, y)
}

func SetCellApple(x, y int) {
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
