package jsutil

import (
	"fmt"
	"syscall/js"
)

/**
 * grid id = gridcontainer
 * cells are cellx-y
 *
 * elementId component placement id ex <div id=snake></div>
 * cellCss css placed on each cell
 * gridCss css placed on grid container
 */
func CreateGrid(elementId, cellCss, gridCss string) {
	global := js.Global()
	document := global.Get("document")

	html := ""
	for y := range g.Height {
		for x := range g.Width {
			html += fmt.Sprintf(`<div><div id=%s class="%s"></div></div>`,
				Cellid(x, y), cellCss)
		}
	}
	element := document.Call("getElementById", elementId)
	newDiv := document.Call("createElement", "div")
	newDiv.Set("className", gridCss)
	newDiv.Set("id", "gridcontainer")
	newDiv.Set("innerHTML", html)
	element.Call("appendChild", newDiv)
}

func RootStyleSetProperty(args ...any) js.Value {
	return js.Global().Get("document").Get("documentElement").Get("style").Call("setProperty", args...)

}

func Cellid(x, y int) string {
	return fmt.Sprintf("cell%d-%d", x, y)
}
