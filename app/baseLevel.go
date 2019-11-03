package app

import (
	_ "fmt"
	_"math/rand"
	"time"

	"github.com/SynthBrain/synthBrain/appGui"
	"github.com/SynthBrain/synthBrain/baseLogic"
	_"github.com/SynthBrain/synthBrain/vision"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
)

func init() {
	DemoMap["BaseLevel"] = &Level{}
}

type Level struct {
	// BaseLogic
	logic *baseLogic.Logic

	onOff bool
	// GUI
	Restart *gui.Button
	Exit    *gui.Button
	WebCam  *gui.Button
	// mesh *graphic.Points
	// positions math32.ArrayF32
}

// Start is called once at the start of the demo.
func (level *Level) Start(app *App) {
	// Init Base Logic
	level.logic = baseLogic.InitLogic()

	// Create and add a button to the scene
	level.onOff = false
	chOnOffFlag := make(chan bool, 1)
	level.WebCam = appGui.WebCam(10, 10, &level.onOff, chOnOffFlag, level.logic.VisionChan)
	app.DemoPanel().Add(level.WebCam)

	level.Exit = appGui.Exit(10, 40, &level.onOff, app.Application, chOnOffFlag)
	app.DemoPanel().Add(level.Exit)

	// Create axes helper
	axes := helper.NewAxes(2)
	app.Scene().Add(axes)

	// //Creates geometry
	// geom := geometry.NewGeometry()
	// positions := math32.NewArrayF32(0, 0)
	// colors := math32.NewArrayF32(0, 16)

	// numPoints := 300000
	// coord := float32(10)
	// for i := 0; i < numPoints; i++ {
	// 	var vertex math32.Vector3
	// 	vertex.Set(
	// 		rand.Float32()*coord, //-coord/2,
	// 		rand.Float32()*coord, //-coord/2,
	// 		rand.Float32()*coord, //-coord/2,
	// 	)
	// 	positions.AppendVector3(&vertex)
	// 	colors.Append(rand.Float32(), rand.Float32(), rand.Float32())
	// }

	//// for i := 0; i < 100000; i++{
	//// 	positions.Append(float32(rand.Int31n(50)), float32(rand.Int31n(50)), float32(rand.Int31n(50)))
	//// 	colors.Append(rand.Float32(), rand.Float32(), rand.Float32())
	//// 	//colors.Append(1, 0, 0)
	//// }

	// geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	// geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))
	// positions = nil // Positions cannot be used after transfering to VBO
	// colors = nil

	// // Creates point material
	// //mat := material.NewPoint(&math32.Color{0, 0, 0})
	// mat := material.NewBasic()
	// //mat.SetSize(50)

	// // Creates points mesh
	// mesh := graphic.NewPoints(geom, mat)
	// mesh.SetScale(1, 1, 1)
	// //a.Scene().Add(t.mesh)
	// app.Scene().AddAt(1, mesh)
}

// Update is called every frame.
func (level *Level) Update(app *App, deltaTime time.Duration) {
	// vision.ReadImg(app.dirData, "/0.jpg")

	// update baseLogic.upd()
	level.logic.Update()
	app.Scene().RemoveAll(true)
	//****************************DEMO*************************************
	// get data from baseLogic and use data for update 3D objects on scene
	level.make3DLayer(0, 0, *level.logic.GetData(), app)
}

// Cleanup is called once at the end of the demo.
func (level *Level) Cleanup(app *App) {
	app.Scene().RemoveAll(true)
}

func (level *Level) make3DLayer(index int, size int, coords map[math32.Vector3]float32, app *App) {
	// Creates geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	colors := math32.NewArrayF32(0, 16)

	//numPoints := size
	//coord := float32(10)
	// for i := 0; i < numPoints; i++ {
	// 	var vertex math32.Vector3
	// 	vertex.Set(
	// 		coords[i].Y,
	// 		coords[i].X,
	// 		coords[i].Z,
	// 	)
	// 	coords[i] = coords / 255
	// 	positions.AppendVector3(&vertex)
	// 	//colors.Append(rand.Float32(), rand.Float32(), rand.Float32())
	// 	colors.Append(coords, coords, coords)
	// }

	// ПЛОХОЙ ВАРИАНТ РИСОВАНИЯ СЦЕНЫ
	for i := 0; i < 480; i++ {
		for j := 0; j < 640; j++ {
			temp := coords[*math32.NewVector3(float32(i), float32(j), 0)]
			var vertex math32.Vector3
			vertex.Set(
				float32(j),
				0,
				float32(i),
			)
			//temp = temp / 100
			positions.AppendVector3(&vertex)
			colors.Append(float32(temp), float32(temp), float32(temp))
		}
	}

	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))
	positions = nil // Positions cannot be used after transfering to VBO
	colors = nil

	// Creates point material
	//mat := material.NewPoint(&math32.Color{0, 0, 0})
	mat := material.NewBasic()
	//mat.SetSize(50)

	// Creates points mesh
	mesh := graphic.NewPoints(geom, mat)
	//mesh.SetScale(1, 1, 1)
	//app.Scene().Add(mesh)
	app.Scene().AddAt(index, mesh)
}
