package samples

import (
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/util/application"
	"math/rand"
)
// 3840 * 2160 = 8 294 400
func notmain() {

	app, err := application.Create(application.Options{
		Title:  "NeuroMatrix",
		Width:  800,
		Height: 600,
	})
	if err != nil {
		panic(err)
	}

	// Create a blue torus and add it to the scene
	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	//mat := material.NewPhong(math32.NewColor("DarkBlue"))
	//torusMesh := graphic.NewMesh(geom, mat)
	//app.Scene().Add(torusMesh)

	// Try Create Dots
	geomDots := geometry.NewCircle(0.02, 3)
	matDots := material.NewPhong(math32.NewColor("DarkBlue"))
	dotsMesh := graphic.NewMesh(geomDots, matDots)

	dotsMesh.SetPosition(0, 0, 0)
	app.Scene().Add(dotsMesh)

	myDots := 5000
	for i := 0; i < myDots; i++ {
		go func() {
			geomDots := geometry.NewCircle(0.01, 3)
			matDots := material.NewPhong(math32.NewColor("DarkBlue"))
			dotsMesh := graphic.NewMesh(geomDots, matDots)

			dotsMesh.SetPosition(rand.Float32(), rand.Float32(), rand.Float32())
			app.Scene().Add(dotsMesh)
		}()
	}



	// Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{1.0, 1.0, 1.0}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	app.Scene().Add(pointLight)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(0.5)
	app.Scene().Add(axis)

	app.CameraPersp().SetPosition(0, 0, 3)
	//app.Run()

	err = app.Run()
	if err != nil {
		panic(err)
	}
	//fmt.Printf("app was running for %f \n", application.Get().RunSeconds())
}
