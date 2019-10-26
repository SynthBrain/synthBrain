package app

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/util/helper"
	"time"
)

func init() {
	DemoMap["other.skybox"] = &Skybox{}
}

type Skybox struct{}

// Start is called once at the start of the demo.
func (t *Skybox) Start(a *App) {

	// Create Skybox
	skybox, err := graphic.NewSkybox(graphic.SkyboxData{
		a.DirData() + "/images/sanfrancisco/", "jpg",
		[6]string{"posx", "negx", "posy", "negy", "posz", "negz"}})
	if err != nil {
		panic(err)
	}
	a.Scene().Add(skybox)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)
}

// Update is called every frame.
func (t *Skybox) Update(a *App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *Skybox) Cleanup(a *App) {}
