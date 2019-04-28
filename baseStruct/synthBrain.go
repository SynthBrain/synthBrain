package baseStruct

import (
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

	//UserData *UserData

	Root     *gui.Root
	Menu     *gui.Panel
	Main     *gui.Panel
	Controls *gui.Panel

	StepDelta *math32.Vector2

	LoadingLabel *gui.ImageLabel

	RestartButton    *gui.ImageButton
	MenuButton       *gui.ImageButton
	QuitButton       *gui.ImageButton
	PlayButton       *gui.ImageButton
	FullScreenButton *gui.ImageButton

	LevelScene *core.Node
	LevelStyle *LevelStyle
	Levels     []*Level
	LevelsRaw  []string
	Level      *Level
	Leveln     int
}

// RestartLevel restarts the current level
func (synB *SynthBrain) RestartLevel(playSound bool) {
	//log.Debug("Restart Level")

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
	//if !synB.Menu.Visible() {
	//	fmt.Println("menu visible")
	//}

	//synB.Levels[synB.Leveln].Restart(playSound)
}

// ToggleFullScreen toggles whether is game is fullscreen or windowed
func (synB *SynthBrain) ToggleFullScreen() {
	//log.Debug("Toggle FullScreen")

	synB.Win.SetFullScreen(!synB.Win.FullScreen())
}

// ToggleMenu switched the menu, title, and credits overlay for the in-level corner buttons
func (synB *SynthBrain) ToggleMenu() {
	//log.Debug("Toggle Menu")

	if synB.Menu.Visible() {

		// Dispatch OnMouseUp to clear the orbit control if user had mouse button pressed when they pressed Esc to hide menu
		synB.Win.Dispatch(gui.OnMouseUp, &window.MouseEvent{})

		synB.Menu.SetVisible(false)
		synB.Controls.SetVisible(true)
		synB.OrbitControl.Enabled = true

	} else {
		synB.Menu.SetVisible(true)
		synB.Controls.SetVisible(false)
		synB.OrbitControl.Enabled = false
	}
}

// Quit saves the user data and quits the game
func (synB *SynthBrain) Quit() {
	//log.Debug("Quit")

	// Copy settings into user data and save
	//synB.userData.SfxVol = synB.sfxSlider.Value()
	//synB.userData.MusicVol = synB.musicSlider.Value()
	//synB.userData.FullScreen = synB.win.FullScreen()
	//synB.userData.Save(synB.dataDir)

	// Close the window
	synB.Win.SetShouldClose(true)
}

// onKey handles keyboard events for the game
func (synB *SynthBrain) OnKey(evname string, ev interface{}) {

	kev := ev.(*window.KeyEvent)
	switch kev.Keycode {
	case window.KeyEscape:
		synB.ToggleMenu()
	case window.KeyF:
		synB.ToggleFullScreen()
	}
}

// onMouse handles mouse events for the game
func (synB *SynthBrain) OnMouse(evname string, ev interface{}) {
	//mev := ev.(*window.MouseEvent)
	//
	//if g.gopherLocked == false && g.leveln > 0 {
	//	// Mouse button pressed
	//	if mev.Action == window.Press {
	//		// Left button pressed
	//		if mev.Button == window.MouseButtonLeft {
	//			g.arrowNode.SetVisible(true)
	//		}
	//	} else if mev.Action == window.Release {
	//		g.arrowNode.SetVisible(false)
	//	}
	//}
}

// onCursor handles cursor movement for the game
func (synB *SynthBrain) OnCursor(evname string, ev interface{}) {

	// Calculate direction of potential movement based on camera angle
	var dir math32.Vector3
	synB.Camera.WorldDirection(&dir)
	synB.StepDelta.Set(0, 0)

	if math32.Abs(dir.Z) > math32.Abs(dir.X) {
		if dir.Z > 0 {
			//synB.arrowNode.SetRotationY(3 * math32.Pi / 2)
			synB.StepDelta.Y = 1
		} else {
			//synB.arrowNode.SetRotationY(1 * math32.Pi / 2)
			synB.StepDelta.Y = -1
		}
	} else {
		if dir.X > 0 {
			//synB.arrowNode.SetRotationY(4 * math32.Pi / 2)
			synB.StepDelta.X = 1
		} else {
			//synB.arrowNode.SetRotationY(2 * math32.Pi / 2)
			synB.StepDelta.X = -1
		}
	}

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

	// Always enable the button to return to the previous level except when we are in the very first level
	//synB.PrevButton.SetEnabled(n != 0)

	// The button to go to the next level has 3 different states: disabled, locked and enabled
	// If this is the very last level - disable it completely
	//if n == len(synB.Levels)-1 {
	//	synB.nextButton.SetImage(gui.ButtonDisabled, synB.DataDir + "/assets/right_disabled2.png")
	//	synB.nextButton.SetEnabled(false)
	//} else {
	//	synB.nextButton.SetImage(gui.ButtonDisabled, synB.DataDir + "/assets/right_disabled_locked.png")
	//	// check last completed level
	//	if synB.UserData.LastUnlockedLevel == n {
	//		synB.nextButton.SetEnabled(false)
	//	} else {
	//		synB.nextButton.SetEnabled(true)
	//	}
	//}

	// Remove level.scene from levelScene and unsubscribe from events
	if len(synB.LevelScene.Children()) > 0 {
		synB.LevelScene.Remove(synB.Level.scene)
		synB.Win.UnsubscribeID(window.OnKeyDown, synB.Leveln)
	}

	// Update current level index and level reference
	synB.Leveln = n
	//synB.userData.LastLevel = n
	synB.Level = synB.Levels[synB.Leveln]

	synB.RestartLevel(false)
	//synB.Level.gopherNodeRotate.Add(g.gopherNode)
	//synB.Level.gopherNodeTranslate.Add(g.arrowNode)
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

/*
func (synB *SynthBrain) SetupGui(width, height int) {
	log.Debug("Creating GUI...")

	transparent := math32.Color4{0, 0, 0, 0}
	blackTextColor := math32.Color4{0.3, 0.3, 0.3, 1.0}
	creditsColor := math32.Color{0.6, 0.6, 0.6}
	sliderColor := math32.Color4{0.628, 0.882, 0.1, 1}
	sliderColorOff := math32.Color4{0.82, 0.48, 0.48, 1}
	sliderColorOver := math32.Color4{0.728, 0.982, 0.2, 1}
	sliderBorderColor := math32.Color4{0.71, 0.482, 0.26, 1}

	sliderBorder := gui.RectBounds{3, 3, 3, 3}
	//zeroBorder := gui.RectBounds{0, 0, 0, 0}

	s := gui.StyleDefault()
	s.ImageButton = gui.ImageButtonStyles{}
	s.ImageButton.Normal = gui.ImageButtonStyle{}
	s.ImageButton.Normal.BgColor = transparent
	s.ImageButton.Normal.FgColor = blackTextColor
	s.ImageButton.Over = s.ImageButton.Normal
	s.ImageButton.Focus = s.ImageButton.Normal
	s.ImageButton.Pressed = s.ImageButton.Normal
	s.ImageButton.Disabled = s.ImageButton.Normal

	s.Slider = gui.SliderStyles{}
	s.Slider.Normal = gui.SliderStyle{}
	s.Slider.Normal.Border = sliderBorder
	s.Slider.Normal.BorderColor = sliderBorderColor
	s.Slider.Normal.BgColor = math32.Color4{0.2, 0.2, 0.2, 1}
	s.Slider.Normal.FgColor = sliderColor
	s.Slider.Over = s.Slider.Normal
	s.Slider.Over.BgColor = math32.Color4{0.3, 0.3, 0.3, 1}
	s.Slider.Over.FgColor = sliderColorOver
	s.Slider.Focus = s.Slider.Over
	s.Slider.Disabled = s.Slider.Normal
	s.Slider.Disabled.FgColor = sliderColorOff

	var err error

	hoverSound := func(evname string, ev interface{}) {
		g.PlaySound(g.hoverPlayer, nil)
	}

	// Menu
	g.menu = gui.NewPanel(100, 100)
	g.menu.SetColor4(&math32.Color4{0.1, 0.1, 0.1, 0.6})
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.menu.SetWidth(g.root.ContentWidth())
		g.menu.SetHeight(g.root.ContentHeight())
	})

	// Controls
	g.controls = gui.NewPanel(100, 100)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.controls.SetWidth(g.root.ContentWidth())
		g.controls.SetHeight(g.root.ContentHeight())
	})

	// Header panel
	header := gui.NewPanel(0, 0)
	header.SetPosition(0, 0)
	header.SetLayout(gui.NewHBoxLayout())
	header.SetPaddings(20, 20, 20, 20)
	header.SetSize(float32(width), 160)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		header.SetWidth(g.root.ContentWidth())
	})
	g.controls.Add(header)

	// Previous Level Button
	g.prevButton, err = gui.NewImageButton(g.dataDir + "/gui/left_normal.png")
	g.prevButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/left_hover.png")
	g.prevButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/left_click.png")
	g.prevButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/left_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.prevButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.PreviousLevel()
	})
	g.prevButton.Subscribe(gui.OnCursorEnter, hoverSound)
	header.Add(g.prevButton)

	params := gui.HBoxLayoutParams{Expand: 1, AlignV: gui.AlignCenter}

	spacer1 := gui.NewPanel(0, 0)
	spacer1.SetLayoutParams(&params)
	header.Add(spacer1)

	// Level Number Label
	g.levelLabel, err = gui.NewImageButton(g.dataDir + "/gui/panel.png")
	g.levelLabel.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/panel.png")
	g.levelLabel.SetColor(&math32.Color{0.8, 0.8, 0.8})
	g.levelLabel.SetText("Level")
	g.levelLabel.SetFontSize(35)
	g.levelLabel.SetEnabled(false)
	header.Add(g.levelLabel)

	spacer2 := gui.NewPanel(0, 0)
	spacer2.SetLayoutParams(&params)
	header.Add(spacer2)

	// Next Level Button
	g.nextButton, err = gui.NewImageButton(g.dataDir + "/gui/right_normal.png")
	g.nextButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/right_hover.png")
	g.nextButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/right_click.png")
	g.nextButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/right_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.nextButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.NextLevel()
	})
	g.nextButton.Subscribe(gui.OnCursorEnter, hoverSound)
	header.Add(g.nextButton)

	// Footer panel
	footer := gui.NewPanel(0, 0)
	footer_height := 140
	footer.SetLayout(gui.NewHBoxLayout())
	footer.SetPaddings(20, 20, 20, 20)
	footer.SetSize(g.root.ContentHeight(), float32(footer_height))
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		footer.SetWidth(g.root.ContentWidth())
		footer.SetPositionY(g.root.ContentHeight() - float32(footer_height))
	})
	g.controls.Add(footer)

	// Restart Level Button
	g.restartButton, err = gui.NewImageButton(g.dataDir + "/gui/restart_normal.png")
	g.restartButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/restart_hover.png")
	g.restartButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/restart_click.png")
	g.restartButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/restart_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.restartButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.RestartLevel(true)
	})
	g.restartButton.Subscribe(gui.OnCursorEnter, hoverSound)
	footer.Add(g.restartButton)

	spacer3 := gui.NewPanel(0, 0)
	spacer3.SetLayoutParams(&params)
	footer.Add(spacer3)

	// Restart Level Button
	g.menuButton, err = gui.NewImageButton(g.dataDir + "/gui/menu_normal.png")
	g.menuButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/menu_hover.png")
	g.menuButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/menu_click.png")
	g.menuButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/menu_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.menuButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.ToggleMenu()
	})
	g.menuButton.Subscribe(gui.OnCursorEnter, hoverSound)
	footer.Add(g.menuButton)

	g.controls.SetVisible(false)
	g.root.Add(g.controls)

	// Title
	g.titleImage, err = gui.NewImageButton(g.dataDir + "/gui/title3.png")
	g.titleImage.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/title3.png")
	g.titleImage.SetEnabled(false)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.titleImage.SetPositionX((g.root.ContentWidth() - g.titleImage.ContentWidth()) / 2)
	})
	g.menu.Add(g.titleImage)

	// Loading Text
	g.loadingLabel = gui.NewImageLabel("Loading...")
	g.loadingLabel.SetColor(&math32.Color{1, 1, 1})
	g.loadingLabel.SetFontSize(40)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.loadingLabel.SetPositionX((g.root.ContentWidth() - g.loadingLabel.ContentWidth()) / 2)
		g.loadingLabel.SetPositionY((g.root.ContentHeight() - g.loadingLabel.ContentHeight()) / 2)
	})
	g.root.Add(g.loadingLabel)

	// Instructions
	g.instructions1 = gui.NewImageLabel(INSTRUCTIONS_LINE1)
	g.instructions1.SetColor(&creditsColor)
	g.instructions1.SetFontSize(28)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.instructions1.SetWidth(g.root.ContentWidth())
		g.instructions1.SetPositionY(4 * g.instructions1.ContentHeight())
	})
	g.controls.Add(g.instructions1)

	g.instructions2 = gui.NewImageLabel(INSTRUCTIONS_LINE2)
	g.instructions2.SetColor(&creditsColor)
	g.instructions2.SetFontSize(28)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.instructions2.SetWidth(g.root.ContentWidth())
		g.instructions2.SetPositionY(5 * g.instructions2.ContentHeight())
	})
	g.controls.Add(g.instructions2)

	g.instructions3 = gui.NewImageLabel(INSTRUCTIONS_LINE3)
	g.instructions3.SetColor(&creditsColor)
	g.instructions3.SetFontSize(28)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.instructions3.SetWidth(g.root.ContentWidth())
		g.instructions3.SetPositionY(g.root.ContentHeight() - 2*g.instructions3.ContentHeight())
	})
	g.controls.Add(g.instructions3)

	buttonInstructionsPad := float32(24)

	g.instructionsRestart = gui.NewImageLabel("Restart Level (R)")
	g.instructionsRestart.SetColor(&creditsColor)
	g.instructionsRestart.SetFontSize(20)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.instructionsRestart.SetPosition(buttonInstructionsPad, g.root.ContentHeight()-6*g.instructionsRestart.ContentHeight())
	})
	g.controls.Add(g.instructionsRestart)

	g.instructionsMenu = gui.NewImageLabel("Show Menu (Esc)")
	g.instructionsMenu.SetColor(&creditsColor)
	g.instructionsMenu.SetFontSize(20)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.instructionsMenu.SetPosition(g.root.ContentWidth()-g.instructionsMenu.ContentWidth()-buttonInstructionsPad, g.root.ContentHeight()-6*g.instructionsMenu.ContentHeight())
	})
	g.controls.Add(g.instructionsMenu)

	// Main panel
	g.main = gui.NewPanel(600, 300)
	mainLayout := gui.NewVBoxLayout()
	mainLayout.SetAlignV(gui.AlignHeight)
	g.main.SetLayout(mainLayout)
	g.main.SetBorders(2, 2, 2, 2)
	g.main.SetBordersColor4(&sliderBorderColor)
	g.main.SetColor4(&math32.Color4{0.2, 0.2, 0.2, 0.6})
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g.main.SetPositionX((g.root.Width() - g.main.Width()) / 2)
		g.main.SetPositionY((g.root.Height()-g.main.Height())/2 + 50)
	})

	topRow := gui.NewPanel(g.main.ContentWidth(), 100)
	topRowLayout := gui.NewHBoxLayout()
	topRowLayout.SetAlignH(gui.AlignWidth)
	topRow.SetLayout(topRowLayout)
	alignCenterVerical := gui.HBoxLayoutParams{Expand: 0, AlignV: gui.AlignCenter}

	// Music Control
	musicControl := gui.NewPanel(130, 100)
	musicControl.SetLayout(topRowLayout)

	g.musicButton, err = gui.NewImageButton(g.dataDir + "/gui/music_normal.png")
	g.musicButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/music_hover.png")
	g.musicButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/music_click.png")
	g.musicButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/music_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.musicButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.userData.MusicOn = !g.userData.MusicOn
		g.UpdateMusicButton(g.userData.MusicOn)
	})
	g.musicButton.Subscribe(gui.OnCursorEnter, hoverSound)
	musicControl.Add(g.musicButton)

	// Music Volume Slider
	g.musicSlider = gui.NewVSlider(20, 80)
	g.musicSlider.SetValue(g.userData.MusicVol)
	g.musicSlider.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		g.SetMusicVolume(g.musicSlider.Value())
	})
	g.musicSlider.Subscribe(gui.OnCursorEnter, hoverSound)
	g.musicSlider.SetLayoutParams(&alignCenterVerical)
	musicControl.Add(g.musicSlider)

	topRow.Add(musicControl)

	// Sound Effects Control
	sfxControl := gui.NewPanel(130, 100)
	sfxControl.SetLayout(topRowLayout)

	g.sfxButton, err = gui.NewImageButton(g.dataDir + "/gui/sound_normal.png")
	g.sfxButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/sound_hover.png")
	g.sfxButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/sound_click.png")
	g.sfxButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/sound_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.sfxButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.userData.SfxOn = !g.userData.SfxOn
		g.UpdateSfxButton(g.userData.SfxOn)
	})
	g.sfxButton.Subscribe(gui.OnCursorEnter, hoverSound)
	sfxControl.Add(g.sfxButton)

	// Sound Effects Volume Slider
	g.sfxSlider = gui.NewVSlider(20, 80)
	g.sfxSlider.SetValue(g.userData.SfxVol)
	g.sfxSlider.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		g.SetSfxVolume(3 * g.sfxSlider.Value())
	})
	g.sfxSlider.Subscribe(gui.OnCursorEnter, hoverSound)
	g.sfxSlider.SetLayoutParams(&alignCenterVerical)
	sfxControl.Add(g.sfxSlider)

	topRow.Add(sfxControl)

	// FullScreen Button
	g.fullScreenButton, err = gui.NewImageButton(g.dataDir + "/gui/screen_normal.png")
	g.fullScreenButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/screen_hover.png")
	g.fullScreenButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/screen_click.png")
	g.fullScreenButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/screen_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.fullScreenButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.ToggleFullScreen()
	})
	g.fullScreenButton.Subscribe(gui.OnCursorEnter, hoverSound)
	topRow.Add(g.fullScreenButton)

	g.main.Add(topRow)

	buttonRow := gui.NewPanel(g.main.ContentWidth(), 100)
	buttonRowLayout := gui.NewHBoxLayout()
	buttonRowLayout.SetAlignH(gui.AlignWidth)
	buttonRow.SetLayout(buttonRowLayout)

	// Quit Button
	g.quitButton, err = gui.NewImageButton(g.dataDir + "/gui/quit_normal.png")
	g.quitButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/quit_hover.png")
	g.quitButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/quit_click.png")
	g.quitButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/quit_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.quitButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.Quit()
	})
	g.quitButton.Subscribe(gui.OnCursorEnter, hoverSound)
	buttonRow.Add(g.quitButton)

	// Play Button
	g.playButton, err = gui.NewImageButton(g.dataDir + "/gui/play_normal.png")
	g.playButton.SetImage(gui.ButtonOver, g.dataDir + "/gui/play_hover.png")
	g.playButton.SetImage(gui.ButtonPressed, g.dataDir + "/gui/play_click.png")
	g.playButton.SetImage(gui.ButtonDisabled, g.dataDir + "/gui/play_disabled2.png")
	if err != nil {
		panic(err)
	}
	g.playButton.Subscribe(gui.OnMouseUp, func(evname string, ev interface{}) {
		g.PlaySound(g.clickPlayer, nil)
		g.ToggleMenu()
	})
	g.playButton.Subscribe(gui.OnCursorEnter, hoverSound)
	buttonRow.Add(g.playButton)

	g.main.Add(buttonRow)

	// Add credits labels
	lCredits1 := gui.NewImageLabel(CREDITS_LINE1)
	lCredits1.SetColor(&creditsColor)
	lCredits1.SetFontSize(20)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		lCredits1.SetWidth(g.root.ContentWidth())
		lCredits1.SetPositionY(g.root.ContentHeight() - 2*lCredits1.ContentHeight())
	})
	g.menu.Add(lCredits1)

	lCredits2 := gui.NewImageLabel(CREDITS_LINE2)
	lCredits2.SetColor(&creditsColor)
	lCredits2.SetFontSize(20)
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		lCredits2.SetWidth(g.root.ContentWidth())
		lCredits2.SetPositionY(g.root.ContentHeight() - lCredits2.ContentHeight())
	})
	g.menu.Add(lCredits2)

	g3n := gui.NewImageLabel("")
	g3n.SetSize(57, 50)
	g3n.SetImageFromFile(g.dataDir + "/img/g3n.png")
	g.root.Subscribe(gui.OnResize, func(evname string, ev interface{}) {
		g3n.SetPositionX(g.root.ContentWidth() - g3n.Width())
		g3n.SetPositionY(g.root.ContentHeight() - 1.3*g3n.Height())
	})
	g.menu.Add(g3n)

	g.root.Add(g.menu)

	// Dispatch a fake OnResize event to update all subscribed elements
	g.root.Dispatch(gui.OnResize, nil)

	log.Debug("Done creating GUI.")
}*/
