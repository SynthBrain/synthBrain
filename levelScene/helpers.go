package levelScene

import (
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/util/application"
)


func AxisHelper(size float32, app *application.Application){
	axis := graphic.NewAxisHelper(size)
	app.Scene().Add(axis)
}

func GridHelper(size int, app *application.Application){
	// Add an grid helper to the scene
	gridHelp := graphic.NewGridHelper(float32(size), 1, math32.NewColor("LightGrey"))
	gridHelp.SetPosition(float32(size/2), -0.2, float32(size/2))
	app.Scene().Add(gridHelp)
}