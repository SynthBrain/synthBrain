package baseStruct

import (
	"fmt"
	"github.com/SynthBrain/synthBrain/interfaces"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
	"math/rand"
)

// Level stores all the operational data for a level
type Level struct {
	synB   *SynthBrain
	scene  *core.Node
	camera *camera.Perspective

	style *LevelStyle

	neurons         []*Neuron3D
	sizeListNeurons int
}

// NewLevel - new Level object
func NewLevel(synB *SynthBrain, ls *LevelStyle, cam *camera.Perspective) *Level {

	l := new(Level)
	l.synB = synB
	l.style = ls
	l.camera = cam

	l.scene = core.NewNode()
	//l.scene.SetPosition(-ld.center.X, -ld.center.Y, -ld.center.Z)

	//log.Debug("Starting NewLevel loop")

	l.sizeListNeurons = 30000
	fmt.Println("Start new scene")
	l.neurons = make([]*Neuron3D, l.sizeListNeurons)
	//if obj != nil {
	//	switch obj := obj.(type) {
	//	case *Neuron3D:
	//		l.neurons = append(l.neurons, obj)
	//
	//		mesh := ls.makeRedBox()
	//		light := light.NewPoint(l.style.boxLightColorOff, 1.0)
	//
	//		obj.SetMeshAndLight(mesh, light)
	//		l.scene.Add(mesh)
	//
	//	}
	//}
	for i := 0; i < l.sizeListNeurons; i++ {
		//l.neurons[i] = NewNeuron3D(*math32.NewVector3(0,0,0))
		l.neurons[i] = NewNeuron3D(*math32.NewVector3(float32(rand.Int31n(20)), float32(rand.Int31n(20)), float32(rand.Int31n(20))))
		//l.neurons = append(l.neurons, obj)

		//mesh := ls.MakeWhiteNeuron()
		dot := ls.MakeSuperWhiteDot()
		//light := light.NewPoint(l.style.activeOff, 1.0)

		l.neurons[i].SetMeshPoint(dot) //, light)

		l.scene.Add(l.neurons[i].point)

		//if i >= 1{
		//	meshSynapse := ls.MakeSynapseLine(l.neurons[i].mesh.Position(), l.neurons[i-1].mesh.Position(), math32.NewColor("White"))
		//	l.scene.Add(meshSynapse())
		//}
	}

	// Add a single point light above the level
	light := light.NewPoint(&math32.Color{1, 1, 1}, 8.0)
	//light.SetPosition(l.data.center.X, l.data.center.Y*2+2, l.data.center.Z)
	l.scene.Add(light)

	return l
}

// SetPosition moves an object in the data grid along with its node to the desired position
func (l *Level) SetPosition(obj interfaces.IBaseObj, dest math32.Vector3) {
	//l.data.Set(obj.Location(), nil)
	//obj.SetLocation(dest)
	//l.data.Set(obj.Location(), obj)
	//obj.Node().SetPositionVec(&dest)
}

// Update updates all ongoing animations for the level
func (l *Level) Update(timeDelta float64) {
	//l.neurons[rand.Int31n(7000)].mesh.SetPositionVec(math32.NewVector3(float32(rand.Int31n(20)),
	//	float32(rand.Int31n(20)),
	//	float32(rand.Int31n(20))))

	//for i := 0; i < 70; i++ {
	//	l.neurons[i].mesh.SetPositionVec(math32.NewVector3(float32(rand.Int31n(20)),
	//		float32(rand.Int31n(20)),
	//		float32(rand.Int31n(20))))
	//	//time.Sleep(time.Millisecond * 10)
	//	//fmt.Println(i," ", l.neurons[i].GetLocation())
	//}

	for i := 0; i < l.sizeListNeurons; i++ {
		go l.NeuronGoThere(i)
	}
}

func (l *Level) NeuronGoThere(i int) {
	l.neurons[i].point.SetPositionVec(math32.NewVector3(float32(rand.Int31n(100)),
		float32(rand.Int31n(100)),
		float32(rand.Int31n(100))))
}
