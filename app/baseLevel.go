package app

import (
	_ "fmt"
	"github.com/g3n/engine/texture"
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
	meshTest *graphic.Points
	generalSphere *graphic.Mesh

	plane *graphic.Mesh
	// positions math32.ArrayF32
	angl float32

	flag bool
}

// Start is called once at the start of the demo.
func (level *Level) Start(app *App) {
	// Init Base Logic
	level.logic = baseLogic.InitLogic()

	// Create and add a button to the scene
	level.flag = true
	level.onOff = false
	chOnOffFlag := make(chan bool, 1)
	level.WebCam = appGui.WebCam(10, 10, &level.onOff, chOnOffFlag, level.logic.VisionChan)
	app.DemoPanel().Add(level.WebCam)

	level.Exit = appGui.Exit(10, 40, &level.onOff, app.Application, chOnOffFlag)
	app.DemoPanel().Add(level.Exit)

	// Create axes helper
	axes := helper.NewAxes(30)
	app.Scene().Add(axes)

	// Creates generalSphere
	level.makeSphere(app, 1000, *math32.NewVector3(0, 0,0))
	// start all sensor part from this place
	level.makeSphere(app, 100, *math32.NewVector3(0, -500,0))

	// vision
	level.makeSphere(app, 500, *math32.NewVector3(500, 0,0))
	// motion system for Vision
	level.makeSphere(app, 250, *math32.NewVector3(500, 625,0))

	// sound
	level.makeSphere(app, 500, *math32.NewVector3(-500, 0,0))

	// motivation
	level.makeSphere(app, 50, *math32.NewVector3(0, -750,0))

	level.initInputLayerVision(app)
	// Create plane with VisionData**********************************
	//level.initInputDataPlane(app)
	//*****************************************************************
}

// Update is called every frame.
func (level *Level) Update(app *App, deltaTime time.Duration) {
	level.logic.Update()
	if(level.logic.GetReady()){
		//level.Dispose(app)
		level.DrawInputLayerVision(app)
	}
}

func (level *Level) Dispose(app *App) {
	//fmt.Printf("Length: ",  len(app.Scene().Children()))
	for i := len(app.Scene().Children()); i > 8; i-- {
		app.Scene().ChildAt(i-1).Dispose()
		app.Scene().RemoveAt(i-1)
	}
}

// Cleanup is called once at the end of the demo.
func (level *Level) Cleanup(app *App) {
	app.Scene().RemoveAll(true)
}

func (level *Level) DrawInputLayerVision(app *App) {
	temp := level.mesh.GetGeometry().VBOs()
	positions := math32.NewArrayF32(0, 0)
	colors := math32.NewArrayF32(0, 16)
	var vertex math32.Vector3
	for i := 0; i < len(level.logic.DataVision); i++ {
		for j := 0; j < len(level.logic.DataVision[0]); j++ {
			color := level.logic.DataVision[i][j]
			vertex.Set(
				float32(j),
				0,
				float32(i),
			)
			positions.AppendVector3(&vertex)
			colors.Append(float32(color), float32(color), float32(color))
		}
	}
	if(level.flag){
		temp[0].SetBuffer(positions)
		level.flag = false
	}
	//temp[0].SetBuffer(positions)
	temp[1].SetBuffer(colors)
}

func (level *Level) initInputLayerVision(app *App){
	geom := geometry.NewGeometry()
	colors := math32.NewArrayF32(0, 16)
	color := 1
	colors.Append(float32(color), float32(color), float32(color))
	positions := math32.NewArrayF32(0, 0)
	positions.Append(
		0, 0, 0,
	)
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
	//level.mesh.SetScale(1, 1, 1)
	app.Scene().Add(level.mesh)
}

func (level *Level) makeSphere(app *App, size float64, pos math32.Vector3){
	geom := geometry.NewSphere(size, 32, 16)
	mat := material.NewStandard(&math32.Color{1, 1, 1})
	mat.SetWireframe(true)
	mat.SetSide(material.SideDouble)
	level.generalSphere = graphic.NewMesh(geom, mat)
	level.generalSphere.SetPosition(pos.X, pos.Y, pos.Z)
	app.Scene().Add(level.generalSphere)
}

func (level *Level) initGeom(app *App) {
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)

	var vertex math32.Vector3
	for i := 0; i < len(level.logic.DataVision); i++ {
		for j := 0; j < len(level.logic.DataVision[0]); j++ {
			vertex.Set(
				float32(j),
				0,
				float32(i),
			)
			positions.AppendVector3(&vertex)
		}
	}

	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	level.mesh = graphic.NewPoints(geom, nil)
	//app.Scene().Add(level.mesh)
}

func (level *Level) initMaterial(app *App) {
	tex, err := texture.NewTexture2DFromImage(app.DirData() + "/images/" + "snowflake1.png")
	if err != nil {
		app.Log().Fatal("Error loading texture: %s", err)
	}

	mat := material.NewPoint(&math32.Color{1, 1, 1})
	mat.SetTransparent(true)
	mat.SetOpacity(0.6)
	mat.AddTexture(tex)
	mat.SetSize(600)
	mat.SetBlending(material.BlendAdditive)
	mat.SetDepthMask(false)
	level.mesh.AddMaterial(level.mesh, mat, 0, len(level.logic.DataVision) * len(level.logic.DataVision[0]))
}


//func (level *Level) inputLayerVision(index float32, app *App) {
//	// Creates geometry
//	geom := geometry.NewGeometry()
//	positions := math32.NewArrayF32(0, 0)
//	colors := math32.NewArrayF32(0, 16)
//
//	var vertex math32.Vector3
//	for i := 0; i < len(level.logic.DataVision); i++ {
//		for j := 0; j < len(level.logic.DataVision[0]); j++ {
//			color := level.logic.DataVision[i][j]
//			vertex.Set(
//				float32(j),
//				index,
//				float32(i),
//			)
//			positions.AppendVector3(&vertex)
//			colors.Append(float32(color), float32(color), float32(color))
//		}
//	}
//
//	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
//	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))
//	positions = nil // Positions cannot be used after transfering to VBO
//	colors = nil
//
//	// Creates point material
//	//mat := material.NewPoint(&math32.Color{0, 0, 0})
//	mat := material.NewBasic()
//	//mat.SetSize(50)
//
//	// Creates points mesh
//	level.mesh = graphic.NewPoints(geom, mat)
//	app.Scene().Add(level.mesh)
//}

func (level *Level) updPlaneMaterial(app *App) {
	texfile := app.DirData() + "\\webCam.jpg"
	tex, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		return
	} else {
		plane_mat := material.NewStandard(&math32.Color{1, 1, 1})
		plane_mat.SetSide(material.SideDouble)
		plane_mat.AddTexture(tex)
		level.plane.SetMaterial(plane_mat)
	}
}

func (level *Level) initInputDataPlane(app *App) {
	// Create plane with VisionData**********************************
	texfile := app.DirData() + "\\webCam.jpg"//"/images/tiger1.jpg"
	tex, err := texture.NewTexture2DFromImage(texfile)
	if err != nil {
		//app.Log().Fatal("Error:%s loading texture:%s", err, texfile)
		return
	} else {
		// Creates plane2
		plane_geom := geometry.NewPlane(640, 480)
		plane_mat := material.NewStandard(&math32.Color{1, 1, 1})
		plane_mat.SetSide(material.SideDouble)
		plane_mat.AddTexture(tex)
		level.plane = graphic.NewMesh(plane_geom, plane_mat)
		level.plane.SetPosition(320, -3, 240)
		level.angl = -1.57
		level.plane.RotateX(level.angl)
		app.Scene().Add(level.plane)
	}
}