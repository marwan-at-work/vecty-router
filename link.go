package router

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
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
			prop.Href(route),
			vecty.MarkupIf(opts.ID != "", prop.ID(opts.ID)),
			vecty.MarkupIf(opts.Class != "", vecty.Class(opts.Class)),
			event.Click(onClick(route)).PreventDefault(),
		),
		vecty.Text(text),
	)
}

func onClick(route string) EventCallback {
	return func(e *vecty.Event) {
		Redirect(route)
	}
}
