package baseLogic

import (
	"image"

	// "github.com/g3n/engine/geometry"
	// "github.com/g3n/engine/texture"
	// "github.com/g3n/engine/graphic"
	// "github.com/g3n/engine/material"

	"github.com/g3n/engine/math32"
)

/*
	Доступное поле зрения на нем фокус для которого есть
	1. нейронка которая учиться его двигать
	2. нейронка которая запоминает паттерны
	3. потребность которая заставляет работать первую на основе данных со второй
*/

/*
	для каждого органа свой полюсный пузырь
	такие пузыри пересекаються так как находяться в одном общем пузыре
	на местах пересечений создаються связи
*/

type Logic struct {
	VisionChan chan *[][]float32
	ImgChan    chan *image.Image
	data       [][]float32
	image      image.Image
	mapData    map[math32.Vector3]float32
}

func InitLogic() *Logic {
	return &Logic{
		VisionChan: make(chan *[][]float32, 1),
		ImgChan: make(chan *image.Image, 1),
	}
}

func (l *Logic) Update() {
	l.getDataFromChan()
}

func (l *Logic) getDataFromChan() {
	select {
	case dataKey := <-l.VisionChan:
		tempImg := <- l.ImgChan
		l.image = *tempImg
		l.data = *dataKey
		count := 0
		//vision.ImgToDataSlice(data)
		l.mapData = make(map[math32.Vector3]float32, len(l.data)*len(l.data[0]))
		tempPosition := *math32.NewVector3(0, 0, 0)
		for i := 0; i < len(l.data); i++ {
			for j := 0; j < len(l.data[i]); j++ {
				//fmt.Println("Start 2 ", j)
				//fmt.Print(data[i][j], " ")
				tempPosition.Set(float32(i), float32(j), 0)
				l.mapData[tempPosition] = l.data[i][j]
				//coords[count].Set(float32(i), float32(j), 0) //data[i][j]
				count++
			}
		}


		// Loads texture from image
		// texfile := app.DirData() + "/webCam.jpg"//"/images/tiger1.jpg"
		// tex2, err := texture.NewTexture2DFromImage(texfile)

		// //tex2, err := texture.NewTexture2DFromRGBA()
		// if err != nil {
		// 	app.Log().Fatal("Error:%s loading texture:%s", err, texfile)
		// }
		// // Creates plane2
		// plane2_geom := geometry.NewPlane(640, 480)
		// plane2_mat := material.NewStandard(&math32.Color{1, 1, 1})
		// plane2_mat.SetSide(material.SideDouble)
		// plane2_mat.AddTexture(tex2)
		// plane2 := graphic.NewMesh(plane2_geom, plane2_mat)
		// plane2.SetPosition(0, 0, 0)
		// app.Scene().Add(plane2)

		//fmt.Println(count)
		//app.Scene().RemoveAll(true) // NOTBAD
		//level.make3DLayer(0, count, coords, app)
	default:
		// залишаєм щоб не стопати фрейм рейт
	}
}

// отдавать данные для рисования сцени
func (l *Logic) GetData() *map[math32.Vector3]float32 {
	return &l.mapData
}
// отдавать данные для рисования сцени
func (l *Logic) GetImage() *image.Image {
	return &l.image
}
