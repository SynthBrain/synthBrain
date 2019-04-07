package neurons

import (

	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/util/application"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/geometry"
)

/*
dotGeom := geometry.NewCircle(0, 3)
dotMat := material.NewPhong(math32.NewColor("White"))
dotMesh := graphic.NewPoints(dotGeom, dotMat)
dotMesh.SetPosition(
	float32(rand.Int31n(20)),
	float32(rand.Int31n(20)),
	float32(rand.Int31n(20)))
app.Scene().Add(dotMesh)
*/


/*
	3D тело нейрона для отрисовки и позиционирования
	инициализировать на момент отрисовки, функция затухания, и удалять
	плавная градация света от белый - красный - жёлтый - зеленый - синий - чёрный
*/

// Neuron3DBody base 3d struct
type Neuron3DBody struct{
	Geom *geometry.Circle
	Mat  *material.Phong
	Mesh *graphic.Points
	app *application.Application
}

// NewBody - constructor
func NewBody(app *application.Application) *Neuron3DBody{
	return &Neuron3DBody{
		Geom: 	geometry.NewCircle(0, 3),
		Mat: 	material.NewPhong(math32.NewColor("White")),
		//Mesh: 	graphic.NewPoints(Geom, Mat),
		app:    app,
	}
}

// CreateBody new body
func (nBody *Neuron3DBody) CreateBody(){
	nBody.Mesh = graphic.NewPoints(nBody.Geom, nBody.Mat)
	nBody.app.Scene().Add(nBody.Mesh)
}

// SetPosition 3D Body
func (nBody *Neuron3DBody) SetPosition(x, y, z float32){
	nBody.Mesh.SetPosition(x, y, z)
}


//Create a blue torus and add it to the scene
	//geom := geometry.NewTorus(1, .4, 12, 32, math32.Pi*2)
	//mat := material.NewPhong(math32.NewColor("DarkBlue"))
	//torusMesh := graphic.NewMesh(geom, mat)
	//app.Scene().Add(torusMesh)