package myGui

import (
	"github.com/g3n/engine/gui"
)

//var l2 *gui.Label

func LabelFps(posX, posY float32, textLable string) *gui.Label{
	l1 := gui.NewLabel("GUI FPS: " + textLable)
	//l1 := gui.NewLabel("Simple GUI FPS: " + strconv.Itoa(int(fps)))
	l1.SetPosition(posX, posY)
	l1.SetPaddings(2, 2, 2, 2)
	return l1
}

// func LabelFpsTest(posX, posY float32, textLable string, app *application.Application){
// 	l2 = gui.NewLabel("GUI FPS: " + textLable)
// 	//l1 := gui.NewLabel("Simple GUI FPS: " + strconv.Itoa(int(fps)))
// 	l2.SetPosition(posX, posY)
// 	l2.SetPaddings(2, 2, 2, 2)
// 	app.Gui().Root().Add(l2)
// }