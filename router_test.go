package router

import "testing"

func TestRouter(t *testing.T) {
	usersPath := "/users"

	tt := []struct {
		name     string
		route    *Route
		match    string
		expected bool
	}{
		{
			"exact match positiive",
			NewRoute(usersPath, nil, NewRouteOpts{ExactMatch: true}),
			"/users",
			true,
		},
		{
			"exact match negative",
			NewRoute(usersPath, nil, NewRouteOpts{ExactMatch: true}),
			"/users/1",
			false,
		},
		{
			"loose match exact",
			NewRoute(usersPath, nil, NewRouteOpts{ExactMatch: false}),
			"/users",
			true,
		},
		{
			"loose match loose",
			NewRoute(usersPath, nil, NewRouteOpts{ExactMatch: false}),
			"/users/1",
			true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.route.p.MatchString(tc.match) != tc.expected {
				t.Errorf("expected %v to have a %v match for %v", usersPath, tc.expected, tc.match)
			}
		})
	}
}
