package app

import (
	_ "fmt"
	_ "math/rand"
	"time"

	"github.com/SynthBrain/synthBrain/appGui"
	"github.com/SynthBrain/synthBrain/baseLogic"
	_ "github.com/SynthBrain/synthBrain/vision"
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
	mesh *graphic.Points
	// positions math32.ArrayF32
	angl float32
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
}

// Update is called every frame.
func (level *Level) Update(app *App, deltaTime time.Duration) {
	level.logic.Update()
	if(level.logic.GetReady()){
		app.Scene().ChildAt(0).Dispose()
		app.Scene().RemoveAll(true)
		level.make3DLayer(0, app)
	}
}

func (level *Level) Dispose() {
	level.mesh.Dispose()
}

// Cleanup is called once at the end of the demo.
func (level *Level) Cleanup(app *App) {
	app.Scene().RemoveAll(true)
}

func (level *Level) make3DLayer(index float32, app *App) {
	// Creates geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	colors := math32.NewArrayF32(0, 16)

	var vertex math32.Vector3
	for i := 0; i < len(level.logic.Data); i++ {
		for j := 0; j < len(level.logic.Data[0]); j++ {
			color := level.logic.Data[i][j]
			vertex.Set(
				float32(j),
				index,
				float32(i),
			)
			positions.AppendVector3(&vertex)
			colors.Append(float32(color), float32(color), float32(color))
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
	level.mesh = graphic.NewPoints(geom, mat)
	app.Scene().Add(level.mesh)
}
