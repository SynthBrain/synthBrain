package myGui

import (
	"SynthBrainGO/vision"
	"fmt"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/util/application"
)

func WebCam(posX, posY float32, onOff *bool) *gui.Button{
	b1 := *gui.NewButton("WebCam")
	b1.SetPosition(10, 40)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
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
	return &b1
}

func Exit(posX, posY float32, onOff *bool, app *application.Application) *gui.Button{
	button := *gui.NewButton("Exit ")
	button.SetPosition(10, 70)
	button.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("Application Close")
		defer closeWebCam(onOff)
		app.Window().SetShouldClose(true)
	})
	return &button
}

func closeWebCam(onOff *bool){
	if *onOff == true {
		fmt.Println("stop WebCam")
		vision.OnOff = true
		//onOff = false
	}
}