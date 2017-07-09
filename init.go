package router

import (
	"github.com/gopherjs/gopherjs/js"
)

func init() {
	js.Global.Set("onpopstate", func(e *js.Object) {
		refreshRoutes()
	})
}
