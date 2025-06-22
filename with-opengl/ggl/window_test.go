package ggl

import (
	"fmt"
	"testing"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var h = 400
var w = 300
var title = "test title"

func newTestGwin() *gwin {
	return NewGwin(w, h, title)
}

func getTestWidth() int {
	return w
}

func getTestHeight() int {
	return h
}

func fmtExpRes(exp, res interface{}) string {
	return fmt.Sprintf("expected : %v, result : %v", exp, res)
}

func assert(exp, res interface{}, t *testing.T) {
	if exp != res {
		t.Error(fmtExpRes(exp, res))
	}
}

func TestGetWidth(t *testing.T) {
	g := newTestGwin()
	defer g.Term()
	exp := getTestWidth()
	res := g.Width()
	assert(exp, res, t)
}

func TestGetHeight(t *testing.T) {
	g := newTestGwin()
	defer g.Term()
	exp := getTestHeight()
	res := g.Height()
	assert(exp, res, t)
}

func TestShow(t *testing.T) {
	g := newTestGwin()
	defer g.Term()
	g.Show()
	exp := glfw.True
	res := g.win.GetAttrib(glfw.Visible)
	assert(exp, res, t)
}

func TestHide(t *testing.T) {
	g := newTestGwin()
	defer g.Term()
	g.Show()
	g.Hide()
	exp := glfw.False
	res := g.win.GetAttrib(glfw.Visible)
	assert(exp, res, t)
}

func TestIsVis(t *testing.T) {
	g := newTestGwin()
	defer g.Term()
	g.Hide()
	exp := false
	res := g.IsVis()
	assert(exp, res, t)

	g.Show()
	exp = true
	res = g.IsVis()
	assert(exp, res, t)
}
