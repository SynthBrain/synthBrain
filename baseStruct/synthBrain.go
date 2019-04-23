package baseStruct


import (
	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/camera/control"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/window"
)

// SynthBrain base struct application
type SynthBrain struct {
	wmgr         window.IWindowManager
	win          window.IWindow
	gs           *gls.GLS
	renderer     *renderer.Renderer
	scene        *core.Node
	camera       *camera.Perspective
	orbitControl *control.OrbitControl
	dataDir      string

	userData *UserData

	root     *gui.Root
	menu     *gui.Panel
	main     *gui.Panel
	controls *gui.Panel

	stepDelta     *math32.Vector2
	musicCheckbox *gui.CheckRadio
	musicSlider   *gui.Slider

	sfxCheckbox *gui.CheckRadio
	sfxSlider   *gui.Slider

	loadingLabel        *gui.ImageLabel

	levelLabel       *gui.ImageButton
	titleImage       *gui.ImageButton
	nextButton       *gui.ImageButton
	prevButton       *gui.ImageButton
	restartButton    *gui.ImageButton
	menuButton       *gui.ImageButton
	quitButton       *gui.ImageButton
	playButton       *gui.ImageButton
	sfxButton        *gui.ImageButton
	musicButton      *gui.ImageButton
	fullScreenButton *gui.ImageButton

	levelScene *core.Node
	levelStyle *LevelStyle
	levels     []*Level
	levelsRaw  []string
	level      *Level
	leveln     int

	gopherLocked   bool
	gopherNode     *core.Node
	arrowNode      *core.Node
	steps          int
	audioAvailable bool

	// Sound/music players
	musicPlayer           *audio.Player
	musicPlayerMenu       *audio.Player
	clickPlayer           *audio.Player
	hoverPlayer           *audio.Player
	walkPlayer            *audio.Player
	bumpPlayer            *audio.Player
	gopherHurtPlayer      *audio.Player
	gopherFallEndPlayer   *audio.Player
	gopherFallStartPlayer *audio.Player
	boxPushPlayer         *audio.Player
	boxOnPadPlayer        *audio.Player
	boxOffPadPlayer       *audio.Player
	boxFallEndPlayer      *audio.Player
	boxFallStartPlayer    *audio.Player
	elevatorUpPlayer      *audio.Player
	elevatorDownPlayer    *audio.Player
	levelDonePlayer       *audio.Player
	levelRestartPlayer    *audio.Player
	levelFailPlayer       *audio.Player
	gameCompletePlayer    *audio.Player
}

// RestartLevel restarts the current level
func (synB *SynthBrain) RestartLevel(playSound bool) {
	log.Debug("Restart Level")

	// if synB.leveln == 0 {
	// 	synB.instructions3.SetText(INSTRUCTIONS_LINE3)
	// }

	// synB.instructions1.SetVisible(synB.leveln == 0)
	// synB.instructions2.SetVisible(synB.leveln == 0)
	// synB.instructions3.SetVisible(synB.leveln == 0)
	// synB.instructionsRestart.SetVisible(synB.leveln == 0)
	// synB.instructionsMenu.SetVisible(synB.leveln == 0)
	// synB.arrowNode.SetVisible(synB.leveln == 0)

	// If the menu is not visible then "free" the gopher
	// The menu would be visible if the user fell or dropped a box and then opened the menu before the fall ended
	// If the menu is visible then we want to keep the gopher locked
	if !synB.menu.Visible() {
		synB.gopherLocked = false
	}

	synB.levels[synB.leveln].Restart(playSound)
}

// ToggleFullScreen toggles whether is game is fullscreen or windowed
func (synB *SynthBrain) ToggleFullScreen() {
	log.Debug("Toggle FullScreen")

	synB.win.SetFullScreen(!synB.win.FullScreen())
}

// ToggleMenu switched the menu, title, and credits overlay for the in-level corner buttons
func (synB *SynthBrain) ToggleMenu() {
	log.Debug("Toggle Menu")

	if synB.menu.Visible() {

		// Dispatch OnMouseUp to clear the orbit control if user had mouse button pressed when they pressed Esc to hide menu
		synB.win.Dispatch(gui.OnMouseUp, &window.MouseEvent{})

		// Dispatch OnCursorLeave to sliders in case user had cursor over sliders when they pressed Esc to hide menu
		synB.sfxSlider.Dispatch(gui.OnCursorLeave, &window.MouseEvent{})
		synB.musicSlider.Dispatch(gui.OnCursorLeave, &window.MouseEvent{})

		synB.menu.SetVisible(false)
		synB.controls.SetVisible(true)
		synB.orbitControl.Enabled = true
		synB.gopherLocked = false
		if synB.audioAvailable {
			synB.musicPlayerMenu.Stop()
			synB.musicPlayer.Play()
		}
	} else {
		synB.menu.SetVisible(true)
		synB.controls.SetVisible(false)
		synB.orbitControl.Enabled = false
		synB.gopherLocked = true
		if synB.audioAvailable {
			synB.musicPlayer.Stop()
			synB.musicPlayerMenu.Play()
		}
	}
}

// Quit saves the user data and quits the game
func (synB *SynthBrain) Quit() {
	log.Debug("Quit")

	// Copy settings into user data and save
	synB.userData.SfxVol = synB.sfxSlider.Value()
	synB.userData.MusicVol = synB.musicSlider.Value()
	synB.userData.FullScreen = synB.win.FullScreen()
	synB.userData.Save(synB.dataDir)

	// Close the window
	synB.win.SetShouldClose(true)
}

// onKey handles keyboard events for the game
func (synB *SynthBrain) onKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	switch kev.Keycode {
	case window.KeyEscape:
		g.ToggleMenu()
	case window.KeyF:
		g.ToggleFullScreen()
	case window.KeyR:
		if !g.menu.Visible() && g.steps > 0 {
			g.RestartLevel(true)
		}
	}
}

// onMouse handles mouse events for the game
func (synB *SynthBrain) onMouse(evname string, ev interface{}) {
	mev := ev.(*window.MouseEvent)

	if g.gopherLocked == false && g.leveln > 0 {
		// Mouse button pressed
		if mev.Action == window.Press {
			// Left button pressed
			if mev.Button == window.MouseButtonLeft {
				g.arrowNode.SetVisible(true)
			}
		} else if mev.Action == window.Release {
			g.arrowNode.SetVisible(false)
		}
	}
}

// onCursor handles cursor movement for the game
func (synB *SynthBrain) onCursor(evname string, ev interface{}) {

	// Calculate direction of potential movement based on camera angle
	var dir math32.Vector3
	g.camera.WorldDirection(&dir)
	g.stepDelta.Set(0, 0)

	if math32.Abs(dir.Z) > math32.Abs(dir.X) {
		if dir.Z > 0 {
			g.arrowNode.SetRotationY(3 * math32.Pi / 2)
			g.stepDelta.Y = 1
		} else {
			g.arrowNode.SetRotationY(1 * math32.Pi / 2)
			g.stepDelta.Y = -1
		}
	} else {
		if dir.X > 0 {
			g.arrowNode.SetRotationY(4 * math32.Pi / 2)
			g.stepDelta.X = 1
		} else {
			g.arrowNode.SetRotationY(2 * math32.Pi / 2)
			g.stepDelta.X = -1
		}
	}

}

// Update updates the current level if any
func (synB *SynthBrain) Update(timeDelta float64) {
	if g.level != nil {
		g.level.Update(timeDelta)
	}
}


// InitLevel initializes the level associated to the provided index
func (synB *SynthBrain) InitLevel(n int) {
	log.Debug("Initializing Level %v", n+1)

	// Always enable the button to return to the previous level except when we are in the very first level
	g.prevButton.SetEnabled(n != 0)

	// The button to go to the next level has 3 different states: disabled, locked and enabled
	// If this is the very last level - disable it completely
	if n == len(g.levels)-1 {
		g.nextButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/right_disabled2.png")
		g.nextButton.SetEnabled(false)
	} else {
		g.nextButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/right_disabled_locked.png")
		// check last completed level
		if g.userData.LastUnlockedLevel == n {
			g.nextButton.SetEnabled(false)
		} else {
			g.nextButton.SetEnabled(true)
		}
	}

	// Remove level.scene from levelScene and unsubscribe from events
	if len(g.levelScene.Children()) > 0 {
		g.levelScene.Remove(g.level.scene)
		g.win.UnsubscribeID(window.OnKeyDown, g.leveln)
	}

	// Update current level index and level reference
	g.leveln = n
	g.userData.LastLevel = n
	g.level = g.levels[g.leveln]

	g.RestartLevel(false)
	g.level.gopherNodeRotate.Add(g.gopherNode)
	g.level.gopherNodeTranslate.Add(g.arrowNode)
	g.levelLabel.SetText("Level " + strconv.Itoa(n+1))
	g.levelScene.Add(g.level.scene)
	g.win.SubscribeID(window.OnKeyDown, g.leveln, g.level.onKey)

}

// LoadSkybox loads the space skybox and adds it to the scene
func (synB *SynthBrain) LoadSkyBox() {
	log.Debug("Creating Skybox...")

	// Load skybox textures
	skyboxData := graphic.SkyboxData{
		g.dataDir + "/img/skybox/", "jpg",
		[6]string{"px", "nx", "py", "ny", "pz", "nz"}}

	skybox, err := graphic.NewSkybox(skyboxData)
	if err != nil {
		panic(err)
	}
	skybox.SetRenderOrder(-1) // The skybox should always be rendered first

	// For each skybox face - set the material to not use lights and to have emissive color.
	brightness := float32(0.6)
	sbmats := skybox.Materials()
	for i := 0; i < len(sbmats); i++ {
		sbmat := skybox.Materials()[i].GetMaterial().(*material.Standard)
		sbmat.SetUseLights(material.UseLightNone)
		sbmat.SetEmissiveColor(&math32.Color{brightness, brightness, brightness})
	}
	g.scene.Add(skybox)

	log.Debug("Done creating skybox")
}

// RenderFrame renders a frame of the scene with the GUI overlaid
func (synB *SynthBrain) RenderFrame() {

	// Process GUI timers
	synB.root.TimerManager.ProcessTimers()

	// Render the scene/gui using the specified camera
	rendered, err := synB.renderer.Render(synB.camera)
	if err != nil {
		panic(err)
	}

	// Check I/O events
	synB.wmgr.PollEvents()

	// Update window if necessary
	if rendered {
		synB.win.SwapBuffers()
	}
}