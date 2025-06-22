package ggl

import (
	//"runtime"

	"fmt"
	"testing"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type gwin struct {
	win *glfw.Window
}

func initGlfw() {
	//runtime.LockOSThread()

	log("trying to initialize glfw")
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	log("successfully initialized glfw")

	vis := glfw.False
	var visstr string
	if vis == glfw.False {
		visstr = "false"
	} else {
		visstr = "true"
	}

	log("set window visibility hint to %s", visstr)
	glfw.WindowHint(glfw.Visible, vis)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	log("glfwgl package is initialized")
}

func log(s string, a ...any) {
	if !testing.Testing() {
		fmt.Printf(s, a...)
		fmt.Println()
	}
}

func NewGwin(w int, h int, title string) *gwin {
	initGlfw()

	g := &gwin{
		win: nil,
	}

	win, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		panic(err)
	}

	g.win = win

	return g
}

func (g *gwin) Width() int {
	w, _ := g.win.GetSize()
	return w
}

func (g *gwin) Height() int {
	_, h := g.win.GetSize()
	return h
}

func (g *gwin) Show() {
	g.win.Show()
}

func (g *gwin) Hide() {
	g.win.Hide()
}

func (g *gwin) IsVis() bool {
	vis := g.win.GetAttrib(glfw.Visible)
	switch vis {
	case glfw.False:
		return false
	case glfw.True:
		return true
	default:
		estr := fmt.Sprintf("unknown integer returned from GetAttrib (visibility) : %d", vis)
		panic(estr)
	}
}

func (g *gwin) Term() {
	glfw.Terminate()
}

func (g *gwin) MakeContextCurrent() {
	g.win.MakeContextCurrent()
}