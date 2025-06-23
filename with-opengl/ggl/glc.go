package ggl

import "github.com/go-gl/gl/v3.3-core/gl"

type opengl struct {
	win *gwin
}

func NewOpengl(width int, height int, title string) *opengl {
	g := &opengl{
		win: NewGwin(width, height, title),
	}

	g.win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		g.Term()
		panic(err)
	}

	return g
}

func Version() string {
	v := gl.GoStr(gl.GetString(gl.VERSION))
	return v
}

func (g *opengl) ViewPort(x, y, w, h int32) {
	gl.Viewport(x, y, w, h)
}

func (g *opengl) Term() {
	g.win.Term()
}
