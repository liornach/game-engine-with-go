package main

import (
	"example.com/go-engine/ggl"
)

// "log"
// "runtime"

// "github.com/go-gl/gl/v3.3-core/gl"
// "github.com/go-gl/glfw/v3.3/glfw"

func init() {
	// OpenGL requires locking to main thread
	//	runtime.LockOSThread()
}

func main() {
	/*
		// Initialize GLFW
		if err := glfw.Init(); err != nil {
			log.Fatalln("Failed to initialize GLFW:", err)
		}
		defer glfw.Terminate()

		// Set GLFW version hints for OpenGL 3.3 Core
		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.Resizable, glfw.False)

		// Create a window
		window, err := glfw.CreateWindow(800, 600, "OpenGL in Go", nil, nil)
		if err != nil {
			log.Fatalln("Failed to create window:", err)
		}
		window.MakeContextCurrent()

		// Initialize OpenGL bindings
		if err := gl.Init(); err != nil {
			log.Fatalln("Failed to initialize OpenGL bindings:", err)
		}

		version := gl.GoStr(gl.GetString(gl.VERSION))
		log.Println("OpenGL version:", version)

		// Main loop
		for !window.ShouldClose() {
			gl.ClearColor(0.2, 0.3, 0.3, 1.0)
			gl.Clear(gl.COLOR_BUFFER_BIT)

			window.SwapBuffers()
			glfw.PollEvents()
		}
	*/

	g := ggl.NewGwin(800, 500, "Lior Nachmias 2")
	defer g.Term()
}
