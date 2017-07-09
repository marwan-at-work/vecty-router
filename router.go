package router

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/vecty"
)

var routes = []*Route{}
var hasNamedVar = regexp.MustCompile("{[^/]+}")

// Route keeps track of the component and the route it's meant to be matched against.
// so when a the URL of the browser changes (pushState), we can check whether we need to render it or not.
type Route struct {
	vecty.Core
	pattern string
	c       vecty.Component
	e       *EmptyComponent
	p       *regexp.Regexp
}

// NewRoute takes a pattern and a component, if the pattern matches the
// current URL, then it returns the component, otherwise it returns
// an EmptyComponent.
func NewRoute(pattern string, c vecty.Component) *Route {
	r := &Route{
		pattern: pattern,
		c:       c,
		e:       &EmptyComponent{},
	}

	if hasNamedVar.MatchString(pattern) {
		r.p = regexp.MustCompile(
			fmt.Sprintf("^%v$", hasNamedVar.ReplaceAllString(pattern, "([^/]+)")),
		)
	}

	register(r)

	return r
}

// Render renders the underlying component or EmptyComponent if route does not match
func (r *Route) Render() *vecty.HTML {
	path := pathname()
	if matches(r.pattern, path) {
		return r.c.Render()
	}

	if r.p != nil && r.p.MatchString(path) {
		return r.c.Render()
	}

	return r.e.Render()
}

// GetNamedVar returns the parsed named variables from a URL.
// If you did NewRoute("/blog/{id}", someComponent{}).
// Then you can do GetNamedVar(someComponent{})
// and get a ma like this: {"id": "id-var"}
func GetNamedVar(c vecty.Component) map[string]string {
	vars := map[string]string{}

	var givenRoute *Route
	for _, r := range routes {
		if r.c == c {
			givenRoute = r
			break
		}
	}

	if givenRoute == nil || givenRoute.p == nil {
		return vars
	}

	// extract the named vars from url: "/users/{id}/{dog}" => ["{id}", "{dog}"]
	namedVars := hasNamedVar.FindAllString(givenRoute.pattern, -1)

	// remove the surrounding brackets from each named var
	for i := 0; i < len(namedVars); i++ {
		namedVars[i] = strings.Replace(
			strings.Replace(namedVars[i], "{", "", 1),
			"}",
			"",
			1,
		)
	}

	namedValues := givenRoute.p.FindAllStringSubmatch(pathname(), -1)[0][1:]

	for i := 0; i < len(namedVars); i++ {
		vars[namedVars[i]] = namedValues[i]
	}

	return vars
}

// matches checks if the given pattern is the same as the browser URL.
func matches(pattern, pathname string) bool {
	if pathname == pattern {
		return true
	}

	return false
}

// pathname gets the relative pathname from the browser.
func pathname() string {
	return js.Global.Get("location").Get("pathname").String()
}

func register(r *Route) {
	routes = append(routes, r)
}

func refreshRoutes() {
	for _, r := range routes {
		vecty.Rerender(r)
	}
}
