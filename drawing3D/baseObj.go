package drawing3D

import (
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/math32"
)

// BaseObj base obj struct
type BaseObj struct {
	Node3D *core.Node
	Vec    math32.Vector3
}

func (bo *BaseObj) Node() *core.Node {
	return bo.Node3D
}

func (bo *BaseObj) GetLocation() *math32.Vector3 {
	return &bo.Vec
}

func (bo *BaseObj) SetLocation(vec math32.Vector3) {
	bo.Vec = vec
}

//// Neuron3D
//type Neuron3D struct {
//	BaseObj
//	dot  *graphic.Points
//	mesh *graphic.Mesh
//}
//
//func NewNeuron3D(vec math32.Vector3) *Neuron3D {
//	n := new(Neuron3D)
//	n.vec = vec
//	return n
//}
//
//func (n *Neuron3D) SetMeshPoint(point *graphic.Points) { //, light *light.Point) {
//	n.dot = point
//	n.node = &point.Node
//	point.SetPositionVec(&n.vec)
//}
//
//func (n *Neuron3D) SetMesh(mesh *graphic.Mesh) {
//	n.mesh = mesh
//	n.node = &mesh.Node
//	mesh.SetPositionVec(&n.vec)
//}

// Synapse
//type Synapse struct {
//	BaseObj
//	mesh *graphic.Lines
//}
//
//func NewSynapse(vec math32.Vector3) *Synapse {
//	s := new(Synapse)
//	s.Vec = vec
//	return s
//}
//
//func (s *Synapse) SetMeshLines(mesh *graphic.Lines) {
//	s.mesh = mesh
//	s.Node3D = &mesh.Node
//	mesh.SetPositionVec(&s.Vec)
//}
