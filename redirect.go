package router

import (
	"github.com/gopherjs/gopherjs/js"
)

// Redirect ...
func Redirect(route string) {
	js.Global.Get("history").Call(
		"pushState",
		map[string]string{"redirectRoute": route},
		route,
		route,
	)
	refreshRoutes()
}
