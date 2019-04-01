package main


// // Create a tic-tac-toe board.
// board := [][]string{
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// }

// for i := 0; i < len(board); i++ {
// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
// }

// 3840 * 2160 = 8 294 400
import (
	"fmt"
	"math/rand"
	"synthBrain/levelScene"
	"synthBrain/myGui"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)
/*
	Рисовать только тех что имеют достаточный уровень активность и окончательно не затухли
*/
func main() {
	//принимать в канал указатель *
	//chImg := make(chan *[640][480]byte)

	fmt.Println("Start NeuroMatrix")
	app, err := application.Create(application.Options{
		Title:     "NeuroMatrix",
		Width:     1280,
		Height:    600,
		TargetFPS: 140,
	})
	if err != nil {
		panic(err)
	}

	// add GUI*********************************************************
	// Create and add a label to the root panel
	l1 := myGui.LabelFps(10, 10, "240")
	app.Gui().Root().Add(l1)

	// go func() {
	// 	for {
	// 		if a, b, c := app.FrameRater().FPS(60); a > 0 && b > 0 && c == true {
	// 			fmt.Println("FPS ", int(b))
	// 		}
	// 	}
	// }()

	//fps := float32(app.FrameCount()) / application.Get().RunSeconds()

	//go myGui.LabelFpsTest(10, 10, strconv.Itoa(int(app.FrameCount()) / int(application.Get().RunSeconds())), app)

	// Create and add button 1 to the root panel
	onOff := false
	b1 := myGui.WebCam(10, 40, &onOff, app)
	app.Gui().Root().Add(b1)

	// Create and add exit button to the root panel
	b2 := myGui.Exit(10, 70, &onOff, app)
	app.Gui().Root().Add(b2)
	//******************************************************************

	//Create a blue torus and add it to the scene
	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	//mat := material.NewPhong(math32.NewColor("DarkBlue"))
	//torusMesh := graphic.NewMesh(geom, mat)
	//app.Scene().Add(torusMesh)

	myDots := 7000
	for i := 0; i < myDots; i++ {
		go func() {
			// dotGeom := geometry.NewCircle(0.2, 3)
			// dotMat := material.NewPhong(math32.NewColor("White"))
			// dotMesh := graphic.NewMesh(dotGeom, dotMat)
			// dotMesh.SetPosition(
			// 	float32(rand.Int31n(15)),
			// 	float32(rand.Int31n(15)),
			// 	float32(rand.Int31n(15)))
			// app.Scene().Add(dotMesh)

			dotGeom := geometry.NewCircle(0, 3)
			//dotGeom := geometry.NewGeometry()
			dotMat := material.NewPhong(math32.NewColor("White"))
			dotMesh := graphic.NewPoints(dotGeom, dotMat)
			dotMesh.SetPosition(
				float32(rand.Int31n(20)),
				float32(rand.Int31n(20)),
				float32(rand.Int31n(20)))
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
	levelScene.LightsScene(app)

	// Add an axis helper to the scene
	levelScene.AxisHelper(0.5, app)

	// Add an grid helper to the scene
	levelScene.GridHelper(10, app)

	// Add camera to the scene
	app.CameraPersp().SetPosition(15, 15, 15)
	//app.Gl().ClearColor(0, 0.5, 0.7, 1)
	app.Gl().ClearColor(0, 0.2, 0.4, 1)

	// Start application
	err = app.Run()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("app was running for %f \n", application.Get().RunSeconds())
}
