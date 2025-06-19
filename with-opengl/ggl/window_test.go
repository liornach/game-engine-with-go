package ggl

import "testing"

var h = 400
var w = 300
var title = "test title"

func newTestGwin() *gwin {
	return NewGwin(w, h, title)
}

func TestGetWidth(t *testing.T) {
	g := newTestGwin()
	exp := w
	res := g.GetWidth()
	if res != exp {
		t.Errorf("expected : %d, result : %d", exp, res)
	}
}
