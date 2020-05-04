package app

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"time"
)

func init() {
	DemoMap["A_testLogic"] = &testLogic{}
}

type testLogic struct {
	points *graphic.Points
}

// Start is called once at the start of the demo.
func (t *testLogic) Start(a *App) {

	a.Gls().ClearColor(0, 0, 0, 1)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)

	// Load textures for the sprites
	spnames := []string{"snowflake1.png"}
	sprites := []*texture.Texture2D{}
	for _, name := range spnames {
		tex, err := texture.NewTexture2DFromImage(a.DirData() + "/images/" + name)
		if err != nil {
			a.Log().Fatal("Error loading texture: %s", err)
		}
		sprites = append(sprites, tex)
	}

	// Creates geometry with random points
	geom := geometry.NewGeometry()
	positions := math32.NewArrayF32(0, 0)
	numPoints := 100
	var vertex math32.Vector3
	var temp float32
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (j % 2.0 == 1) {
				temp = 0.5
			}
			var tempI float32
			tempI += float32(i)+ temp
			var tempZ float32
			tempZ += float32(j)
			vertex.Set(
				float32(tempZ),
				0,
				float32(tempI),
			)
			positions.AppendVector3(&vertex)
			temp = 0
		}
	}
	geom.AddVBO(gls.NewVBO(positions).AddAttrib(gls.VertexPosition))

	t.points = graphic.NewPoints(geom, nil)
	start := 0
	count := numPoints / len(sprites)
	for _, tex := range sprites {
		mat := material.NewPoint(&math32.Color{1, 1, 1})
		mat.SetTransparent(true)
		mat.SetOpacity(0.6)
		mat.AddTexture(tex)
		mat.SetSize(500)
		mat.SetBlending(material.BlendAdditive)
		mat.SetDepthMask(false)
		t.points.AddMaterial(t.points, mat, start, count)
		start += count
	}
	t.points.SetScale(1,1,1)
	a.Scene().Add(t.points)
}

// Update is called every frame.
func (t *testLogic) Update(a *App, deltaTime time.Duration) {

	//rps := float32(deltaTime.Seconds()) * 2 * math32.Pi
	//t.points.RotateY(rps * 0.05)
}

// Cleanup is called once at the end of the demo.
func (t *testLogic) Cleanup(a *App) {}
