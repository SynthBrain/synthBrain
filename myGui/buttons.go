package myGui

import (
	"fmt"
	"github.com/SynthBrain/synthBrain/vision"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/window"
)

func WebCam(posX, posY float32, onOff *bool) *gui.Button {
	button := *gui.NewButton("WebCam")
	button.SetPosition(posX, posY)
	button.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		// new gorutine for non-block app
		if *onOff == false {
			fmt.Println("start WebCam")
			vision.OnOff = false
			//go vision.StartWebCam(chImg)
			go vision.StartWebCam()
			*onOff = true
		} else {
			fmt.Println("stop WebCam")
			vision.OnOff = true
			*onOff = false
		}
	})
	return &button
}

func Exit(posX, posY float32, onOff *bool, win window.IWindow) *gui.Button {
	button := *gui.NewButton("Exit ")
	button.SetPosition(posX, posY)
	button.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("Application Close")
		defer closeWebCam(onOff)
		win.SetShouldClose(true)
	})
	return &button
}

func closeWebCam(onOff *bool) {
	if *onOff == true {
		fmt.Println("stop WebCam")
		vision.OnOff = true
		//onOff = false
	}
}
