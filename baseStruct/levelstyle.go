package baseStruct

import (
	"github.com/g3n/engine/geometry"
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

	MakeWhiteNeuron func() *graphic.Mesh
	MakeBlackNeuron func() *graphic.Mesh
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

	// Create functions that return a cube mesh using the provided material, reusing the same cube geometry

	sharedCubeGeom := geometry.NewCube(1)
	makeObjWithMaterial := func(mat *material.Phong) func() *graphic.Mesh {
		return func() *graphic.Mesh { return graphic.NewMesh(sharedCubeGeom, mat) }
	}

	s.MakeWhiteNeuron = makeObjWithMaterial(s.materialWhite)
	s.MakeBlackNeuron = makeObjWithMaterial(s.materialBlack)

	return s
}
