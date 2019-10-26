package main

import (
	"github.com/SynthBrain/synthBrain/app"
	// _ "github.com/SynthBrain/synthBrain/demos/animation"
	// _ "github.com/SynthBrain/synthBrain/demos/audio"
	// _ "github.com/SynthBrain/synthBrain/demos//experimental/physics"
	// _ "github.com/SynthBrain/synthBrain/demos//geometry"
	// _ "github.com/SynthBrain/synthBrain/demos//gui"
	// _ "github.com/SynthBrain/synthBrain/demos//helper"
	// _ "github.com/SynthBrain/synthBrain/demos//light"
	// _ "github.com/SynthBrain/synthBrain/demos//loader"
	// _ "github.com/SynthBrain/synthBrain/demos//material"
	// _ "github.com/SynthBrain/synthBrain/demos//other"
	// _ "github.com/SynthBrain/synthBrain/demos//shader"
	// _ "github.com/SynthBrain/synthBrain/demos//tests"
	// _ "github.com/SynthBrain/synthBrain/demos//texture"
)

func main() {
	// Create and run application
	app.Create().Run()
}

//package main
//
//// 3840 * 2160 = 8 294 400
//import (
//	"flag"
//	"github.com/SynthBrain/synthBrain/baseStruct"
//	"github.com/SynthBrain/synthBrain/myGui"
//	"github.com/g3n/engine/app"
//	"github.com/g3n/engine/camera"
//	"github.com/g3n/engine/core"
//	"github.com/g3n/engine/geometry"
//	"github.com/g3n/engine/gls"
//	"github.com/g3n/engine/graphic"
//	"github.com/g3n/engine/gui"
//	"github.com/g3n/engine/light"
//	"github.com/g3n/engine/material"
//	"github.com/g3n/engine/math32"
//	"github.com/g3n/engine/renderer"
//
//	"github.com/g3n/engine/util/helper"
//	"github.com/g3n/engine/util/logger"
//	"github.com/g3n/engine/window"
//
//	"time"
//)
//
//var (
//	// TODO uncomment and implement usage of the following flags
//	//oFullScreen   = flag.Bool("fullscreen", false, "Starts application with full screen")
//	//oSwapInterval = flag.Int("swapinterval", -1, "Sets the swap buffers interval to this value")
//	oHideFPS     = flag.Bool("hidefps", false, "Do now show calculated FPS in the GUI")
//	oUpdateFPS   = flag.Uint("updatefps", 1000, "Time interval in milliseconds to update the FPS in the GUI")
//	oTargetFPS   = flag.Uint("targetfps", 60, "Sets the frame rate in frames per second")
//	oNoglErrors  = flag.Bool("noglerrors", false, "Do not check OpenGL errors at each call (may increase FPS)")
//	oCpuProfile  = flag.String("cpuprofile", "", "Activate cpu profiling writing profile to the specified file")
//	oExecTrace   = flag.String("exectrace", "", "Activate execution tracer writing data to the specified file")
//	oNogui       = flag.Bool("nogui", false, "Do not show the GUI, only the specified demo")
//	oLogs        = flag.String("logs", "", "Set log levels for packages. Ex: gui:debug,gls:info")
//	oStats       = flag.Bool("stats", false, "Shows statistics control panel in the GUI")
//	oRenderStats = flag.Bool("renderstats", false, "Shows gui renderer statistics in the console")
//)
//
//var log *logger.Logger
//
///*
//	Рисовать только тех что имеют достаточный уровень активность и окончательно не затухли
//*/
//func main() {
//	// Create application and scene
//	app := app.App()
//	scene := core.NewNode()
//
//	// Set the scene to be managed by the gui manager
//	gui.Manager().Set(scene)
//
//	// Create SynthBrain struct
//	synB := new(baseStruct.SynthBrain)
//	//frameRater := synB.FrameRater
//
//	// Create perspective camera
//	cam := camera.New(1)
//	cam.SetPosition(0, 0, 3)
//	scene.Add(cam)
//
//	// Set up orbit control for the camera
//	camera.NewOrbitControl(cam)
//
//	// Set up callback to update viewport and camera aspect ratio when the window is resized
//	onResize := func(evname string, ev interface{}) {
//		// Get framebuffer size and update viewport accordingly
//		width, height := app.GetSize()
//		app.Gls().Viewport(0, 0, int32(width), int32(height))
//		// Update the camera's aspect ratio
//		cam.SetAspect(float32(width) / float32(height))
//	}
//	app.Subscribe(window.OnWindowSize, onResize)
//	onResize("", nil)
//
//	// Create a blue torus and add it to the scene
//	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
//	geom := geometry.NewTorus(0, .4, 3, 3, math32.Pi*2)
//	mat := material.NewStandard(math32.NewColor("DarkBlue"))
//	mesh := graphic.NewMesh(geom, mat)
//	scene.Add(mesh)
//
//	// Create and add a button to the scene
//	btn := gui.NewButton("Make Red")
//	btn.SetPosition(30, 40)
//	btn.SetSize(40, 40)
//	btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
//		mat.SetColor(math32.NewColor("DarkRed"))
//	})
//	scene.Add(btn)
//	// Create and add a button to the scene
//	btn1 := gui.NewButton("Make Blue")
//	btn1.SetPosition(30, 90)
//	btn1.SetSize(40, 40)
//	btn1.Subscribe(gui.OnClick, func(name string, ev interface{}) {
//		mat.SetColor(math32.NewColor("DarkBlue"))
//	})
//	scene.Add(btn1)
//	// Create and add a button to the scene
//	onOff := false
//	chOnOffFlag := make(chan bool, 1)
//	exit :=  myGui.Exit(30, 240, &onOff, app, chOnOffFlag)
//	scene.Add(exit)
//
//	// Create and add a button to the scene
//	lbl := gui.NewLabel("FPS: ")
//	lbl.SetPosition(10, 10)
//	lbl.SetPaddings(2, 2, 2, 2)
//	scene.Add(lbl)
//
//	// Create and add lights to the scene
//	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
//	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 5.0)
//	pointLight.SetPosition(1, 0, 2)
//	scene.Add(pointLight)
//
//	// Create and add an axis helper to the scene
//	scene.Add(helper.NewAxes(0.5))
//
//	// Set background color to gray
//	app.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)
//
//	//synB.InitLevel(0)
//
//	now := time.Now()
//	newNow := time.Now()
//	//log.Info("Starting Render Loop")
//	//Run the application
//	app.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
//		app.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
//		newNow = time.Now()
//		timeDelta := now.Sub(newNow)
//		now = newNow
//		//fps, pfps, _ := frameRater.FPS(time.Duration(*oUpdateFPS) * time.Millisecond)
//		//
//		//lbl.SetText("FPS: ")
//		//lbl.SetText("FPS: " + fmt.Sprintf("%3.1f / %3.1f", fps, pfps) )
//		synB.Update(timeDelta.Seconds())
//		renderer.Render(scene, cam)
//	})
//
//	// ABROAD**********************************************************************************************
//
//	//// OpenGL functions must be executed in the same thread where
//	//// the context was created (by window.New())
//	//runtime.LockOSThread()
//	//
//	//// Parse command line flags
//	//showLog := flag.Bool("debug", false, "display the debug log")
//	//flag.Parse()
//	//
//	//// Create logger
//	//log = logger.New("SynthBrain", nil)
//	//log.AddWriter(logger.NewConsole(false))
//	//log.SetFormat(logger.FTIME | logger.FMICROS)
//	//if *showLog == true {
//	//	log.SetLevel(logger.DEBUG)
//	//} else {
//	//	log.SetLevel(logger.INFO)
//	//}
//	//log.Info("Initializing SynthBrain")
//	//
//	//// Create SynthBrain struct
//	//synB := new(baseStruct.SynthBrain)
//	//
//	//// Manually scan the $GOPATH directories to find the data directory
//	//rawPaths := os.Getenv("GOPATH")
//	//paths := strings.Split(rawPaths, ":")
//	//for _, j := range paths {
//	//	// Checks data path
//	//	path := filepath.Join(j, "src", "github.com", "SynthBrain", "synthBrain")
//	//	if _, err := os.Stat(path); err == nil {
//	//		synB.DataDir = path
//	//	}
//	//}
//	//
//	//// Get the window manager
//	//var err error
//	//synB.Wmgr, err = window.Manager("glfw")
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//// Create window and OpenGL context
//	//synB.Win, err = synB.Wmgr.CreateWindow(900, 640, "SynthBrain", false)
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//// Create OpenGL state
//	//synB.Gs, err = gls.New()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//// Speed up a bit by not checking OpenGL errors
//	//synB.Gs.SetCheckErrors(false)
//	//
//	//// Sets window background color
//	//synB.Gs.ClearColor(0, 0.2, 0.4, 1) //(0.1, 0.1, 0.1, 1.0)
//	//
//	//// Sets the OpenGL viewport size the same as the window size
//	//// This normally should be updated if the window is resized.
//	//width, height := synB.Win.Size()
//	//synB.Gs.Viewport(0, 0, int32(width), int32(height))
//	//
//	//// Creates GUI root panel
//	//synB.Root = gui.NewRoot(synB.Gs, synB.Win)
//	//synB.Root.SetSize(float32(width), float32(height))
//	//
//	//// Update window if resize
//	//synB.Win.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) {
//	//	width, height := synB.Win.Size()
//	//	synB.Gs.Viewport(0, 0, int32(width), int32(height))
//	//	synB.Root.SetSize(float32(width), float32(height))
//	//	aspect := float32(width) / float32(height)
//	//	synB.Camera.SetAspect(aspect)
//	//})
//	//
//	////add GUI*********************************************************
//	//// Create and add a label to the root panel
//	//synB.LabelFps = myGui.LabelFps(10, 10, "240")
//	//synB.Root.Add(synB.LabelFps)
//	//
//	//// Create and add button 1 to the root panel
//	//onOff := false
//	//chOnOffFlag := make(chan bool, 1)
//	//synB.WebCam = myGui.WebCam(10, 40, &onOff, chOnOffFlag)
//	//synB.Root.Add(synB.WebCam)
//	//
//	//// Create and add exit button to the root panel
//	//synB.Exit = myGui.Exit(10, 70, &onOff, synB.Win, chOnOffFlag)
//	//synB.Root.Add(synB.Exit)
//	////****************************************************************
//	//
//	//// Creates a renderer and adds default shaders
//	//synB.Renderer = renderer.NewRenderer(synB.Gs)
//	////g.renderer.SetSortObjects(false)
//	//err = synB.Renderer.AddDefaultShaders()
//	//if err != nil {
//	//	panic(err)
//	//}
//	//synB.Renderer.SetGui(synB.Root)
//	//
//	//// Adds a perspective camera to the scene
//	//// The camera aspect ratio should be updated if the window is resized.
//	//aspect := float32(width) / float32(height)
//	//synB.Camera = camera.NewPerspective(65, aspect, 0.01, 1000)
//	//synB.Camera.SetPosition(10, 10, 10)
//	//synB.Camera.LookAt(&math32.Vector3{0, 0, 0})
//	//
//	//// Create orbit control and set limits
//	//synB.OrbitControl = control.NewOrbitControl(synB.Camera, synB.Win)
//	//synB.OrbitControl.Enabled = true //false
//	//synB.OrbitControl.EnablePan = false
//	//synB.OrbitControl.MaxPolarAngle = 2 * math32.Pi / 3
//	//synB.OrbitControl.MinDistance = 0.1
//	//synB.OrbitControl.MaxDistance = 10000
//	//
//	//// Create main scene and child levelScene
//	//synB.Scene = core.NewNode()
//	//synB.LevelScene = core.NewNode()
//	//synB.Scene.Add(synB.Camera)
//	//synB.Scene.Add(synB.LevelScene)
//	//synB.StepDelta = math32.NewVector2(0, 0)
//	//synB.Renderer.SetScene(synB.Scene)
//	//
//	//// Add white ambient light to the scene
//	//ambLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.4)
//	//synB.Scene.Add(ambLight)
//	//
//	//synB.LevelStyle = drawing3D.NewBaseStyle()
//	//
//	////synB.SetupGui(width, height)
//	//synB.RenderFrame()
//	//
//	////synB.LoadSkyBox()
//	//synB.LoadLevels()
//	//
//	//size := 10
//	//gridHelp := graphic.NewGridHelper(float32(size), 1, math32.NewColor("LightGrey"))
//	//gridHelp.SetPosition(float32(size/2), -0.2, float32(size/2))
//	//synB.Scene.Add(gridHelp)
//	//
//	//// Done Loading - hide the loading label, show the menu, and initialize the level
//	////synB.LoadingLabel.SetVisible(false)
//	//
//	//synB.InitLevel(0)
//	//
//	//now := time.Now()
//	//newNow := time.Now()
//	//log.Info("Starting Render Loop")
//	//// Start the render loop
//	//for !synB.Win.ShouldClose() {
//	//	newNow = time.Now()
//	//	timeDelta := now.Sub(newNow)
//	//	now = newNow
//	//	//fmt.Println(now)
//	//
//	//	//vision.ReadImg(synB.DataDir, "/assets/0.jpg")
//	//	synB.Update(timeDelta.Seconds())
//	//	synB.RenderFrame()
//	//}
//
//	//fmt.Printf("app was running for %f \n", application.Get().RunSeconds())
//}
//
//// go func() {
//// 	for {
//// 		if a, b, c := app.FrameRater().FPS(60); a > 0 && b > 0 && c == true {
//// 			fmt.Println("FPS ", int(b))
//// 		}
//// 	}
//// }()
////fps := float32(app.FrameCount()) / application.Get().RunSeconds()
////go myGui.LabelFpsTest(10, 10, strconv.Itoa(int(app.FrameCount()) / int(application.Get().RunSeconds())), app)
//
////IndCh := make(chan int)
//
//// fmt.Println("Start NeuroMatrix")
//// app, err := application.Create(application.Options{
//// 	Title:     "NeuroMatrix",
//// 	Width:     1280,
//// 	Height:    600,
//// })
//// if err != nil {
//// 	panic(err)
//// }
//
//// // add GUI*********************************************************
//// // Create and add a label to the root panel
//// l1 := myGui.LabelFps(10, 10, "240")
//// app.Gui().Root().Add(l1)
//
//// // Create and add button 1 to the root panel
//// onOff := false
//// b1 := myGui.WebCam(10, 40, &onOff, app)
//// app.Gui().Root().Add(b1)
//
//// // Create and add exit button to the root panel
//// b2 := myGui.Exit(10, 70, &onOff, app)
//// app.Gui().Root().Add(b2)
////******************************************************************
//
//// // Создать и протестировать линии - синапсы
//
//// go func() {
//// 	myDots := 0
//// 	maxD := 700
//// 	dotlist := make(map[int]*drawing3D.Neuron3DBody)
//
//// 	for {
//// 		if myDots < maxD {
//// 			dotlist[myDots] = drawing3D.NewBody(app, math32.NewColor("White"))
//// 			dotlist[myDots].CreateBody()
//// 			//dotlist[myDots].SetPosition(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20)))
//// 			myDots++
//// 			//fmt.Println(len(dotlist), myDots)
//// 		}
//// 		if myDots == maxD {
//// 			for _, v := range dotlist {
//// 				v.SetPosition(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20)))
//// 				time.Sleep(time.Millisecond * 10)
//
//// 			}
//// 		}
//// 	}
//// }()
//
//// //Add lights to the scene
//// helpers.LightsScene(app)
//
//// // Add an axis helper to the scene
//// helpers.AxisHelper(0.5, app)
//
//// // Add an grid helper to the scene
//// helpers.GridHelper(10, app)
//
//// // Add camera to the scene
//// app.CameraPersp().SetPosition(15, 15, 15)
//// //app.Gl().ClearColor(0, 0.5, 0.7, 1)
//// app.Gl().ClearColor(0, 0.2, 0.4, 1)
//
//// // Start application
//// err = app.Run()
//// if err != nil {
//// 	panic(err)
//// }
