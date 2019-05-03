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

	materialWhite    *material.Phong
	materialBlack    *material.Phong
	materialDotRed   *material.Phong
	materialDotWhite *material.Phong

	materialSynapse *material.Basic

	MakeWhiteCube func() *graphic.Mesh
	MakeBlackCube func() *graphic.Mesh

	MakeWhiteDot      func() *graphic.Points
	MakeSuperWhiteDot func() *graphic.Points
	MakeRedDot        func() *graphic.Points
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

	//**********************************************************************************
	s.materialSynapse = material.NewBasic()
	s.materialDotRed = material.NewPhong(math32.NewColor("Red"))
	s.materialDotWhite = material.NewPhong(math32.NewColor("White"))

	// Create functions that return a cube mesh using the provided material, reusing the same cube geometry
	//*****************************Cube*******************************
	sharedCubeGeom := geometry.NewCube(0.5)
	makeObjWithMaterial := func(mat *material.Phong) func() *graphic.Mesh {
		return func() *graphic.Mesh { return graphic.NewMesh(sharedCubeGeom, mat) }
	}

	s.MakeWhiteCube = makeObjWithMaterial(s.materialWhite)
	s.MakeBlackCube = makeObjWithMaterial(s.materialBlack)

	//*****************************Dots*******************************
	sharedDotGeom := s.neuronGeom()
	makeDotWithMaterial := func(mat *material.Phong) func() *graphic.Points {
		return func() *graphic.Points { return graphic.NewPoints(sharedDotGeom, mat) }
	}
	s.MakeWhiteDot = makeDotWithMaterial(s.materialDotWhite)
	s.MakeRedDot = makeDotWithMaterial(s.materialDotRed)

	//****************************SuperWhiteDots**************************
	sharedDotGeomSuperWhite := s.neuronGeom()
	makeDotWithMaterialWhite := func(mat *material.Basic) func() *graphic.Points {
		return func() *graphic.Points { return graphic.NewPoints(sharedDotGeomSuperWhite, mat) }
	}
	s.MakeSuperWhiteDot = makeDotWithMaterialWhite(s.materialSynapse)

	return s
}

func (s *LevelStyle) MakeSynapseLine(start math32.Vector3, stop math32.Vector3, color *math32.Color) func() *graphic.Lines {

	makeSynapseWithMaterial := func(mat *material.Basic) func() *graphic.Lines {
		return func() *graphic.Lines {
			return graphic.NewLines(s.synapseGeom(geometry.NewGeometry(), start, stop, color), mat)
		}
	}
	return makeSynapseWithMaterial(s.materialSynapse)
}

func (s *LevelStyle) synapseGeom(geom *geometry.Geometry,
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

func (s *LevelStyle) neuronGeom() *geometry.Geometry {

	geom := geometry.NewGeometry()
	vertices := math32.NewArrayF32(0, 3)
	vertices.Append(
		0, 0, 0,
		//stop.X, stop.Y, stop.Z,
	)

	colors := math32.NewArrayF32(0, 3)
	colors.Append(
		1, 1, 1,
		//color.R, color.G, color.B,
	)
	geom.AddVBO(gls.NewVBO(vertices).AddAttrib(gls.VertexPosition))
	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))

	return geom
}
