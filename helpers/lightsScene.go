package helpers

import (
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
	"github.com/g3n/engine/light"
)

// LightsScene add lights 
func LightsScene(app *application.Application){
	ambientLight := light.NewAmbient(math32.NewColor("White"), 0.8)//{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(math32.NewColor("White"), 5.0) //(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(0, 0, 0)
	app.Scene().Add(pointLight)
}