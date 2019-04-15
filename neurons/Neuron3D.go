package neurons

import (
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/util/application"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/geometry"
)

/*
	3D тело нейрона для отрисовки и позиционирования
	инициализировать на момент отрисовки, функция затухания, и удалять
	плавная градация света от белый - красный - жёлтый - зеленый - синий - чёрный
*/

// Neuron3DBody base 3d struct
type Neuron3DBody struct{
	Geom 	*geometry.Circle
	Mat  	*material.Phong
	Mesh 	*graphic.Points
	app 	*application.Application
}

// NewBody - constructor
func NewBody(app *application.Application) *Neuron3DBody{
	return &Neuron3DBody{
		Geom: 	geometry.NewCircle(0, 3),
		Mat: 	material.NewPhong(math32.NewColor("White")),
		app:    app,
	}
}

// CreateBody new body
func (nBody *Neuron3DBody) CreateBody(){
	nBody.Mesh = graphic.NewPoints(nBody.Geom, nBody.Mat)
	nBody.app.Scene().Add(nBody.Mesh)
	//nBody.IndxBody = nBody.app.Scene().ChildIndex(nBody.Mesh)
}

// SetPosition 3D Body
func (nBody *Neuron3DBody) SetPosition(x, y, z float32){
	nBody.Mesh.SetPosition(x, y, z)
}

// GetPosition 3D Body
func (nBody *Neuron3DBody) GetPosition() math32.Vector3{
	return nBody.Mesh.Position()
}

// DrawSynapse - create and draw lines(synapse)
// func(nBody *Neuron3DBody) DrawSynapse(start math32.Vector3, stop math32.Vector3, color *math32.Color){
// 	geom := geometry.NewGeometry()
// 	vertices := math32.NewArrayF32(0, 6)
// 	vertices.Append(
// 			start.X, start.Y, start.Z,
// 			stop.X, stop.Y, stop.Z,
// 	)
		
// 	colors := math32.NewArrayF32(0, 6)
// 	colors.Append(
// 		color.R, color.G, color.B,
// 		color.R, color.G, color.B,
// 	)
// 	geom.AddVBO(gls.NewVBO(vertices).AddAttrib(gls.VertexPosition))
// 	geom.AddVBO(gls.NewVBO(colors).AddAttrib(gls.VertexColor))

// 	// Creates basic material
// 	mat := material.NewBasic()

// 	// Creates lines with the specified geometry and material
// 	nBody.Synapse = graphic.NewLines(geom, mat)
// 	nBody.app.Scene().Add(nBody.Synapse)
// 	nBody.IndxSynapse = nBody.app.Scene().ChildIndex(nBody.Synapse)
// 	//fmt.Println(nBody.IndxSynapse)
// }
//Create a blue torus and add it to the scene
	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	//mat := material.NewPhong(math32.NewColor("DarkBlue"))
	//torusMesh := graphic.NewMesh(geom, mat)
	//app.Scene().Add(torusMesh)