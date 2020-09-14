package router

import (
	"regexp"
	"strings"
	"syscall/js"

	fmt "github.com/cathalgarvey/fmtless"
	"github.com/hexops/vecty"
)

var routes = []*Route{}
var hasNamedVar = regexp.MustCompile("{[^/]+}")

// Route keeps track of the component and the route it's meant to be matched against.
// so when a the URL of the browser changes (pushState), we can check whether we need to render it or not.
type Route struct {
	vecty.Core
	pattern string
	c       vecty.Component
	p       *regexp.Regexp
}

// NewRouteOpts let you specify route matching options, such as exact match.
type NewRouteOpts struct {
	ExactMatch bool
}

// NewRoute takes a pattern and a component, if the pattern matches the
// current URL, then it returns the component, otherwise it returns
// an EmptyComponent.
func NewRoute(pattern string, c vecty.Component, opts NewRouteOpts) *Route {
	r := &Route{
		pattern: pattern,
		c:       c,
	}

	routePattern := pattern

	if hasNamedVar.MatchString(pattern) {
		routePattern = hasNamedVar.ReplaceAllString(pattern, "([^/]+)")
	}

	if opts.ExactMatch {
		routePattern = fmt.Sprintf("^%v$", routePattern)
	}

	r.p = regexp.MustCompile(routePattern)

	register(r)

	return r
}

// Render renders the underlying component or EmptyComponent if route does not match
func (r *Route) Render() vecty.ComponentOrHTML {
	path := pathname()

	if r.p.MatchString(path) {
		return r.c
	}

	return nil
}

// GetNamedVar returns the parsed named variables from a URL.
// If you did NewRoute("/blog/{id}", someComponent{}).
// Then you can do GetNamedVar(someComponent{})
// and get a map like this: {"id": "id-var"}
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

// pathname gets the relative pathname from the browser.
func pathname() string {
	return js.Global().Get("location").Get("pathname").String()
}

func register(r *Route) {
	routes = append(routes, r)
}

func refreshRoutes() {
	for _, r := range routes {
		vecty.Rerender(r)
	}
	if nf != nil {
		vecty.Rerender(nf)
	}
}
