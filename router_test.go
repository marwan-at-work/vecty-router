package router

import "testing"

func TestSimpleMatch(t *testing.T) {
	if !matches("pattern/one", "pattern/one") {
		t.Error("expected pattern/one to match itself")
	}

	if matches("pattern/one", "pattern/two") {
		t.Error("expected pattern/one to not match pattern/two")
	}
}
