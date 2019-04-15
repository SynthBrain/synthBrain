package helpers

import (
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/util/application"
)

// AxisHelper add an axis helper
func AxisHelper(size float32, app *application.Application){
	axis := graphic.NewAxisHelper(size)
	app.Scene().Add(axis)
}

//GridHelper Add an grid helper to the scene
func GridHelper(size int, app *application.Application){
	gridHelp := graphic.NewGridHelper(float32(size), 1, math32.NewColor("LightGrey"))
	gridHelp.SetPosition(float32(size/2), -0.2, float32(size/2))
	app.Scene().Add(gridHelp)
}