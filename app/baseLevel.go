package app

import (
	"math/rand"
	"time"

	"github.com/SynthBrain/synthBrain/appGui"
	"github.com/SynthBrain/synthBrain/vision"
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
	// GUI
	Restart  *gui.Button
	Exit     *gui.Button
	WebCam   *gui.Button
	// mesh *graphic.Points
	// positions math32.ArrayF32
}

// Start is called once at the start of the demo.
func (level *Level) Start(app *App) {
	// Create and add a button to the scene
	onOff := false
	chOnOffFlag := make(chan bool, 1)
	level.Exit =  appGui.Exit(10, 40, &onOff, app.Application, chOnOffFlag)
	level.WebCam =  appGui.WebCam(10, 10, &onOff, chOnOffFlag)
	app.DemoPanel().Add(level.WebCam)
	app.DemoPanel().Add(level.Exit)

	// Create axes helper
	axes := helper.NewAxes(2)
	app.Scene().Add(axes)

	// Creates geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	colors := math32.NewArrayF32(0, 16)

	numPoints := 100000
	coord := float32(10)
	for i := 0; i < numPoints; i++ {
		var vertex math32.Vector3
		vertex.Set(
			rand.Float32()*coord, //-coord/2,
			rand.Float32()*coord, //-coord/2,
			rand.Float32()*coord, //-coord/2,
		)
		positions.AppendVector3(&vertex)
		colors.Append(rand.Float32(), rand.Float32(), rand.Float32())
	}

	// for i := 0; i < 100000; i++{
	// 	positions.Append(float32(rand.Int31n(50)), float32(rand.Int31n(50)), float32(rand.Int31n(50)))
	// 	colors.Append(rand.Float32(), rand.Float32(), rand.Float32())
	// 	//colors.Append(1, 0, 0)
	// }

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
	mesh.SetScale(1, 1, 1)
	//a.Scene().Add(t.mesh)
	app.Scene().AddAt(0, mesh)
}

// Update is called every frame.
func (level *Level) Update(app *App, deltaTime time.Duration) {
	//vision.ReadImg(app.dirData, "/0.jpg")


	//app.Scene().RemoveAt(0)
	//level.Start(app)
}

// Cleanup is called once at the end of the demo.
func (level *Level) Cleanup(app *App) {}
