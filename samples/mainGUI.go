// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package samples

import (
	"fmt"
	"runtime"

	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

func notGuimain() {

	// Creates window and OpenGL context
	wmgr, err := window.Manager("glfw")
	if err != nil {
		panic(err)
	}
	win, err := wmgr.CreateWindow(800, 600, "GUI Button", false)
	if err != nil {
		panic(err)
	}

	// OpenGL functions must be executed in the same thread where
	// the context was created (by window.New())
	runtime.LockOSThread()

	// Create OpenGL state
	gs, err := gls.New()
	if err != nil {
		panic(err)
	}

	// Creates GUI root panel and sets the background color
	root := gui.NewRoot(gs, win)
	root.SetColor(math32.NewColor("darkgray"))

	// Initial setting of the viewport and root panel size
	width, height := win.Size()
	gs.Viewport(0, 0, int32(width), int32(height))
	root.SetSize(float32(width), float32(height))

	//app, err := application.Create(application.Options{
	//	Title:  "NeuroMatrix",
	//	Width:  800,
	//	Height: 600,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	// Create a blue torus and add it to the scene
	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	//mat := material.NewPhong(math32.NewColor("DarkBlue"))
	//torusMesh := graphic.NewMesh(geom, mat)
	//app.Scene().Add(torusMesh)

	// Creates a camera
	aspect := float32(width) / float32(height)
	camera := camera.NewPerspective(65, aspect, 0.01, 1000)

	// Subscribe to window resize events. When the window is resized:
	// - Update the viewport size
	// - Update the root panel size
	// - Update the camera aspect ratio
	win.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) {
		width, height := win.Size()
		gs.Viewport(0, 0, int32(width), int32(height))
		root.SetSize(float32(width), float32(height))
		aspect := float32(width) / float32(height)
		camera.SetAspect(aspect)
	})

	// Create and add a label to the root panel
	l1 := gui.NewLabel("Simple GUI button demo")
	l1.SetPosition(10, 10)
	l1.SetPaddings(2, 2, 2, 2)
	root.Add(l1)

	// Create and add button 1 to the root panel
	b1 := gui.NewButton("button 1")
	b1.SetPosition(10, 40)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("button 1 clicked")
	})
	root.Add(b1)

	// Create and add button 2 to the root panel
	b2 := gui.NewButton("button 2")
	b2.SetPosition(100, 40)
	b2.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("button 2 clicked")
	})
	root.Add(b2)

	// Create and add exit button to the root panel
	b3 := gui.NewButton("exit ")
	b3.SetPosition(190, 40)
	b3.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		win.SetShouldClose(true)
	})
	root.Add(b3)

	// Creates a renderer and adds default shaders
	rend := renderer.NewRenderer(gs)
	err = rend.AddDefaultShaders()
	if err != nil {
		panic(err)
	}
	rend.SetGui(root)

	// Sets window background color
	gs.ClearColor(0.6, 0.6, 0.6, 1.0)

	// Render loop
	for !win.ShouldClose() {

		// Render the root GUI panel using the specified camera
		rendered, err := rend.Render(camera)
		if err != nil {
			panic(err)
		}
		wmgr.PollEvents()

		// Update window and checks for I/O events
		if rendered {
			win.SwapBuffers()
		}
	}
}
