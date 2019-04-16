package main

// // Create a tic-tac-toe board.
// board := [][]string{
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// 	[]string{"_", "_", "_"},
// }

// for i := 0; i < len(board); i++ {
// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
// }

// 3840 * 2160 = 8 294 400
import (
	"github.com/g3n/engine/math32"
	"time"
	"math/rand"
	"synthBrain/neurons"
	
	"fmt"
	"synthBrain/helpers"
	"synthBrain/myGui"
	"github.com/g3n/engine/util/application"
)

/*
	Рисовать только тех что имеют достаточный уровень активность и окончательно не затухли
*/
func main() {
	//IndCh := make(chan int)

	fmt.Println("Start NeuroMatrix")
	app, err := application.Create(application.Options{
		Title:     "NeuroMatrix",
		Width:     1280,
		Height:    600,
	})
	if err != nil {
		panic(err)
	}

	// add GUI*********************************************************
	// Create and add a label to the root panel
	l1 := myGui.LabelFps(10, 10, "240")
	app.Gui().Root().Add(l1)

	// Create and add button 1 to the root panel
	onOff := false
	b1 := myGui.WebCam(10, 40, &onOff, app)
	app.Gui().Root().Add(b1)

	// Create and add exit button to the root panel
	b2 := myGui.Exit(10, 70, &onOff, app)
	app.Gui().Root().Add(b2)
	//******************************************************************

	// Создать и протестировать линии - синапсы

	go func() {
		myDots := 70
		// заменить на мапы 
		var dotlist []*neurons.Neuron3DBody
		var synList []*neurons.Synapse
		for {
			if myDots > 0{
				nn := neurons.NewBody(app)
	 			nn.CreateBody()
				nn.SetPosition(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20)))
				dotlist = append(dotlist, nn)

				syn := neurons.NewSynapse(app, nn.GetPosition(), 
					dotlist[rand.Int31n(int32(len(dotlist)))].GetPosition() , math32.NewColor("LightGrey"))
				synList = append(synList, syn)
				//app.Scene().Add(syn)
				

				// dotlist[len(dotlist)-1].DrawSynapse(nn.GetPosition(),
				// 	dotlist[rand.Int31n(int32(len(dotlist)))].GetPosition() , math32.NewColor("LightGrey"))


				myDots--
			}
			if myDots == 0 {
				for _, v := range dotlist {
					v.SetPosition(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20)))
					//fmt.Println(v.IndxBody, " ", v.IndxSynapse)
				
					// if v.IndxSynapse >= 0{
					// 	app.Scene().RemoveAt(v.IndxSynapse)
					// }
					
					// IndCh <- v.IndxSynapse
					
					// v.DrawSynapse(v.GetPosition(),
					// 	dotlist[rand.Int31n(int32(len(dotlist)))].GetPosition(), math32.NewColor("LightGrey"))
					
					//fmt.Println(app.Scene().ChildIndex(v.Synapse), "Input")
					time.Sleep(time.Millisecond * 10)
				}
			}
		}
	}()

	// go func(){
	// 	for{
	// 		select{
	// 		case i := <-IndCh:
	// 			//app.Scene().RemoveAt(i)
	// 			fmt.Println(i, "Out")
	// 		case <-time.After(time.Second):
	// 			fmt.Println("OutEmpty")
	// 		}
	// 	}
	// }()
	

	//Add lights to the scene
	helpers.LightsScene(app)

	// Add an axis helper to the scene
	helpers.AxisHelper(0.5, app)

	// Add an grid helper to the scene
	helpers.GridHelper(10, app)

	// Add camera to the scene
	app.CameraPersp().SetPosition(15, 15, 15)
	//app.Gl().ClearColor(0, 0.5, 0.7, 1)
	app.Gl().ClearColor(0, 0.2, 0.4, 1)

	// Start application
	err = app.Run()
	if err != nil {
		panic(err)
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