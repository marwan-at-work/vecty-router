# Vecty Router

A declarative client-side router for [Vecty](https://www.github.com/gopherjs/vecty) applications. 
Similar to [react-router v4](https://github.com/ReactTraining/react-router)

### Installation

`go get github.com/marwan-at-work/vecty-router`

### Usage

You don't need to declare your routes at the top level. You can decalre them inside any component
and if they match they will render, otherwise, router will render an empty div instead. 

```go
package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/marwan-at-work/vecty-router"
)

// Body renders the <body>  tag
type Body struct {
	vecty.Core
}

// Render renders the <body> tag with the App as its children
func (b *Body) Render() *vecty.HTML {
	return elem.Body(
		router.NewRoute("/", &MainView{}),
		router.NewRoute("/blog", &Blog{}),
		router.NewRoute("/blog/{id}", &PostView{}),
	)
}
```

To retrieve a named variable like {id} in the example above you can do

```go
// Render returns every title
func (pv *PostView) Render() *vecty.HTML {
	id := router.GetNamedVar(pv)["id"]
	return elem.Div(
		vecty.Text(id),
	)
}
```

### Alternatives

- https://github.com/go-humble/router