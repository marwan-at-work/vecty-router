package router

import (
	"syscall/js"
)

// Redirect ...
func Redirect(route string) {
	js.Global().Get("history").Call(
		"pushState",
		map[string]interface{}{"redirectRoute": route},
		route,
		route,
	)
	refreshRoutes()
}
