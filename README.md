# Vecty Router

A declarative client-side router for [Vecty](https://www.github.com/hexops/vecty) applications. 
Similar to [react-router v4](https://github.com/ReactTraining/react-router)

### Installation

`go get marwan.io/vecty-router`

### Usage

You don't need to declare your routes at the top level. You can declare them inside any component
and if they match they will render, otherwise, router will render an empty div instead. 

```go
package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"marwan.io/vecty-router"
)

// Body renders the <body> tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		router.NewRoute("/", &MainView{}, router.NewRouteOpts{ExactMatch: true}),
		router.NewRoute("/blog", &Blog{}, router.NewRouteOpts{}),
		router.NewRoute("/blog/{id}", &PostView{}, router.NewRouteOpts{ExactMatch: true}),
	)
}
```

To retrieve a named variable like {id} in the example above you can do

```go
// Render returns every title
func (pv *PostView) Render() vecty.ComponentOrHTML {
	id := router.GetNamedVar(pv)["id"]
	return elem.Div(
		vecty.Text(id),
	)
}
```

### Other features

##### Navigation through links

```go
func (c *component) Render() vecty.ComponentOrHTML {
	return elem.Span(
		router.Link("/my/route", "click here", router.LinkOptions{}),
	)
}
```

##### Programatically navigate to a route

```go
router.Redirect("/my/route")
```

### Status

Currently vecty-router does not fallback to hash routing if the History API is not on your browser.
It also calls vecty.Rerender on all routes whenever a route changes. It should/will do its own deducing of whether to call rerender on a route or not based on route matches and whether it's already mounted or not.

### Alternatives

- https://github.com/go-humble/router
