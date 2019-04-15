package neurons

import (
	"github.com/g3n/engine/util/application"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/graphic"
)

// Synapse - lines
type Synapse struct{
	graphic.Lines
}

// NewSynapse - create and return lines with new parameters
func NewSynapse(app *application.Application, start math32.Vector3, stop math32.Vector3, color *math32.Color) *Synapse{
	syn := new(Synapse)

	geom := geometry.NewGeometry()
	vertices := math32.NewArrayF32(0, 6)
	vertices.Append(
			start.X, start.Y, start.Z,
			stop.X, stop.Y, stop.Z,
	)
		
	colors := math32.NewArrayF32(0, 6)
	colors.Append(
		color.R, color.G, color.B,
		color.R, color.G, color.B,
	)
	geom.AddVBO(gls.NewVBO(vertices).AddAttrib(gls.VertexPosition))
	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))

	// Creates basic material
	mat := material.NewBasic()

	// Creates lines with the specified geometry and material
	syn.Lines.Init(geom, mat)
	app.Scene().Add(syn)
	return syn
}