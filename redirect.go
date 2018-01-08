package router

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

// Redirect ...
func Redirect(route string) EventCallback {
	return func(e *vecty.Event) {
		js.Global.Get("history").Call(
			"pushState",
			map[string]string{"redirectRoute": route},
			route,
			route,
		)
		refreshRoutes()
	}
}
