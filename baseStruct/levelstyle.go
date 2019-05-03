package baseStruct

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
)

// LevelStyle - level styling information and functions
type LevelStyle struct {
	activeOn  *math32.Color
	activeOff *math32.Color

	materialWhite *material.Phong
	materialBlack *material.Phong

	materialSynapse *material.Basic

	MakeWhiteNeuron func() *graphic.Points
	MakeBlackNeuron func() *graphic.Points
}

// NewBaseStyle
func NewBaseStyle(dataDir string) *LevelStyle {

	s := new(LevelStyle)

	s.activeOn = &math32.Color{1, 1, 1}  // white
	s.activeOff = &math32.Color{0, 0, 0} // black

	// Helper function to load texture and handle errors
	//newTexture := func(path string) *texture.Texture2D {
	//	tex, err := texture.NewTexture2DFromImage(path)
	//	if err != nil {
	//		log.Fatal("Error loading texture: %s", err)
	//	}
	//	return tex
	//}

	// Load textures and create materials
	s.materialWhite = material.NewPhong(math32.NewColor("white"))
	//s.materialWhite.AddTexture(newTexture(dataDir + "/assets/white.png"))

	s.materialBlack = material.NewPhong(math32.NewColor("black"))
	//s.materialBlack.AddTexture(newTexture(dataDir + "/assets/black.png"))

	s.materialSynapse = material.NewBasic()

	// Create functions that return a cube mesh using the provided material, reusing the same cube geometry

	//sharedCubeGeom := geometry.NewCube(0.1)

	//*****************************Neuron*******************************
	sharedCircleGeom := geometry.NewCircle(0, 3)
	makeObjWithMaterial := func(mat *material.Phong) func() *graphic.Points {
		return func() *graphic.Points { return graphic.NewPoints(sharedCircleGeom, mat) }
	}

	s.MakeWhiteNeuron = makeObjWithMaterial(s.materialWhite)
	s.MakeBlackNeuron = makeObjWithMaterial(s.materialBlack)

	return s
}

func (s *LevelStyle) MakeSynapseLine(start math32.Vector3, stop math32.Vector3, color *math32.Color) func() *graphic.Lines {

	makeSynapseWithMaterial := func(mat *material.Basic) func() *graphic.Lines {
		return func() *graphic.Lines {
			return graphic.NewLines(s.synapseBody(geometry.NewGeometry(), start, stop, color), mat)
		}
	}
	return makeSynapseWithMaterial(s.materialSynapse)
}

func (s *LevelStyle) synapseBody(geom *geometry.Geometry,
	start math32.Vector3,
	stop math32.Vector3,
	color *math32.Color) *geometry.Geometry {

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

	return geom
}
