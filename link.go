package router

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// EventCallback defines a vecty onClick handler
type EventCallback func(e *vecty.Event)

// LinkOptions - use to pass extra options to the link element
// like an ID, or class attribute.
type LinkOptions struct {
	ID    string
	Class string
}

// Link implements a frontend history Anchor tag.
func Link(route, text string, opts LinkOptions) *vecty.HTML {
	return elem.Anchor(
		vecty.Markup(
			vecty.MarkupIf(opts.ID != "", prop.ID(opts.ID)),
			vecty.MarkupIf(opts.Class != "", vecty.Class(opts.Class)),
			event.Click(onClick(route)).PreventDefault(),
		),
		vecty.Text(text),
	)
}

func onClick(route string) EventCallback {
	return func(e *vecty.Event) {
		js.Global.Get("history").Call(
			"pushState",
			map[string]string{"linkRoute": route},
			route,
			route,
		)
		refreshRoutes()
	}
}
