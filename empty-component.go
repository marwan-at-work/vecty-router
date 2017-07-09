package router

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// EmptyComponent is rendered instead of a given component when
// the route doesn't match.
type EmptyComponent struct {
	vecty.Core
}

// Render renders an empty <div>
func (e *EmptyComponent) Render() *vecty.HTML {
	return elem.Div(
		prop.ID("vecty-router-empty-component"),
	)
}
