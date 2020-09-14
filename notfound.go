package router

import (
	"github.com/hexops/vecty"
)

var nf *notFound

// NotFoundHandler renders if no routes are matched at all.
func NotFoundHandler(c vecty.Component) vecty.Component {
	r := &notFound{c: c}
	nf = r
	return r
}

type notFound struct {
	vecty.Core
	c vecty.Component
}

func (r *notFound) Render() vecty.ComponentOrHTML {
	path := pathname()

	for _, r := range routes {
		if r.p.MatchString(path) {
			return nil
		}
	}

	return r.c
}
