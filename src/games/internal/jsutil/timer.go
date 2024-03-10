package jsutil

import (
	"syscall/js"
	"time"
)

type Watch struct {
	Start   time.Time
	Enabled bool
}

// func NewTimer() *Watch {
// 	c := make(chan struct{}, 0)

// 	println("WASM Go Initialized")

// 	w := &Watch{}

// 	return w
// }

func SetInterval(fn func(this js.Value, args []js.Value) interface{}, delay int32) interface{} {
	return js.Global().Call("setInterval", js.FuncOf(fn), delay)
}

func ClearInterval(i interface{}) js.Value {
	return js.Global().Call("clearInterval", i)
}

// func (w *Watch) update(this js.Value, args []js.Value) interface{} {
// 	if !w.Enabled {
// 		return nil
// 	}

// 	since := time.Since(w.Start)
// 	minutes := int(since.Minutes())
// 	// outStr := fmt.Sprintf(
// 	// 	"%02d:%02d",
// 	// 	minutes,
// 	// 	int(since.Seconds())-(minutes*60),
// 	// )
// 	// js.Global().Get("document").Call("getElementById", "timer").Set("innerHTML", outStr)

// 	return nil
// }

// func (w *Watch) start(this js.Value, args []js.Value) interface{} {
// 	if w.Enabled {
// 		return nil
// 	}

// 	w.Start = time.Now()
// 	w.Enabled = true

// 	return nil
// }

// func (w *Watch) stop(this js.Value, args []js.Value) interface{} {
// 	if !w.Enabled {
// 		return nil
// 	}

// 	w.Enabled = false

// 	return nil
// }

// func main() {
// 	c := make(chan struct{}, 0)

// 	println("WASM Go Initialized")

// 	w := &Watch{}

// 	js.Global().Set("start", js.FuncOf(w.start))
// 	js.Global().Set("stop", js.FuncOf(w.stop))

// 	// Auto call update

// 	<-c
// }
