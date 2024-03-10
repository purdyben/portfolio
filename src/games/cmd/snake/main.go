package main

import (
	_ "embed"
	"fmt"
	"syscall/js"
	"time"

	"wasm/cmd/snake/util"
	"wasm/internal/games"
	"wasm/internal/jsutil"
	"wasm/snake"
)

var (
	userInput chan string
	global    = js.Global()
	document  = global.Get("document")

	ticker = games.NewTicker(80 * time.Millisecond)
)

func init() {
	userInput = make(chan string, 10)
	// js.Global().Set("updateRow", js.FuncOf(updateGameView))
	ResetGame()
	jsutil.RootStyleSetProperty("--board-rows", snake.Game().Rows())
	jsutil.RootStyleSetProperty("--board-cols", snake.Game().Cols())

	fmt.Println("Setting Controls")
	js.Global().Set("snakeStop", util.Stop(ticker))
	js.Global().Set("snakeStart", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			jsutil.Console("start")
			ticker.Start()
			return "hellp"
		}))
	js.Global().Set("snakeRestart", util.Restart(ticker))
}

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	insertHtml("game", `<div>Snake</div>
						<button onclick="snakeStart()">Start</button>
						<button onclick="snakeStop()">Stop</button>
						<button onclick="snakeRestart()">Restart</button>`)

	document.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		userInput <- event.Get("key").String()
		return nil
	}))
	ticker.Ticker.Stop()
	go handelUserInput(userInput)
	go rungame(ticker)
	select {}
}

func rungame(ticker *games.GameTicker) {
	game := snake.Game()
	s := snake.NewSnake(5, 5)

	n := games.NewNode(5, 6)
	s.Head().Next = n

	n.Next = games.NewNode(5, 7)
	game.SetSnake(s)

	for {
		select {
		case <-ticker.Tick():
			// Update Snake
			game.Snake().UpdateBasedOnVelocity()
			//

			// Render Updates
			util.Render(game)
			a, b := game.Snake().Velocity()
			fmt.Printf("Velocity %d %d\n", a, b)
		}
	}
}

func ResetGame() {
	snake.CreateBoardHtml("GameContainer", snake.Game().Rows(), snake.Game().Cols())
}

func handelUserInput(userInput chan string) {
	// game := snake.Game()
	for {
		select {
		case key := <-userInput:
			func() {
				defer recover()
				fmt.Println(key)
				snake.HandleUserInput(key)

				// document := js.Global().Get("document")
				// cell := document.Call("getElementById", fmt.Sprintf("cell%d-%d", game.Head().X(), game.Head().Y()))
				// cell.Set("className", "grid-cell-snake")
				// Render(game.Head().X(), game.Head().Y())
			}()
		}
	}
}

func insertHtml(divID, content string) {
	document := js.Global().Get("document")
	gameDiv := document.Call("getElementById", divID)
	newDiv := document.Call("createElement", "div")
	newDiv.Set("innerHTML", content)
	gameDiv.Call("appendChild", newDiv)
}

func SetTemplateColumns() {
	js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", "--board-rows", snake.Game().Rows())
	js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", "--board-cols", snake.Game().Cols())
	// fmt.Sprintf("--board-rows:%d;", snake.Currme().Height))
	// js.Global().Get("document").Get("documenvvctElement").Set("style", fmt.Sprintf("--board-cols:%d;", snake.CurrGame().Width))
	// element := js.Global().Get("document").Call("getElementById", "gridcontainer")
	// element.Set("style", fmt.Sprintf("--board-rows:%d;", snake.CurrGame().Height))
	// element.Set("style", fmt.Sprintf("--board-cols:%d;", snake.CurrGame().Width))
}
