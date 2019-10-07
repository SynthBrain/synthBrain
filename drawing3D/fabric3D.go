package drawing3D

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/math32"
)

// написати функцію що буде повертати різні варіанти нейрона для малювання в сцені
func FabricNeuron(style *LevelStyle, scene *core.Node, form int8) *Neuron3D {
	neuron := NewNeuron3D()
	switch form {
	case 0:
		neuron.SetMesh(style.MakeWhiteCube())
		scene.Add(neuron.Mesh)
		return neuron
	case 1:
		neuron.SetPoint(style.MakeSuperWhiteDot())
		scene.Add(neuron.Points)
		return neuron
	case 2:
		// сделать фабрику для синапсов с пересозданием
		tempStyle := style.MakeSynapseLine(neuron.GetLocation(), &math32.Vector3{
			neuron.GetLocation().X + 0.7,
			neuron.GetLocation().Y + 0.7,
			neuron.GetLocation().Z + 0.7}, math32.NewColor("White"))
		neuron.SetLines(tempStyle())
		scene.Add(neuron.Lines)
		return neuron
	default:
		neuron.SetPoint(style.MakeSuperWhiteDot())
		scene.Add(neuron.Points)
		return neuron
	}
}

func FabricSynapse(style *LevelStyle, scene *core.Node) (func(start *math32.Vector3, stop *math32.Vector3) *Neuron3D, func(i int)) {

	return func(start *math32.Vector3, stop *math32.Vector3) *Neuron3D {
		neuron := NewNeuron3D()
		tempStyle := style.MakeSynapseLine(start, stop, math32.NewColor("White"))
		neuron.SetLines(tempStyle())

		//neuron.IndexScene = len(scene.Children())
		scene.Add(neuron.Lines)
		return neuron

	}, func (i int) {
		scene.RemoveAt(i)
		//scene.Dispose()
	}
}


