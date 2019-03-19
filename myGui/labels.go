package myGui

import (
	"github.com/g3n/engine/gui"
)

func LabelFps(posX, posY float32, textLable string) *gui.Label{
	l1 := gui.NewLabel("GUI FPS: " + textLable)
	//l1 := gui.NewLabel("Simple GUI FPS: " + strconv.Itoa(int(fps)))
	l1.SetPosition(10, 10)
	l1.SetPaddings(2, 2, 2, 2)
	return l1
}