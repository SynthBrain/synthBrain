package baseStruct

import (
	"github.com/SynthBrain/synthBrain/drawing3D"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/camera/control"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/window"
)

// SynthBrain base struct application
type SynthBrain struct {
	Wmgr         window.IWindowManager
	Win          window.IWindow
	Gs           *gls.GLS
	Renderer     *renderer.Renderer
	Scene        *core.Node
	Camera       *camera.Perspective
	OrbitControl *control.OrbitControl
	DataDir      string

	Restart  *gui.Button
	Exit     *gui.Button
	WebCam   *gui.Button
	LabelFps *gui.Label

	Root *gui.Root

	StepDelta *math32.Vector2

	LoadingLabel *gui.ImageLabel

	LevelScene *core.Node
	LevelStyle *drawing3D.LevelStyle
	Levels     []*Level
	LevelsRaw  []string
	Level      *Level
	Leveln     int
}

// Update updates the current level if any
func (synB *SynthBrain) Update(timeDelta float64) {
	if synB.Level != nil {
		synB.Level.Update(timeDelta)
	}
}

// InitLevel initializes the level associated to the provided index
func (synB *SynthBrain) InitLevel(n int) {
	//log.Debug("Initializing Level %v", n+1)

	// Remove level.scene from levelScene and unsubscribe from events
	if len(synB.LevelScene.Children()) > 0 {
		synB.LevelScene.Remove(synB.Level.scene)
		synB.Win.UnsubscribeID(window.OnKeyDown, synB.Leveln)
	}

	// Update current level index and level reference
	synB.Leveln = n
	//synB.userData.LastLevel = n
	synB.Level = synB.Levels[synB.Leveln]

	//synB.LevelLabel.SetText("Level " + strconv.Itoa(n+1))
	synB.LevelScene.Add(synB.Level.scene)
	//synB.Win.SubscribeID(window.OnKeyDown, synB.Leveln, synB.Level.onKey)
}

// LoadSkybox loads the space skybox and adds it to the scene
func (synB *SynthBrain) LoadSkyBox() {
	//log.Debug("Creating Skybox...")

	// Load skybox textures
	skyboxData := graphic.SkyboxData{
		synB.DataDir + "/assets/skybox/", "jpg",
		[6]string{"px", "nx", "py", "ny", "pz", "nz"}}

	skybox, err := graphic.NewSkybox(skyboxData)
	if err != nil {
		panic(err)
	}
	skybox.SetRenderOrder(-1) // The skybox should always be rendered first

	//// For each skybox face - set the material to not use lights and to have emissive color.
	//brightness := float32(0.6)
	//sbmats := skybox.Materials()
	//for i := 0; i < len(sbmats); i++ {
	//	sbmat := sbmats[i].GetMaterial().(*material.Standard)
	//	sbmat.SetUseLights(material.UseLightNone)
	//	sbmat.SetEmissiveColor(&math32.Color{brightness, brightness, brightness})
	//}
	synB.Scene.Add(skybox)

	//log.Debug("Done creating skybox")
}

// LoadLevels reads and parses the level files inside ./levels, building an array of Level objects
func (synB *SynthBrain) LoadLevels() {
	//log.Debug("Load Levels")

	synB.Levels = make([]*Level, 1)
	synB.Levels[0] = NewLevel(synB, synB.LevelStyle, synB.Camera)
}

// RenderFrame renders a frame of the scene with the GUI overlaid
func (synB *SynthBrain) RenderFrame() {

	// Process GUI timers
	synB.Root.TimerManager.ProcessTimers()

	// Render the scene/gui using the specified camera
	rendered, err := synB.Renderer.Render(synB.Camera)
	if err != nil {
		panic(err)
	}

	// Check I/O events
	synB.Wmgr.PollEvents()

	// Update window if necessary
	if rendered {
		synB.Win.SwapBuffers()
	}
}

// // Create a tic-tac-toe board.
// board := [][]string{
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// }

// for i := 0; i < len(board); i++ {
// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
// }
