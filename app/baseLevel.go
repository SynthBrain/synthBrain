package app

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	_"github.com/g3n/engine/util/helper"
	"time"
	"math/rand"
)

func init() {
	DemoMap["Baselevel"] = &Level{}
}

type Level struct{
	// mesh *graphic.Points
	// positions math32.ArrayF32
}

// Start is called once at the start of the demo.
func (level *Level) Start(a *App) {

	// Create axes helper
	//axes := helper.NewAxes(2)
	//a.Scene().Add(axes)

	// Creates geometry
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	colors := math32.NewArrayF32(0, 16)
	
	for i := 0; i < 100000; i++{
		positions.Append(float32(rand.Int31n(50)), float32(rand.Int31n(50)), float32(rand.Int31n(50)))
		colors.Append(rand.Float32(), rand.Float32(), rand.Float32())
		//colors.Append(1, 0, 0)
	}

	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))
	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))
	positions = nil // Positions cannot be used after transfering to VBO
	colors = nil

	// Creates point material
	//mat := material.NewPoint(&math32.Color{0.3, 0.3, 0.3})
	mat := material.NewBasic()
	//mat.SetSize(50)

	// Creates points mesh
	mesh := graphic.NewPoints(geom, mat)
	mesh.SetScale(1, 1, 1)
	//a.Scene().Add(t.mesh)
	a.Scene().AddAt(0, mesh)
}

// Update is called every frame.
func (level *Level) Update(a *App, deltaTime time.Duration) {
	a.Scene().RemoveAt(0)
	level.Start(a)
}

// Cleanup is called once at the end of the demo.
func (level *Level) Cleanup(a *App) {}
