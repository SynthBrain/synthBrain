package baseStruct

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/math32"
)

// BaseObj base obj struct
type BaseObj struct {
	node *core.Node
	vec  math32.Vector3
}

func (bo *BaseObj) Node() *core.Node {
	return bo.node
}

func (bo *BaseObj) GetLocation() math32.Vector3 {
	return bo.vec
}

func (bo *BaseObj) SetLocation(vec math32.Vector3) {
	bo.vec = vec
}

// Neuron3D
type Neuron3D struct {
	BaseObj
	mesh  *graphic.Mesh
	light *light.Point
}

func NewNeuron3D(vec math32.Vector3) *Neuron3D {
	n := new(Neuron3D)
	n.vec = vec
	return n
}

func (n *Neuron3D) SetMeshAndLight(mesh *graphic.Mesh, light *light.Point) {
	n.mesh = mesh
	n.node = &mesh.Node
	mesh.SetPositionVec(&n.vec)
	n.light = light
	n.node.Add(light)
}

// Synapse
type Synapse struct {
	BaseObj
	mesh *graphic.Mesh
}

func NewSynapse(vec math32.Vector3) *Synapse {
	s := new(Synapse)
	s.vec = vec
	return s
}

func (s *Synapse) SetMesh(mesh *graphic.Mesh) {
	s.mesh = mesh
	s.node = &mesh.Node
	mesh.SetPositionVec(&s.vec)
}
