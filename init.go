package router

import (
	"syscall/js"
)

func init() {
	js.Global().Set("onpopstate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		refreshRoutes()
		return nil
	}))
}
