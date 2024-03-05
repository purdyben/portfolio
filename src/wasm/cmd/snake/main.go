package main

import (
	_ "embed"
	"fmt"
	"syscall/js"
	"wasm/snake"
)

func init() {
	js.Global().Set("updateRow", js.FuncOf(updateGameView))
}

// Declare a main function, this is the entrypoint into our go module
// That will be run. In our example, we won't need this
func main() {
	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("handleClick", js.FuncOf(handleClick))

	insertHtml("game", `
			<head>
			  <title>htmx Local Row Update Example</title>
			  <script src="https://unpkg.com/htmx.org@1.7.0/dist/htmx.min.js"></script>
			</head>
			<body>

			<table>
			  <thead>
			    <tr>
			      <th>ID</th>
			      <th>Name</th>
			      <th>Action</th>
			    </tr>
			  </thead>
			  <tbody>
			    <tr id="row1"  hx-trigger="click" hx-get="javascript: updateGameView()">
					before updating
			    </tr>
			    <!-- Other rows... -->
			  </tbody>
			</table>`)

	global := js.Global()
	document := global.Get("document")

	global.Get("console").Call("log", "User Input:", "33")
	ch := make(chan string, 10)
	document.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// fmt.Println("button clicked")
		event := args[0]
		ch <- event.Get("key").String()
		return nil
	}))

	// gameDiv := document.Call("getElementById", "game")
	// newDiv := document.Call("createElement", "div")
	// newDiv.Set("innerHTML", templates.Squarehtml)
	// gameDiv.Call("appendChild", newDiv)

	snake.CreateBoardHtml("GameContainer", snake.CurrGame())
	SetTemplateColumns()
	// fmt.Println("game h =", snake.CurrGame().Height)
	go func() {
		for {
			select {
			case key := <-ch:
				func() {
					defer recover()
					fmt.Println(key)
					snake.HandleUserInput(key)

					document := js.Global().Get("document")
					cell := document.Call("getElementById", fmt.Sprintf("cell%d-%d", snake.Head().Point.X, snake.Head().Point.Y))
					cell.Set("className", "grid-cell-snake")
				}()
			}
		}
	}()

	select {}

}

func SetTemplateColumns() {
	js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", "--board-rows", snake.CurrGame().Height)
	js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", "--board-cols", snake.CurrGame().Width)
	// fmt.Sprintf("--board-rows:%d;", snake.CurrGame().Height))
	// js.Global().Get("document").Get("documentElement").Set("style", fmt.Sprintf("--board-cols:%d;", snake.CurrGame().Width))
	// element := js.Global().Get("document").Call("getElementById", "gridcontainer")
	// element.Set("style", fmt.Sprintf("--board-rows:%d;", snake.CurrGame().Height))
	// element.Set("style", fmt.Sprintf("--board-cols:%d;", snake.CurrGame().Width))
}

func updateGameView(this js.Value, p []js.Value) interface{} {
	rowID := p[0].String()
	document := js.Global().Get("document")
	row := document.Call("getElementById", rowID)
	row.Call("querySelector", "td:nth-child(2)").Set("innerHTML", "<p>"+snake.GameState()+"</p>")
	return nil
}

//export add
func add(x int, y int) int {
	return x + y
}

func handleClick(this js.Value, inputs []js.Value) interface{} {
	println("Button clicked!")
	return nil
}
func insertHtml(divID, content string) {
	document := js.Global().Get("document")
	gameDiv := document.Call("getElementById", divID)
	newDiv := document.Call("createElement", "div")
	newDiv.Set("innerHTML", content)
	gameDiv.Call("appendChild", newDiv)
}

func inspectObject(obj js.Value) {
	// Example: Print all properties of the object
	fmt.Println("Object properties:")
	keys := js.Global().Get("Object").Call("keys", obj)
	for i := 0; i < keys.Length(); i++ {
		key := keys.Index(i).String()
		value := obj.Get(key)
		fmt.Printf("%s: %v\n", key, value)
		// Check if the property is an object and inspect it further
		if value.Type() == js.TypeObject {
			fmt.Printf("%s: (nested object)\n", key)
			inspectNestedObject(value)
		} else {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
}

func inspectNestedObject(obj js.Value) {
	// Example: Print all properties of the nested object
	keys := js.Global().Get("Object").Call("keys", obj)
	for i := 0; i < keys.Length(); i++ {
		key := keys.Index(i).String()
		value := obj.Get(key)
		fmt.Printf("  %s: %v\n", key, value)
	}
}
