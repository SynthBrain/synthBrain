package main

// 3840 * 2160 = 8 294 400
import (
	"github.com/g3n/engine/light"
	"SynthBrainGO/vision"
	"fmt"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
	"math/rand"
	"strconv"
)

func main() {
	//принимать в канал указатель *
	//chImg := make(chan *[640][480]byte)

	fmt.Println("Start NeuroMatrix")
	app, err := application.Create(application.Options{
		Title:     "NeuroMatrix",
		Width:     1280,
		Height:    600,
		TargetFPS: 60,
	})
	if err != nil {
		panic(err)
	}

	// add GUI
	// Create and add a label to the root panel
	//fps := float32(app.FrameCount()) / application.Get().RunSeconds()
	fps := 2
	l1 := gui.NewLabel("Simple GUI FPS: " + strconv.Itoa(int(fps)))
	l1.SetPosition(10, 10)
	l1.SetPaddings(2, 2, 2, 2)
	app.Gui().Root().Add(l1)

	// Create and add button 1 to the root panel
	onOff := false
	b1 := *gui.NewButton("WebCam")
	b1.SetPosition(10, 40)
	b1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		// new gorutine for non-block app
		if onOff == false {
			fmt.Println("start WebCam")
			vision.OnOff = false
			//go vision.StartWebCam(chImg)
			go vision.StartWebCam()

			onOff = true
		} else {
			fmt.Println("stop WebCam")
			vision.OnOff = true
			onOff = false
		}
	})
	app.Gui().Root().Add(b1)

	// Create and add exit button to the root panel
	b3 := *gui.NewButton("Exit ")
	b3.SetPosition(10, 70)
	b3.Subscribe(gui.OnClick, func(name string, ev interface{}) {
		fmt.Println("Application Close")
		//if onOff == true {
		//	fmt.Println("stop WebCam")
		//	vision.OnOff = true
		//	//onOff = false
		//}
		app.Window().SetShouldClose(true)
	})
	app.Gui().Root().Add(b3)

	//Create a blue torus and add it to the scene
	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	//mat := material.NewPhong(math32.NewColor("DarkBlue"))
	//torusMesh := graphic.NewMesh(geom, mat)
	//app.Scene().Add(torusMesh)

	myDots := 700
	for i := 0; i < myDots; i++ {
		go func() {
			dotGeom := geometry.NewCircle(0, 3)
			dotMat := material.NewPhong(math32.NewColor("DarkBlue"))
			dotMesh := graphic.NewPoints(dotGeom, dotMat)
			dotMesh.SetPosition(
				float32(rand.Int31n(15)),
				float32(rand.Int31n(15)),
				float32(rand.Int31n(15)))
			app.Scene().Add(dotMesh)
			//fmt.Println(dotMesh.Position())

			//if (i / 2) == 0 {
			//	dotsMesh.SetPosition(float32(rand.Int31n(20)),
			//float32(rand.Int31n(20)), float32(rand.Int31n(20)))
			//} else {
			//	dotsMesh.SetPosition(float32(rand.Int31n(20))+0.5,
			//float32(rand.Int31n(20)), float32(rand.Int31n(20))+0.5)
			//}

		}()
	}

	//Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
	pointLight.SetPosition(0, 0, 0)
	app.Scene().Add(pointLight)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(0.5)
	app.Scene().Add(axis)

	gridSizePosition := 10
	gridHelp := graphic.NewGridHelper(float32(gridSizePosition), 1, math32.NewColor("LightGrey"))
	gridHelp.SetPosition(float32(gridSizePosition/2), -0.2, float32(gridSizePosition/2))
	app.Scene().Add(gridHelp)

	app.CameraPersp().SetPosition(15, 15, 15)
	app.Gl().ClearColor(0, 0.5, 0.7, 1)
	//app.Gl().ClearColor(0.5,0.5,0.5,1)
	//app.Run()

	err = app.Run()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("app was running for %f \n", application.Get().RunSeconds())
}
