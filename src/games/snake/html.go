package snake

import (
	"fmt"
	"syscall/js"
)

func CreateBoardHtml(id string, rows, cols int) {
	fmt.Println("Creating Board", rows, cols)
	global := js.Global()
	document := global.Get("document")

	html := ""
	for y := range rows {
		// html += "<div id=row" + strconv.Itoa(y) + ">"
		for x := range cols {
			html += `<div> 
				<div id=` + fmt.Sprintf("cell%d-%d", x, y) + ` class="grid-cell"></div>
			</div>`
		}
		// html += "</div>"
	}
	element := document.Call("getElementById", id)
	element.Set("innerHTML", "")

	newDiv := document.Call("createElement", "div")
	newDiv.Set("className", "grid-container")
	newDiv.Set("id", "gridcontainer")
	newDiv.Set("innerHTML", html)
	element.Call("appendChild", newDiv)

	// gridcontainer := document.Call("getElementById", "gridcontainer")
	// gridcontainer.Set("innerHTML", "")
	// gridcontainer.Set("innerHTML", html)
}
