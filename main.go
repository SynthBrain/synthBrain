package main

// 3840 * 2160 = 8 294 400
import (
	"flag"
	"github.com/SynthBrain/synthBrain/baseStruct"
	"github.com/SynthBrain/synthBrain/myGui"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/camera/control"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/logger"
	"github.com/g3n/engine/window"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var log *logger.Logger

/*
	Рисовать только тех что имеют достаточный уровень активность и окончательно не затухли
*/
func main() {
	// OpenGL functions must be executed in the same thread where
	// the context was created (by window.New())
	runtime.LockOSThread()

	// Parse command line flags
	showLog := flag.Bool("debug", false, "display the debug log")
	flag.Parse()

	// Create logger
	log = logger.New("SynthBrain", nil)
	log.AddWriter(logger.NewConsole(false))
	log.SetFormat(logger.FTIME | logger.FMICROS)
	if *showLog == true {
		log.SetLevel(logger.DEBUG)
	} else {
		log.SetLevel(logger.INFO)
	}
	log.Info("Initializing SynthBrain")

	// Create SynthBrain struct
	synB := new(baseStruct.SynthBrain)

	// Manually scan the $GOPATH directories to find the data directory
	rawPaths := os.Getenv("GOPATH")
	paths := strings.Split(rawPaths, ":")
	for _, j := range paths {
		// Checks data path
		path := filepath.Join(j, "src", "github.com", "SynthBrain", "synthBrain")
		if _, err := os.Stat(path); err == nil {
			synB.DataDir = path
		}
	}

	// Get the window manager
	var err error
	synB.Wmgr, err = window.Manager("glfw")
	if err != nil {
		panic(err)
	}

	// Create window and OpenGL context
	synB.Win, err = synB.Wmgr.CreateWindow(900, 640, "SynthBrain", false)
	if err != nil {
		panic(err)
	}

	// Create OpenGL state
	synB.Gs, err = gls.New()
	if err != nil {
		panic(err)
	}

	// Speed up a bit by not checking OpenGL errors
	synB.Gs.SetCheckErrors(false)

	// Sets window background color
	synB.Gs.ClearColor(0, 0.2, 0.4, 1) //(0.1, 0.1, 0.1, 1.0)

	// Sets the OpenGL viewport size the same as the window size
	// This normally should be updated if the window is resized.
	width, height := synB.Win.Size()
	synB.Gs.Viewport(0, 0, int32(width), int32(height))

	// Creates GUI root panel
	synB.Root = gui.NewRoot(synB.Gs, synB.Win)
	synB.Root.SetSize(float32(width), float32(height))

	// Update window if resize
	synB.Win.Subscribe(window.OnWindowSize, func(evname string, ev interface{}) {
		width, height := synB.Win.Size()
		synB.Gs.Viewport(0, 0, int32(width), int32(height))
		synB.Root.SetSize(float32(width), float32(height))
		aspect := float32(width) / float32(height)
		synB.Camera.SetAspect(aspect)
	})

	//add GUI*********************************************************
	// Create and add a label to the root panel
	synB.LabelFps = myGui.LabelFps(10, 10, "240")
	synB.Root.Add(synB.LabelFps)

	// Create and add button 1 to the root panel
	onOff := false
	chOnOffFlag := make(chan bool, 1)
	synB.WebCam = myGui.WebCam(10, 40, &onOff, chOnOffFlag)
	synB.Root.Add(synB.WebCam)

	// Create and add exit button to the root panel
	synB.Exit = myGui.Exit(10, 70, &onOff, synB.Win, chOnOffFlag)
	synB.Root.Add(synB.Exit)
	//****************************************************************

	// Creates a renderer and adds default shaders
	synB.Renderer = renderer.NewRenderer(synB.Gs)
	//g.renderer.SetSortObjects(false)
	err = synB.Renderer.AddDefaultShaders()
	if err != nil {
		panic(err)
	}
	synB.Renderer.SetGui(synB.Root)

	// Adds a perspective camera to the scene
	// The camera aspect ratio should be updated if the window is resized.
	aspect := float32(width) / float32(height)
	synB.Camera = camera.NewPerspective(65, aspect, 0.01, 1000)
	synB.Camera.SetPosition(10, 10, 10)
	synB.Camera.LookAt(&math32.Vector3{0, 0, 0})

	// Create orbit control and set limits
	synB.OrbitControl = control.NewOrbitControl(synB.Camera, synB.Win)
	synB.OrbitControl.Enabled = true //false
	synB.OrbitControl.EnablePan = false
	synB.OrbitControl.MaxPolarAngle = 2 * math32.Pi / 3
	synB.OrbitControl.MinDistance = 0.1
	synB.OrbitControl.MaxDistance = 10000

	// Create main scene and child levelScene
	synB.Scene = core.NewNode()
	synB.LevelScene = core.NewNode()
	synB.Scene.Add(synB.Camera)
	synB.Scene.Add(synB.LevelScene)
	synB.StepDelta = math32.NewVector2(0, 0)
	synB.Renderer.SetScene(synB.Scene)

	// Add white ambient light to the scene
	ambLight := light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.4)
	synB.Scene.Add(ambLight)

	synB.LevelStyle = baseStruct.NewBaseStyle(synB.DataDir)

	//synB.SetupGui(width, height)
	synB.RenderFrame()

	//synB.LoadSkyBox()
	synB.LoadLevels()

	size := 10
	gridHelp := graphic.NewGridHelper(float32(size), 1, math32.NewColor("LightGrey"))
	gridHelp.SetPosition(float32(size/2), -0.2, float32(size/2))
	synB.Scene.Add(gridHelp)

	// Done Loading - hide the loading label, show the menu, and initialize the level
	//synB.LoadingLabel.SetVisible(false)

	synB.InitLevel(0)

	now := time.Now()
	newNow := time.Now()
	log.Info("Starting Render Loop")
	// Start the render loop
	for !synB.Win.ShouldClose() {
		newNow = time.Now()
		timeDelta := now.Sub(newNow)
		now = newNow
		//fmt.Println(now)

		//vision.ReadImg(synB.DataDir, "/assets/0.jpg")
		synB.Update(timeDelta.Seconds())
		synB.RenderFrame()
	}

	//fmt.Printf("app was running for %f \n", application.Get().RunSeconds())
}

// go func() {
// 	for {
// 		if a, b, c := app.FrameRater().FPS(60); a > 0 && b > 0 && c == true {
// 			fmt.Println("FPS ", int(b))
// 		}
// 	}
// }()
//fps := float32(app.FrameCount()) / application.Get().RunSeconds()
//go myGui.LabelFpsTest(10, 10, strconv.Itoa(int(app.FrameCount()) / int(application.Get().RunSeconds())), app)

//IndCh := make(chan int)

// fmt.Println("Start NeuroMatrix")
// app, err := application.Create(application.Options{
// 	Title:     "NeuroMatrix",
// 	Width:     1280,
// 	Height:    600,
// })
// if err != nil {
// 	panic(err)
// }

// // add GUI*********************************************************
// // Create and add a label to the root panel
// l1 := myGui.LabelFps(10, 10, "240")
// app.Gui().Root().Add(l1)

// // Create and add button 1 to the root panel
// onOff := false
// b1 := myGui.WebCam(10, 40, &onOff, app)
// app.Gui().Root().Add(b1)

// // Create and add exit button to the root panel
// b2 := myGui.Exit(10, 70, &onOff, app)
// app.Gui().Root().Add(b2)
//******************************************************************

// // Создать и протестировать линии - синапсы

// go func() {
// 	myDots := 0
// 	maxD := 700
// 	dotlist := make(map[int]*neurons.Neuron3DBody)

// 	for {
// 		if myDots < maxD {
// 			dotlist[myDots] = neurons.NewBody(app, math32.NewColor("White"))
// 			dotlist[myDots].CreateBody()
// 			//dotlist[myDots].SetPosition(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20)))
// 			myDots++
// 			//fmt.Println(len(dotlist), myDots)
// 		}
// 		if myDots == maxD {
// 			for _, v := range dotlist {
// 				v.SetPosition(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20)))
// 				time.Sleep(time.Millisecond * 10)

// 			}
// 		}
// 	}
// }()

// //Add lights to the scene
// helpers.LightsScene(app)

// // Add an axis helper to the scene
// helpers.AxisHelper(0.5, app)

// // Add an grid helper to the scene
// helpers.GridHelper(10, app)

// // Add camera to the scene
// app.CameraPersp().SetPosition(15, 15, 15)
// //app.Gl().ClearColor(0, 0.5, 0.7, 1)
// app.Gl().ClearColor(0, 0.2, 0.4, 1)

// // Start application
// err = app.Run()
// if err != nil {
// 	panic(err)
// }
