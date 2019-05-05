package myGui

import (
	"fmt"
	"github.com/SynthBrain/synthBrain/vision"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/window"
	"time"
)

func WebCam(posX, posY float32, onOff *bool, chFlag chan bool) *gui.Button {
	button := *gui.NewButton("WebCam")
	button.SetPosition(posX, posY)
	button.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		// new gorutine for non-block app
		if *onOff == false {
			fmt.Println("start WebCam")
			go vision.StartWebCam(chFlag)
			*onOff = true
		} else {
			//fmt.Println("stop WebCam")
			closeWebCam(onOff, chFlag)
		}
	})
	return &button
}

func Exit(posX, posY float32, onOff *bool, win window.IWindow, chFlag chan bool) *gui.Button {
	button := *gui.NewButton("Exit ")
	button.SetPosition(posX, posY)
	button.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		if *onOff {
			closeWebCam(onOff, chFlag)
			fmt.Println("Application Close")
			time.Sleep(time.Second)
			win.SetShouldClose(true)
		} else {
			fmt.Println("Application Close")
			win.SetShouldClose(true)
		}
	})

	return &button
}

func closeWebCam(onOff *bool, chFlag chan bool) {
	chFlag <- true
	*onOff = false
}
