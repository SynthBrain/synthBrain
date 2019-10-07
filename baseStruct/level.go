package baseStruct

import (
	"github.com/SynthBrain/synthBrain/drawing3D"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/math32"
	"math/rand"
)

// Level stores all the operational data for a level
type Level struct {
	synB   *SynthBrain
	scene  *core.Node
	camera *camera.Camera

	style *drawing3D.LevelStyle

	neurons         []*drawing3D.Neuron3D
	sizeListNeurons int

	//makeSynapse func(start *math32.Vector3, stop *math32.Vector3) *drawing3D.Neuron3D
	//remSynapse func(i int)
}
//
//// NewLevel - new Level object
//func NewLevel(synB *SynthBrain, ls *drawing3D.LevelStyle, cam *camera.Camera) *Level {
//
//	l := new(Level)
//	l.synB = synB
//	l.style = ls
//	l.camera = cam
//
//	l.scene = core.NewNode()
//	//l.scene.SetPosition(-ld.center.X, -ld.center.Y, -ld.center.Z)
//
//	//log.Debug("Starting NewLevel loop")
//
//	// меняеться лиш один а нужно что б все!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
//	l.sizeListNeurons = 1000
//	fmt.Println("Start new scene")
//	l.neurons = make([]*drawing3D.Neuron3D, l.sizeListNeurons)
//	l.makeSynapse, l.remSynapse = drawing3D.FabricSynapse(l.style, l.scene)
//	//if obj != nil {
//	//	switch obj := obj.(type) {
//	//	case *Neuron3D:
//	//		l.drawing3D = append(l.drawing3D, obj)
//	//
//	//		mesh := ls.makeRedBox()
//	//		light := light.NewPoint(l.style.boxLightColorOff, 1.0)
//	//
//	//		obj.SetMeshAndLight(mesh, light)
//	//		l.scene.Add(mesh)
//	//
//	//	}
//	//}
//
//	for i := 0; i < len(l.neurons); i++ {
//		//remSynapse(l.neurons[i].Node3D)
//		l.neurons[i] = l.makeSynapse(math32.NewVector3(float32(1),float32(1),float32(1)), math32.NewVector3(float32(2),float32(2),float32(2)))
//		//l.drawing3D[i] = NewNeuron3D(*math32.NewVector3(0,0,0))
//
//		// переиспользование нейронов возможно а вот синапсы придеться каждый раз пересоздавать
//		// что б луч был направлен куда нужно
//
//		// form:
//		// 0 = cub
//		// 1 = point
//		// 2 = line
//		//l.neurons[i] = drawing3D.FabricNeuron(l.style, l.scene, 0)
//		//l.scene.Add(l.neurons[i].Points)
//		//l.drawing3D = append(l.drawing3D, obj)
//
//		//mesh := ls.MakeWhiteNeuron()
//		//*****************************************
//		//dot := ls.MakeSuperWhiteDot()
//		//l.drawing3D[i].SetPoint(dot) //, light)
//		//l.drawing3D[i].SetPosition(math32.NewVector3(
//		//	float32(rand.Int31n(20)),
//		//	float32(rand.Int31n(20)),
//		//	float32(rand.Int31n(20))))
//		//l.scene.Add(l.drawing3D[i].Points)
//		//*****************************************
//		//dot := ls.MakeWhiteCube()
//		//l.neurons[i].SetMesh(dot) //, light)
//		//l.scene.Add(l.neurons[i].Mesh)
//		//*****************************************
//
//		//light := light.NewPoint(l.style.activeOff, 1.0)
//
//		//if i >= 1{
//		//	meshSynapse := ls.MakeSynapseLine(l.drawing3D[i].mesh.Position(), l.drawing3D[i-1].mesh.Position(), math32.NewColor("White"))
//		//	l.scene.Add(meshSynapse())
//		//}
//	}
//
//	// Add a single point light above the level
//	light := light.NewPoint(&math32.Color{1, 1, 1}, 8.0)
//	//light.SetPosition(l.data.center.X, l.data.center.Y*2+2, l.data.center.Z)
//	l.scene.Add(light)
//
//	return l
//}
//
//// SetPosition moves an object in the data grid along with its node to the desired position
//func (l *Level) SetPosition(obj interfaces.IBaseObj, dest math32.Vector3) {
//	//l.data.Set(obj.Location(), nil)
//	//obj.SetLocation(dest)
//	//l.data.Set(obj.Location(), obj)
//	//obj.Node().SetPositionVec(&dest)
//}
//
// Update updates all ongoing animations for the level
func (l *Level) Update(timeDelta float64) {
	//l.drawing3D[rand.Int31n(7000)].mesh.SetPositionVec(math32.NewVector3(float32(rand.Int31n(20)),
	//	float32(rand.Int31n(20)),
	//	float32(rand.Int31n(20))))

	//for i := 0; i < 70; i++ {
	//	l.drawing3D[i].mesh.SetPositionVec(math32.NewVector3(float32(rand.Int31n(20)),
	//		float32(rand.Int31n(20)),
	//		float32(rand.Int31n(20))))
	//	//time.Sleep(time.Millisecond * 10)
	//	//fmt.Println(i," ", l.drawing3D[i].GetLocation())
	//}

	for i := 0; i < len(l.neurons); i++ {
		l.NeuronGoThere(i)
	}
}

func (l *Level) NeuronGoThere(i int) {

	l.neurons[i].SetPosition(math32.NewVector3(float32(rand.Int31n(20)),
		float32(rand.Int31n(20)),
		float32(rand.Int31n(20))))
	//start := l.neurons[i].GetLocation()

	//l.scene.Remove(l.neurons[i].Node())
	//fmt.Println(l.neurons[i].IndexScene)
	//l.remSynapse(i)
	//l.neurons[i] = l.makeSynapse(start, math32.NewVector3(float32(rand.Int31n(20)),
	//	float32(rand.Int31n(20)),
	//	float32(rand.Int31n(20))))

}
