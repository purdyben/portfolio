package main

import (
	"context"
	_ "embed"
	"fmt"
	"sync"
	"syscall/js"
	"time"

	"wasm/cmd/snake/util"
	"wasm/internal/games"
	"wasm/internal/games/math"
	"wasm/internal/jsutil"
	"wasm/snake"
)

var (
	userInput chan string
	global    = js.Global()
	document  = global.Get("document")

	boardRows = 20
	boardCols = 20

	gameContext, gameCancel = context.WithCancel(context.Background())
	ticker                  *games.GameTicker
	_game                   *snake.GameObject
	mu                      sync.Mutex
)

func init() {
	// js.Global().Set("updateRow", js.FuncOf(updateGameView))

	ticker = games.NewTicker(90 * time.Millisecond)

	fmt.Println("Setting Controls")
	js.Global().Set("snakeStop", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			ticker.Stop()
			jsutil.Console("stop")
			return nil
		}))
	js.Global().Set("snakeStart", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			jsutil.Console("start")
			ticker.Start()
			return "hellp"
		}))
	js.Global().Set("snakeRestart", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			exitGame()
			startGame()
			snake.UpdateApple(Game())
			jsutil.Console("restart")
			snake.CreateBoardHtml("GameContainer", Game().Rows(), Game().Cols())
			util.Render(Game())
			return ""
		}))

	js.Global().Set("updateApple", js.FuncOf(
		func(this js.Value, args []js.Value) interface{} {
			snake.UpdateApple(Game())
			return ""
		}))

	jsutil.RootStyleSetProperty("--board-rows", boardRows)
	jsutil.RootStyleSetProperty("--board-cols", boardCols)
}

func Game() *snake.GameObject {
	mu.Lock()
	defer mu.Unlock()
	return _game
}

func SetGame(g *snake.GameObject) {
	mu.Lock()
	defer mu.Unlock()
	_game = g
}

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	insertHtml("game", `<div>Snake</div>
						<button onclick="snakeStart()">Start</button>
						<button onclick="snakeStop()">Stop</button>
						<button onclick="snakeRestart()">Restart</button>

						<button onclick="updateApple()">updateApple</button>`)

	document.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		userInput <- event.Get("key").String()
		return nil
	}))

	startGame()
	select {}
}

func startGame() {
	fmt.Println("Starting New Game")

	gameContext, gameCancel = context.WithCancel(context.Background())

	userInput = make(chan string, 2)

	// starting snake
	SetGame(snake.InitializeGameObjects(boardRows, boardCols))

	s := snake.NewSnake(5, 5)

	Game().SetSnake(s)
	snake.UpdateApple(Game())

	go gameTicker(gameContext, ticker)
	go handelUserInput(gameContext, userInput)

	snake.CreateBoardHtml("GameContainer", Game().Rows(), Game().Cols())
	util.Render(Game())
}

func exitGame() {
	if gameCancel != nil {
		gameCancel()
	}
}

func gameTicker(ctx context.Context, ticker *games.GameTicker) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.Tick():
			game := Game()
			// Update Snake

			lastCoord := games.Last(game.Snake().Head()).Coord()

			event := game.Snake().UpdateBasedOnVelocity(game.Rows(), game.Cols())
			if event != -1 {
				fmt.Println("Update Event ", event)
			}
			event = snake.CollisionCheck(game.Snake(), game.Apple())

			if event == snake.EventApple {
				Game().Score += 1
				snake.UpdateApple(game)
				Game().Snake().AddNode(lastCoord.X(), lastCoord.Y())
			}

			if event != -1 {
				fmt.Println("Update Event ", event)
			}
			// Render Updates
			util.Render(game)
			math.RandomInRange(0, int(math.Seed()))
		}
	}
}

func handelUserInput(ctx context.Context, userInput chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case key := <-userInput:
			func() {
				defer recover()
				fmt.Println(key)
				snake.HandleUserInput(Game(), key)
			}()
		}
	}
}

func ResetGame() {
	snake.CreateBoardHtml("GameContainer", Game().Rows(), Game().Cols())
}

func insertHtml(divID, content string) {
	document := js.Global().Get("document")
	gameDiv := document.Call("getElementById", divID)
	newDiv := document.Call("createElement", "div")
	newDiv.Set("innerHTML", content)
	gameDiv.Call("appendChild", newDiv)
}

func SetTemplateColumns() {
	js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", "--board-rows", Game().Rows())
	js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", "--board-cols", Game().Cols())
	// fmt.Sprintf("--board-rows:%d;", snake.Currme().Height))
	// js.Global().Get("document").Get("documenvvctElement").Set("style", fmt.Sprintf("--board-cols:%d;", snake.CurrGame().Width))
	// element := js.Global().Get("document").Call("getElementById", "gridcontainer")
	// element.Set("style", fmt.Sprintf("--board-rows:%d;", snake.CurrGame().Height))
	// element.Set("style", fmt.Sprintf("--board-cols:%d;", snake.CurrGame().Width))
}
